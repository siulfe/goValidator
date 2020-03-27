package main


import (
	v "goValidator"
	"log"
)


func main(){
	validator := v.Validator{}
	var w int = 2
	a := &w

	v := validator.Field(*a).Length(0,255).Email()

	log.Println(v.Message)
}