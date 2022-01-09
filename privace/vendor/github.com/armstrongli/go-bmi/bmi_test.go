package gobmi

import (
	"fmt"
	"testing"
)

func TestBMI(t *testing.T) {
	//test1
	bmiDemo1,error1 := BMI(70,1.80)
	fmt.Println("demo1BMi:",bmiDemo1,"error1:",error1)
	if error1 != nil{
		t.Fail()
	}

	//test2
	bmiDemo2,error2 := BMI(-10,1.80)
	fmt.Println("demo2BMi:",bmiDemo2,"error2:",error2)
	if error2 != nil{
		t.Fail()
	}

	//test3
	bmiDemo3,error3 := BMI(80,0)
	fmt.Println("demo3BMi:",bmiDemo3,"error2:",error3)
	if error3 != nil{
		t.Fail()
	}

}
func TestFatRate(t *testing.T) {

	//test1  是否正常
	fatRateTest1, help1, error1 := FatRate(21,30,"男")
	fmt.Println("fatRate:",fatRateTest1,"help:",help1,"error:",error1)
	if error1 != nil {
		t.Fail()
	}

	//test2  当bmi为0 or 负数
	_, _, error2 := FatRate(-0,30,"男")
	//fmt.Println("fatRate:",fatRateTest1,"help:",help1,"error:",error1)
	if error2 == nil {
		t.Fail()
	}

	//test3  年龄是否为0
	_, _, error3 := FatRate(21,0,"男")
	//fmt.Println("fatRate:",fatRateTest1,"help:",help1,"error:",error1)
	if error3 == nil {
		t.Fail()
	}



	//test4  年龄是否小于0
	_, _, error4 := FatRate(23,-10,"男")
	//fmt.Println("fatRate:",fatRateTest1,"help:",help1,"error:",error1)
	if error4 == nil {
		t.Fail()
	}

	//test5  年龄是否大于150
	_, _, error5 := FatRate(23,250,"男")
	//fmt.Println("fatRate:",fatRateTest1,"help:",help1,"error:",error1)
	if error5 == nil {
		t.Fail()
	}

	//test6 当年龄小于18岁
	fatRateTest6, help6, error6 := FatRate(23,13,"男")
	fmt.Println("fatRate:",fatRateTest6,"help:",help6,"error:",error6)
	if error6 != nil {
		t.Fail()
	}

	//test7 测试性别输入
	_, _, error7 := FatRate(21,30,"你好")
	if error7 == nil {
		t.Fail()
	}


}
