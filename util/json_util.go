package util

import (
	"bytes"
	"encoding/json"
	"errors"
)

// ObjectToJson
func ObjectToJson(src interface{}) (string, error) {
	if result, err := json.Marshal(src); err != nil {
		return "", errors.New("JsonStr to Object err: " + err.Error())
	} else {
		return string(result), nil
	}
}

// JsonToObject
func JsonToObject(src string, target interface{}) error {
	if err := json.Unmarshal([]byte(src), target); err != nil {
		return errors.New("JsonStr to Object err: " + err.Error())
	}
	return nil
}

// JsonToAny
func JsonToAny(src interface{}, target interface{}) error {
	if src == nil || target == nil {
		return errors.New("param is not null")
	}
	str, err := ObjectToJson(src)
	if err != nil {
		return err
	}
	if err := JsonToObject(str, target); err != nil {
		return err
	}
	return nil
}

// JsonToObject2
func JsonToObject2(src string, target interface{}) error {
	d := json.NewDecoder(bytes.NewBuffer([]byte(src)))
	d.UseNumber()
	if err := d.Decode(target); err != nil {
		return errors.New("JsonStr to Object err: " + err.Error())
	}
	return nil
}

// JsonToAny2
func JsonToAny2(src interface{}, target interface{}) error {
	if src == nil || target == nil {
		return errors.New("param is not null")
	}
	str, err := ObjectToJson(src)
	if err != nil {
		return err
	}
	if err := JsonToObject2(str, target); err != nil {
		return err
	}
	return nil
}
