package main

import (
	"fmt"
	gobmi "github.com/armstrongli/go-bmi"
)


func main(){
<<<<<<< HEAD
	tall,height := 65.0,1.85
	dbmi,_ := gobmi.BMI(tall,height)
	fatRateDemo,help,_ := gobmi.FatRate(dbmi, 30, "男")
=======

	//完成流程
	tall,height,age,sex := 75.0,1.80,30,"男"
	dbmi,_ := gobmi.BMI(tall,height)
	fatRateDemo,help,_ := gobmi.FatRate(dbmi, age, sex)
>>>>>>> 1c4330c8e69a3b519e0790b9e4b60ddd478720e7
	fmt.Println("fatRateDemo:", fatRateDemo)
	fmt.Println("help:", help)



<<<<<<< HEAD

}
=======
}


>>>>>>> 1c4330c8e69a3b519e0790b9e4b60ddd478720e7
