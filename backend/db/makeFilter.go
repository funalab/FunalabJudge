package db

import (
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func MakeFilter(data interface{}) bson.M {
	bsonMap := bson.M{}
	rv := reflect.ValueOf(data)
	for i := 0; i < rv.NumField(); i++ {
		tag := rv.Type().Field(i).Tag.Get("bson")
		value, isDefault := getValueString(rv.Field(i))
		if !isDefault {
			bsonMap[tag] = value
		}
	}
	return bsonMap
}

func getValueString(f reflect.Value) (interface{}, bool) {
	// compare with defaul value of each type
	// _id(primitive.ObjectId)は対象にしない
	switch v := f.Interface().(type) {
	case int32:
		if v != 0 {
			return v, false
		}
	case string:
		if v != "" {
			return v, false
		}
	case time.Time:
		if v != time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC) {
			return v, false
		}
	case *bool:
		if v != nil {
			return *v, false
		}
	}
	return "", true // default value or unknown type
}
