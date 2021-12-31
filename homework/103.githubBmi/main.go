package main

import (
	"fmt"
	gobmi "github.com/armstrongli/go-bmi"
)


func main(){

	//完成流程
	tall,height,age,sex := 75.0,1.80,30,"男"
	dbmi,_ := gobmi.BMI(tall,height)
	fatRateDemo,help,_ := gobmi.FatRate(dbmi, age, sex)
	fmt.Println("fatRateDemo:", fatRateDemo)
	fmt.Println("help:", help)



}


