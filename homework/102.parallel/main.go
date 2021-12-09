package main

import (
	"fmt"

)


func main()  {

	var line1,line2 []float64
	var a,b,c,d,k1,k2 float64

	fmt.Printf("请输入线条1坐标,格式为x1,y1,x2,y2 \n")
	fmt.Scanf("%v,%v,%v,%v",&a,&b,&c,&d)

	line1= append(line1,a)
	line1= append(line1,b)
	line1= append(line1,c)
	line1= append(line1,d)


	fmt.Printf("请输入线条2坐标,格式为x1,y1,x2,y2 \n")
	fmt.Scanf("%v,%v,%v,%v",&a,&b,&c,&d)
	line2= append(line2,a)
	line2= append(line2,b)
	line2= append(line2,c)
	line2= append(line2,d)

	k1= (line1[3]-line1[1])* (line2[2]-line2[0])
	k2= (line1[0]-line1[2])* (line2[3]-line2[1])

	fmt.Println(line1)
	fmt.Println(line2)

	if k1==k2 {
		fmt.Println("两直线平行")
	}else {
		fmt.Println("两直线不平行")
	}



}


