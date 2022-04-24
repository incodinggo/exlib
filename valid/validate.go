package valid

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

type ValidateRequiredModel struct {
	validateModel
}

func (m *ValidateRequiredModel) validate() (result bool) {
	switch m.fieldT.Type.Kind() {
	case reflect.String:
		//result = len(strings.Replace(m.fieldV.(string), " ", "", -1)) > 0
		//去除前后空白进行验证,而不是整个替换空白符
		result = len(strings.TrimSpace(m.fieldV.(string))) > 0
	case reflect.Int:
		result = m.fieldV.(int) != 0
	case reflect.Int64:
		result = m.fieldV.(int64) != 0
	default:
		result = true
		fmt.Println("未验证参数:", m.fieldT.Name, " --> ", m.fieldV)
	}
	return result
}

type ValidateRangeModel struct {
	validateModel
	condition string
}

func (m *ValidateRangeModel) validate() (result bool) {

	if strings.Contains(m.condition, "[:") || strings.Contains(m.condition, ":]") {
		fmt.Println("不正确的表达式：", m.fieldT.Name, " -> ", m.condition)
		return
	}

	regValues := getRegIntValue(m.condition)
	var min, max = regValues[0], regValues[1]

	switch m.fieldT.Type.Kind() {
	case reflect.String:
		vLen := int64(len(strings.Replace(m.fieldV.(string), " ", "", -1)))
		if m.required && vLen == 0 {
			result = false
		}
		result = vLen == 0 || vLen >= min && vLen <= max
	case reflect.Int:
		v := int64(m.fieldV.(int))
		if m.required && v == 0 {
			result = false
		}
		result = v == 0 || v >= min && v <= max
	case reflect.Int64:
		v := m.fieldV.(int64)
		if m.required && v == 0 {
			result = false
		}
		result = v == 0 || v >= min && v <= max
	case reflect.Uint64:
		v := m.fieldV.(uint64)
		if m.required && v == 0 {
			result = false
		}
		result = v == 0 || v >= uint64(min) && v <= uint64(max)
	default:
		result = true
		fmt.Println("未验证参数:", m.fieldT.Name, " --> ", m.fieldV)
	}

	return result
}

type ValidateMinModel struct {
	validateModel
	condition string
}

func (m *ValidateMinModel) validate() (result bool) {

	regValues := getRegIntValue(m.condition)
	var min int64
	if len(regValues) > 0 {
		min = regValues[0]
	}

	switch m.fieldT.Type.Kind() {
	case reflect.String:
		//vLen := int64(len(strings.Replace(m.fieldV.(string), " ", "", -1)))
		//去除前后空白进行验证,而不是整个替换空白符
		vLen := int64(len(strings.TrimSpace(m.fieldV.(string))))
		if m.required && vLen == 0 {
			// 必填，但是为空
			result = false
		}
		result = vLen == 0 || vLen >= min
	case reflect.Int:
		v := int64(m.fieldV.(int))
		if m.required && v == 0 {
			result = false
		}
		result = v >= min
	case reflect.Int64:
		v := m.fieldV.(int64)
		if m.required || v > 0 {
			result = v == 0 || v >= min
		} else {
			result = true
		}
	case reflect.Uint:
		v := m.fieldV.(uint)
		if m.required && v == 0 {
			result = false
		}
		result = v == 0 || v >= uint(min)
	case reflect.Uint64:
		v := m.fieldV.(uint64)
		if m.required && v == 0 {
			result = false
		}
		result = v == 0 || v >= uint64(min)
	default:
		result = true
		fmt.Println("未验证参数:", m.fieldT.Name, " --> ", m.fieldV)
	}

	return
}

type ValidateMaxModel struct {
	validateModel
	condition string
}

func (m *ValidateMaxModel) validate() (result bool) {
	regValues := getRegIntValue(m.condition)
	var max int64
	if len(regValues) > 0 {
		max = regValues[0]
	}

	switch m.fieldT.Type.Kind() {
	case reflect.String:
		//vLen := int64(len(strings.Replace(m.fieldV.(string), " ", "", -1)))
		//去除前后空白进行验证,而不是整个替换空白符
		vLen := int64(len(strings.TrimSpace(m.fieldV.(string))))
		result = vLen == 0 || vLen <= max
		if !result {
			fmt.Println("ValidateMaxModel:", vLen)
		}
	case reflect.Int:
		v := int64(m.fieldV.(int))
		result = v == 0 || v <= max
	case reflect.Int64:
		v := m.fieldV.(int64)
		result = v == 0 || v <= max
	case reflect.Uint:
		v := m.fieldV.(uint)
		if m.required && v == 0 {
			result = false
		}
		result = v == 0 || v <= uint(max)
	case reflect.Uint64:
		v := m.fieldV.(uint64)
		if m.required && v == 0 {
			result = false
		}
		result = v == 0 || v <= uint64(max)
	default:
		result = true
		fmt.Println("未验证参数:", m.fieldT.Name, " --> ", m.fieldV)
	}

	return
}

type ValidatePassModel struct {
	validateModel
	condition string
}

func (m *ValidatePassModel) validate() (result bool) {

	if strings.Contains(m.condition, "[:") || strings.Contains(m.condition, ":]") {
		fmt.Println("不正确的表达式：", m.fieldT.Name, " -> ", m.condition)
		return
	}

	regValues := getRegIntValue(m.condition)
	var min, max int64
	if len(regValues) > 0 {
		min, max = regValues[0], regValues[1]
	}

	switch m.fieldT.Type.Kind() {
	case reflect.String:
		//pwd := strings.Replace(m.fieldV.(string), " ", "", -1)
		//去除前后空白进行验证,而不是整个替换空白符
		pwd := strings.TrimSpace(m.fieldV.(string))

		if !m.required && len(pwd) == 0 {
			// 非必填，未填写
			result = true
		} else if (min != 0 && len(pwd) < int(min)) || (max != 0 && len(pwd) > int(max)) {
			result = false
		} else if ok, err := regexp.MatchString(`^[\w_.,]+$`, pwd); err != nil {
			fmt.Println(err)
		} else if ok {
			result = true
		}
	default:
		// 密码字段必须为string
		fmt.Println("未验证参数:", m.fieldT.Name, " --> ", m.fieldV)
	}

	return
}

type ValidateEmailModel struct {
	validateModel
	condition string
}

func (m *ValidateEmailModel) validate() (result bool) {
	if strings.Contains(m.condition, "[:") || strings.Contains(m.condition, ":]") {
		fmt.Println("不正确的表达式：", m.fieldT.Name, " -> ", m.condition)
		return
	}

	regValues := getRegIntValue(m.condition)
	var min, max int64
	if len(regValues) > 0 {
		min, max = regValues[0], regValues[1]
	}

	switch m.fieldT.Type.Kind() {
	case reflect.String:
		//email := strings.Replace(m.fieldV.(string), " ", "", -1)
		//去除前后空白进行验证,而不是整个替换空白符
		email := strings.TrimSpace(m.fieldV.(string))
		if !m.required && len(email) == 0 {
			// 非必填，未填写
			result = true
		} else if (min != 0 && len(email) < int(min)) || (max != 0 && len(email) > int(max)) {
			result = false
		} else if ok, err := regexp.MatchString(`^[\w_.]+@[a-zA-Z0-9]{2,4}\.[a-z]{2,3}$`, email); err != nil {
			fmt.Println(err)
		} else if ok {
			result = true
		}
	default:
		// 邮箱字段必须为string
		fmt.Println("未验证参数:", m.fieldT.Name, " --> ", m.fieldV)
	}
	return
}

type ValidateTelModel struct {
	validateModel
	condition string
}

func (m *ValidateTelModel) validate() (result bool) {
	switch m.fieldT.Type.Kind() {
	case reflect.String:
		//tel := strings.Replace(m.fieldV.(string), " ", "", -1)
		//去除前后空白进行验证,而不是整个替换空白符
		tel := strings.TrimSpace(m.fieldV.(string))
		if m.required && len(tel) == 0 {
			// 必填，但为空
			result = false
		}
		if ok, err := regexp.MatchString(`^\+[0-9]{9,13}$`, tel); err != nil { //匹配+号
			fmt.Println(err)
		} else if ok {
			result = true
		} else if ok2, err := regexp.MatchString(`^[0-9]{9,13}$`, tel); err != nil { //不匹配+号
			fmt.Println(err)
		} else if ok2 {
			result = true
		}
	default:
		fmt.Println("未验证参数:", m.fieldT.Name, " --> ", m.fieldV)
	}
	return
}

type ValidateLetterModel struct {
	validateModel
	condition string
}

func (m *ValidateLetterModel) validate() (result bool) {
	if strings.Contains(m.condition, "[:") || strings.Contains(m.condition, ":]") {
		fmt.Println("不正确的表达式：", m.fieldT.Name, " -> ", m.condition)
		return
	}

	regValues := getRegIntValue(m.condition)
	var min, max int64
	if len(regValues) > 0 {
		min, max = regValues[0], regValues[1]
	}

	switch m.fieldT.Type.Kind() {
	case reflect.String:
		//letter := strings.Replace(m.fieldV.(string), " ", "", -1)
		//去除前后空白进行验证,而不是整个替换空白符
		letter := strings.TrimSpace(m.fieldV.(string))
		if m.required && len(letter) == 0 {
			// 必填，但为空
			result = true
		} else if (min != 0 && len(letter) < int(min)) || (max != 0 && len(letter) > int(max)) {
			result = false
		} else if ok, err := regexp.MatchString(`^[a-zA-Z0-9]+$`, letter); err != nil {
			fmt.Println(err)
		} else if ok {
			result = true
		}
	default:
		// 邮箱字段必须为string
		fmt.Println("未验证参数:", m.fieldT.Name, " --> ", m.fieldV)
	}
	return
}

type ValidateEnumModel struct {
	validateModel
	condition string
}

func (m *ValidateEnumModel) validate() (result bool) {

	// 汉字 字母 数字
	reg, _ := regexp.Compile(`[\p{Han}\w]*[^,\[\]]`)

	conds := reg.FindAllString(m.condition, -1)
	condLen := len(conds)

	if condLen < 2 {
		fmt.Println("不正确的表达式：", m.fieldT.Name, " -> ", m.condition)
		return
	}

	var value string
	switch m.fieldT.Type.Kind() {
	case reflect.String:
		//value = strings.Replace(m.fieldV.(string), " ", "", -1)
		//去除前后空白进行验证,而不是整个替换空白符
		value = strings.TrimSpace(m.fieldV.(string))
	case reflect.Int:
		fallthrough
	case reflect.Int64:
		value = fmt.Sprintf("%d", m.fieldV)
	default:
		// 密码字段必须为string
		fmt.Println("未验证参数:", m.fieldT.Name, " --> ", m.fieldV)
		return
	}
	if m.required && (value == "" || value == "0") {
		result = false
		return
	}
	if value == "" || value == "0" {
		result = true
		return
	}
	for i := 1; i < condLen; i++ {
		// 输入符合要求
		if conds[i] == value {
			result = true
			break
		}
	}
	return
}

func getRegIntValue(cond string) (values []int64) {
	reg, _ := regexp.Compile(`[-0-9]+`)
	regs := reg.FindAllString(cond, -1)

	for _, v := range regs {
		value, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			fmt.Println("need range[int:int] or min[int],but give string")
			values = append(values, -1)
			continue
		}
		values = append(values, value)
	}
	return
}

type ValidateDateTimeModel struct {
	validateModel
	condition string
}

func (m *ValidateDateTimeModel) validate() (result bool) {
	switch m.fieldT.Type.Kind() {
	case reflect.String:
		str, ok := m.fieldV.(string)
		if !ok {
			result = true
			return
		}
		if m.required && len(str) == 0 {
			// 必填，但为空
			result = true
		} else if ok, err := regexp.MatchString(`^\d{4}(-)(1[0-2]|0?\d)(-)([0-2]\d|0+\d|30|31)\s+(?:[01]\d|2[0-3]):[0-5]\d:[0-5]\d$`, str); err != nil {
			fmt.Println(err)
		} else if ok {
			result = true
		}
	default:
		fmt.Println("未验证参数:", m.fieldT.Name, " --> ", m.fieldV)
	}
	return
}

type ValidateDateModel struct {
	validateModel
	condition string
}

func (m *ValidateDateModel) validate() (result bool) {
	switch m.fieldT.Type.Kind() {
	case reflect.String:
		str, ok := m.fieldV.(string)
		if !ok {
			result = true
			return
		}
		if m.required && len(str) == 0 {
			// 必填，但为空
			result = true
		} else if ok, err := regexp.MatchString(`^\d{4}(-)(1[0-2]|0?\d)(-)([0-2]\d|0+\d|30|31)$`, str); err != nil {
			fmt.Println(err)
		} else if ok {
			result = true
		}
	default:
		fmt.Println("未验证参数:", m.fieldT.Name, " --> ", m.fieldV)
	}
	return
}

type ValidateTimeZoneModel struct {
	validateModel
	condition string
}

func (m *ValidateTimeZoneModel) validate() (result bool) {
	switch m.fieldT.Type.Kind() {
	case reflect.String:
		str, ok := m.fieldV.(string)
		if !ok {
			result = true
			return
		}
		if m.required && len(str) == 0 {
			// 必填，但为空
			result = true
		} else if ok, err := regexp.MatchString(`^[a-zA-Z]+/[a-zA-Z]+$`, str); err != nil {
			fmt.Println(err)
		} else if ok {
			result = true
		}
	default:
		fmt.Println("未验证参数:", m.fieldT.Name, " --> ", m.fieldV)
	}
	return
}

type ValidateMinLenModel struct {
	validateModel
	condition string
}

func (m *ValidateMinLenModel) validate() (result bool) {
	regValues := getRegIntValue(m.condition)
	var min int64
	if len(regValues) > 0 {
		min = regValues[0]
	}
	str, ok := m.fieldV.(string)
	var vLen int64 = 0
	if !ok {
		fmt.Println("未验证参数:", m.fieldT.Name, " --> ", m.fieldV)
	} else {
		vLen = int64(utf8.RuneCountInString(str))
	}
	if vLen >= min {
		result = true
	}

	return
}

type ValidateMaxLenModel struct {
	validateModel
	condition string
}

func (m *ValidateMaxLenModel) validate() (result bool) {
	regValues := getRegIntValue(m.condition)
	var max int64
	if len(regValues) > 0 {
		max = regValues[0]
	}
	str, ok := m.fieldV.(string)
	var vLen int64 = 0
	if !ok {
		fmt.Println("未验证参数:", m.fieldT.Name, " --> ", m.fieldV)
	} else {
		vLen = int64(utf8.RuneCountInString(str))
	}
	if vLen <= max {
		result = true
	}

	return
}

type ValidateMinTrLenModel struct {
	validateModel
	condition string
}

func (m *ValidateMinTrLenModel) validate() (result bool) {
	regValues := getRegIntValue(m.condition)
	var min int64
	if len(regValues) > 0 {
		min = regValues[0]
	}
	str, ok := m.fieldV.(string)
	var vLen int64 = 0
	if !ok {
		fmt.Println("未验证参数:", m.fieldT.Name, " --> ", m.fieldV)
	} else {
		str := strings.TrimSpace(str)
		vLen = int64(utf8.RuneCountInString(str))
	}
	if vLen >= min {
		result = true
	}

	return
}

type ValidateMaxTrLenModel struct {
	validateModel
	condition string
}

func (m *ValidateMaxTrLenModel) validate() (result bool) {
	regValues := getRegIntValue(m.condition)
	var max int64
	if len(regValues) > 0 {
		max = regValues[0]
	}
	str, ok := m.fieldV.(string)
	var vLen int64 = 0
	if !ok {
		fmt.Println("未验证参数:", m.fieldT.Name, " --> ", m.fieldV)
	} else {
		str := strings.TrimSpace(str)
		vLen = int64(utf8.RuneCountInString(str))
	}
	if vLen <= max {
		result = true
	}

	return
}
