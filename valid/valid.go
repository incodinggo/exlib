package valid

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

/**

校验方式      适用字段  说明
required      string    必填(验证非0值或非空字符串,如果可能包含0值或空字符串,则此检验不适用)
enum          string    在指定的数据中 [q,e,r,中文]
              int                      [1,3,9]
              int64
              uint64

min,max       string    最(大)小长度，可以分开使用,会去除前后空白进行验证
              int       最(大)小值，可以分开使用

range[0:10]   string    长度在指定的范围内,会去除前后空白进行验证
              int       大小在指定的范围内

minlen,maxlen  string    最(大)小长度，可以分开使用,不去除前后空白进行验证,识别中英文长度

minlentr,maxlentr  string    最(大)小长度，可以分开使用,去除前后空白进行验证,识别中英文长度

range[0:10]   string    长度在指定的范围内
              int       大小在指定的范围内

password      string    验证是否符合密码规范(字母/数字/.-_)
password[0:1] string    验证是否符合密码规范(字母/数字/.-_),并在指定的范围
email         string    验证是否符合邮箱规范
email[0:1]    string    验证是否符合邮箱规范,并在指定的范围
letter        string    验证是否是 字母+数字
letter[0:1]   string    验证是否是 字母+数字,并在指定的范围
phone         string    是否符合手机号验证规范
datetime	  string	是否为日期时间格式 2016-01-02 15:04:05
date	      string	是否为日期格式 2016-01-02
timezone	  string	是否为时区格式 Asia/Shanghai
*/

const (
	Required = "required"
	Enum     = "enum"
	Min      = "min"
	Max      = "max"
	Range    = "range"
	Password = "password"
	Email    = "email"
	Phone    = "phone"
	Letter   = "letter"
	DateTime = "datetime"
	Date     = "date"
	TimeZone = "timezone"
	MinLen   = "minlen"
	MaxLen   = "maxlen"
	MinLenTr = "minlentr"
	MaxLenTr = "maxlentr"
)

func Validate(reqModel interface{}) error {

	typeV := reflect.ValueOf(reqModel)
	if typeV.Kind() == reflect.Ptr {
		typeV = typeV.Elem()
	}
	typeT := typeV.Type()

	return validateCheck(typeT, typeV)
}

func validateCheck(typeT reflect.Type, typeV reflect.Value) error {
	for i := 0; i < typeT.NumField(); i++ {
		fieldT := typeT.Field(i)
		fieldV := typeV.Field(i)
		//如果是匿名结构体,需要递归判断
		if fieldT.Anonymous && fieldT.Type.Kind() == reflect.Struct {
			err := validateCheck(fieldT.Type, fieldV)
			if err != nil {
				return err
			}
			continue
		}

		// 是否存在校验字段
		validCond := fieldT.Tag.Get("valid")
		if len(validCond) == 0 {
			continue
		}

		// 如果校验出错，直接返回。不需要判断所有条件
		if err := validate(validCond, fieldT, typeV.FieldByName(fieldT.Name).Interface()); err != nil {
			return err
		}
	}
	return nil
}

func validate(validCond string, fieldT reflect.StructField, fieldV interface{}) error {

	// 是否必须
	_validateModel := validateModel{fieldT: fieldT, fieldV: fieldV}

	validSlice := strings.Split(validCond, ";")
	if strings.Index(validCond, "|") != -1 {
		validSlice = append(validSlice, strings.Split(validCond, "|")...)
	}

	for _, valid := range validSlice {
		slice := strings.Split(valid, "#")
		v := valid
		e := ""
		if len(slice) > 1 {
			v = slice[0]
			e = slice[1]
		}

		if len(v) == 0 {
			continue
		}

		var valid ValidateInterface
		num, err := strconv.ParseFloat(e, 64)
		if err != nil {
			err = nil
			num = 0
		}
		if strings.Index(v, Required) == 0 {
			// 必填
			_validateModel.required = true
			valid = &ValidateRequiredModel{validateModel: _validateModel}
		} else if strings.Index(v, Range) == 0 {
			// range
			valid = &ValidateRangeModel{condition: v, validateModel: _validateModel}
		} else if strings.Index(v, Min) == 0 {
			// min
			valid = &ValidateMinModel{condition: v, validateModel: _validateModel}
		} else if strings.Index(v, Max) == 0 {
			// max
			valid = &ValidateMaxModel{condition: v, validateModel: _validateModel}
		} else if strings.Index(v, Enum) == 0 {
			// enum
			valid = &ValidateEnumModel{condition: v, validateModel: _validateModel}
		} else if strings.Index(v, Password) == 0 {
			// password
			valid = &ValidatePassModel{condition: v, validateModel: _validateModel}
		} else if strings.Index(v, Email) == 0 {
			// email
			valid = &ValidateEmailModel{condition: v, validateModel: _validateModel}
		} else if strings.Index(v, Letter) == 0 {
			// letter
			valid = &ValidateLetterModel{condition: v, validateModel: _validateModel}
		} else if strings.Index(v, Phone) == 0 {
			// phone
			valid = &ValidateTelModel{condition: v, validateModel: _validateModel}
		} else if strings.Index(v, DateTime) == 0 {
			//datetime
			valid = &ValidateDateTimeModel{condition: v, validateModel: _validateModel}
		} else if strings.Index(v, Date) == 0 {
			//date
			valid = &ValidateDateModel{condition: v, validateModel: _validateModel}
		} else if strings.Index(v, TimeZone) == 0 {
			//timezone
			valid = &ValidateTimeZoneModel{condition: v, validateModel: _validateModel}
		} else if strings.Index(v, MinLen) == 0 {
			//minlen
			valid = &ValidateMinLenModel{condition: v, validateModel: _validateModel}
		} else if strings.Index(v, MaxLen) == 0 {
			//maxlen
			valid = &ValidateMaxLenModel{condition: v, validateModel: _validateModel}
		} else if strings.Index(v, MinLenTr) == 0 {
			//minlen
			valid = &ValidateMinLenModel{condition: v, validateModel: _validateModel}
		} else if strings.Index(v, MaxLenTr) == 0 {
			//maxlen
			valid = &ValidateMaxLenModel{condition: v, validateModel: _validateModel}
		} else {
			fmt.Println("不支持的语法:", v)
			continue
		}

		if !valid.validate() {
			return ValidateError{
				Field:   fieldT.Name,
				Valid:   validCond,
				ErrCode: int(num),
			}
		}
	}

	return nil
}
