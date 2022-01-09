package main

import (
	"fmt"
	gobmi "github.com/armstrongli/go-bmi"
)


func main(){
	tall,height := 65.0,1.85
	dbmi,_ := gobmi.BMI(tall,height)
	fatRateDemo,help,_ := gobmi.FatRate(dbmi, 30, "男")
	fmt.Println("fatRateDemo:", fatRateDemo)
	fmt.Println("help:", help)




}