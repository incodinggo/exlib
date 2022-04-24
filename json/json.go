package json

import jsoniter "github.com/json-iterator/go"

func Parse(data []byte, v interface{}) {
	err := jsoniter.Unmarshal(data, v)
	if err != nil {
		return
	}
}

func ParseE(data []byte, v interface{}) error {
	err := jsoniter.Unmarshal(data, v)
	return err
}

func Stringify(v interface{}) []byte {
	b, err := jsoniter.Marshal(v)
	if err != nil {
		return nil
	}
	return b
}

func StringifyE(v interface{}) ([]byte, error) {
	b, err := jsoniter.Marshal(v)
	if err != nil {
		return nil, err
	}
	return b, nil
}
