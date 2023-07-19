package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"leetcodehelper/helper/model"
	"reflect"
	"strconv"
	"strings"
)

func Code(param string, fn interface{}) {
	fnType := reflect.TypeOf(fn)
	fnValue := reflect.ValueOf(fn)

	if fnType.Kind() != reflect.Func {
		fmt.Printf("Error occurred: Second parameter must be a function.")
	}
	fnTypeLen := fnType.NumIn()
	strParameters, err := getStringParameters(param, fnTypeLen)
	if err != nil {
		fmt.Printf("Error occurred: %v", err)
	}

	params, err := strParametersToParameters(strParameters, fnType)
	if err != nil {
		fmt.Printf("Error occurred: %v", err)
	}
	results := fnValue.Call(params)
	for _, result := range results {
		if result.Type().Implements(helperNodeType) {
			fmt.Println(result)
		} else {
			jsonBytes, err := json.Marshal(result.Interface())
			if err == nil {
				fmt.Println(string(jsonBytes))
			} else {
				fmt.Println("Error occurred: %v", err)
			}
		}
	}
}

func getStringParameters(parameter string, length int) ([]string, error) {
	strlen := len(parameter)
	strParameters := make([]string, length)
	index := strings.LastIndex(parameter, " = ")

	if index == -1 {
		return nil, errors.New("Illegal parameter")
	}

	for strlen > 0 {
		length--
		if length < 0 {
			return nil, errors.New("The number of parameter lists does not match")
		} else {
			strParameters[length] = strings.TrimSpace(parameter[index+3 : strlen])
			strlen = strings.LastIndex(parameter[:strlen-1], ", ")
			if strlen == -1 {
				break
			}
			index = strings.LastIndex(parameter[:strlen], " = ")
		}
	}
	return strParameters, nil
}

func strParametersToParameters(paramSlice []string, fnType reflect.Type) ([]reflect.Value, error) {
	fnTypeLen := len(paramSlice)
	params := make([]reflect.Value, fnTypeLen)

	for i := 0; i < fnTypeLen; i++ {
		paramType := fnType.In(i)
		paramName := paramType.Name()
		valueStr := paramSlice[i]

		var value reflect.Value

		paramKind := paramType.Kind()

		if paramType.Implements(helperNodeType) {
			helperObject := reflect.New(paramType).Elem().Interface()
			if helperNode, ok := helperObject.(model.HelperNode); ok {
				// 调用接口的方法
				helperNodeValue, err := helperNode.Convert(valueStr)
				if err != nil {
					return nil, err
				}
				value = reflect.ValueOf(helperNodeValue)
			} else {
				return nil, errors.New("YourObject does not implement YourInterface.")
			}
		} else {
			switch paramKind {
			case reflect.Int:
				intValue, err := strconv.Atoi(valueStr)
				if err != nil {
					return nil, errors.New("Invalid number format for parameter: " + paramName)
				}
				value = reflect.ValueOf(intValue)
			case reflect.String:
				value = reflect.ValueOf(valueStr)
			case reflect.Bool:
				boolValue, err := strconv.ParseBool(valueStr)
				if err != nil {
					return nil, errors.New("Invalid bool format for parameter: " + paramName)
				}
				value = reflect.ValueOf(boolValue)
			case reflect.Slice:
				sliceValue, err := buildSlice(paramType, valueStr)
				if err != nil {
					return nil, err
				}
				value = sliceValue
			default:
				return nil, errors.New("Unsupported parameter type for parameter: " + paramName)
			}
		}
		params[i] = value
	}
	return params, nil
}

func buildSlice(paramType reflect.Type, valueStr string) (reflect.Value, error) {
	var value reflect.Value
	var err error
	switch paramType.Elem().Kind() {
	case reflect.Int:
		var slice []int
		err = json.Unmarshal([]byte(valueStr), &slice)
		if err != nil {
			return value, err
		}
		value = reflect.ValueOf(slice)
	case reflect.String:
		var slice []string
		err = json.Unmarshal([]byte(valueStr), &slice)
		if err != nil {
			return value, err
		}
		value = reflect.ValueOf(slice)
	case reflect.Bool:
		var slice []bool
		err = json.Unmarshal([]byte(valueStr), &slice)
		if err != nil {
			return value, err
		}
		value = reflect.ValueOf(slice)
	default:
		err = errors.New("Unsupported slice element type for parameter")
	}
	return value, err
}

var helperNodeType = reflect.TypeOf((*model.HelperNode)(nil)).Elem()
