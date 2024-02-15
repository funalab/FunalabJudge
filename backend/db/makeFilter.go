package db

import (
	"reflect"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
)

// 未完成
func MakeFilterWithNonnilField(data interface{}) bson.M {
	// TODO 引数の定義にアサーション入れたほうがいい？
	bsonMap := bson.M{}
	v := reflect.ValueOf(data)
	for i := 0; i < v.NumField(); i++ {
		tag := v.Type().Field(i).Tag.Get("bson")
		println(tag)
		println(v.Field(i).Kind())
		println(v.Field(i).IsNil())
		if !v.Field(i).IsNil() {
			value := parseReflectValue(v.Field(i))
			bsonMap[tag] = value
		}
	}
	return bsonMap
}

func parseReflectValue(rv reflect.Value) string {
	switch rv.Kind() {
	case reflect.Bool:
		if rv.Bool() {
			return "true"
		} else {
			return "false"
		}
	case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
		return strconv.FormatInt(rv.Int(), 10)
	case reflect.Uint, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8:
		return strconv.FormatUint(rv.Uint(), 10)
	case reflect.Float32:
		return strconv.FormatFloat(rv.Float(), 'f', -1, 32)
	case reflect.Float64:
		return strconv.FormatFloat(rv.Float(), 'f', -1, 64)
	case reflect.String:
		return rv.String()
	default:
		return ""
	}
}
