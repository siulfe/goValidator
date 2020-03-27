package goValidator

import (
	_"log"
	//"strconv"
	"reflect"
)


type Validator struct{
	field interface{}
	IsValid bool
	Message string
	EmailExpresion string
}

func (v Validator) Field(field interface{}) Validator{
	v.field = field
	v.IsValid = true
	EmailExpresion = ""

	return v
}


func (v Validator) Length(min int,max int) Validator{
	if !v.IsValid{
		return v
	}

	if v = v.LengthMin(min); !v.IsValid{
		return v
	}

	v = v.LengthMax(max)

	return v
}

func (v Validator) LengthMin(min int) Validator{
	if !v.IsValid{
		return v
	}

	item := reflect.ValueOf(v.field)

	if item.Kind() == reflect.String || item.Kind() == reflect.Slice || item.Kind() == reflect.Array || item.Kind() == reflect.Map{
		if item.Len() < min{
			v.IsValid = false
			v.Message = lessMinConst
		}
		return v
	}

	if item.Kind() == reflect.Struct{
		if item.NumField() < min{
			v.IsValid = false
			v.Message = lessMinConst
		}
		return v
	}

	if item.Kind() == reflect.Float32 || item.Kind() == reflect.Float64{
		if item.Float() < float64(min){
			v.IsValid = false
			v.Message = lessMinConst
		}

		return v
	}

	if item.Kind () == reflect.Int || item.Kind () == reflect.Int8 || item.Kind () == reflect.Int16 || item.Kind () == reflect.Int32 || 
		item.Kind () == reflect.Int64 || item.Kind () == reflect.Uint || item.Kind () == reflect.Uint8 || item.Kind () == reflect.Uint16 ||
		item.Kind () == reflect.Uint32 || item.Kind () == reflect.Uint64 {
		if item.Int() < int64(min){
			v.IsValid = false
			v.Message = lessMinConst
		}
		return v
	}

	v.IsValid = false
	v.Message = invalidTypeConst + "Length"

	return v;
}

func (v Validator) LengthMax(max int) Validator{
	if !v.IsValid{
		return v
	}

	item := reflect.ValueOf(v.field)

	if item.Kind() == reflect.String || item.Kind() == reflect.Slice || item.Kind() == reflect.Array || item.Kind() == reflect.Map{
		if item.Len() > max{
			v.IsValid = false
			v.Message = greatherMaxConst
		}
		return v
	}

	if item.Kind() == reflect.Struct{
		if item.NumField() > max{
			v.IsValid = false
			v.Message = greatherMaxConst
		}
		return v
	}

	if item.Kind() == reflect.Float32 || item.Kind() == reflect.Float64{
		if item.Float() > float64(max){
			v.IsValid = false
			v.Message = greatherMaxConst
		}

		return v
	}

	if item.Kind () == reflect.Int || item.Kind () == reflect.Int8 || item.Kind () == reflect.Int16 || item.Kind () == reflect.Int32 || 
		item.Kind () == reflect.Int64 || item.Kind () == reflect.Uint || item.Kind () == reflect.Uint8 || item.Kind () == reflect.Uint16 ||
		item.Kind () == reflect.Uint32 || item.Kind () == reflect.Uint64 {
		if item.Int() > int64(max){
			v.IsValid = false
			v.Message = greatherMaxConst
		}
		return v
	}

	v.IsValid = false
	v.Message = invalidTypeConst + "Length"

	return v;
}

func (v Validator) Email() Validator{
	if !v.IsValid{
		return v
	}
	
}