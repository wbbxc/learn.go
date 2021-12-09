package main

import "fmt"

func main() {
	var weight, tall, bmi, fat [3]float64
	var num [3]int
	var sexNum int
	var sex, name, fatValues [3]string
	var sumFat, avgFat float64

	for i := 0; i < 3; i++ {
		fmt.Println("请输入姓名:")
		fmt.Scan(&name[i])

		fmt.Println("请输入体重(千克):")
		fmt.Scan(&weight[i])

		fmt.Println("请输入身高(米):")
		fmt.Scan(&tall[i])

		fmt.Println("请输入年龄:")
		fmt.Scan(&num[i])

		fmt.Println("请输入性别(男/女)")
		fmt.Scan(&sex[i])

		if sex[i] == "男" {
			sexNum = 1
		} else {
			sexNum = 0
		}

		bmi[i] = weight[i] / (tall[i] * tall[i])

		fat[i] = (1.2*bmi[i] + 0.23*float64(num[i]) - 5.4 - 10.8*float64(sexNum)) / 100
		fmt.Printf("体脂率为:%f \n", fat[i])

		fatValues[i] = getFatValues(num[i], sexNum, fat[i])
	}

	for i := 0; i < 3; i++ {
		fmt.Printf("姓名:%v,性别:%v,身高:%v,体重:%v,年龄:%v,BMI:%v,体脂率:%v,状态:%v \n",
			name[i], sex[i], tall[i], weight[i], num[i], bmi[i], fat[i], fatValues[i])
	}
	for _, v := range fat {
		sumFat += v
	}
	avgFat = sumFat / 3
	fmt.Printf("总人数:3个,平均为%v", avgFat)
}

func getFatValues(numC int, sexNumV int, fatC float64) (fatV string) {

	if sexNumV == 1 {
		switch {
		case numC >= 18 && numC <= 39:
			if fatC >= 0.0 && fatC <= 0.1 {
				fatV = "当前比较偏瘦"
			} else if fatC > 0.1 && fatC <= 0.16 {
				fatV = "标准身材"
			} else if fatC > 0.16 && fatC <= 0.21 {
				fatV = "当前偏重"
			} else if fatC > 0.21 && fatC <= 0.26 {
				fatV = "当前肥胖"
			} else {
				fatV = "严重肥胖"
			}
		case numC > 39 && numC <= 59:
			if fatC >= 0.0 && fatC <= 0.11 {
				fatV = "当前比较偏瘦"
			} else if fatC > 0.11 && fatC <= 0.17 {
				fatV = "标准身材"
			} else if fatC > 0.17 && fatC <= 0.22 {
				fatV = "当前偏重"
			} else if fatC > 0.22 && fatC <= 0.27 {
				fatV = "当前肥胖"
			} else {
				fatV = "严重肥胖"
			}
		case numC >= 60:
			if fatC >= 0.0 && fatC <= 0.13 {
				fatV = "当前比较偏瘦"
			} else if fatC > 0.13 && fatC <= 0.19 {
				fatV = "标准身材"
			} else if fatC > 0.19 && fatC <= 0.24 {
				fatV = "当前偏重"
			} else if fatC > 0.24 && fatC <= 0.29 {
				fatV = "当前肥胖"
			} else {
				fatV = "严重肥胖"
			}
		default:
			fatV = "小于18岁不统计该数据"
		}
	} else {
		switch {
		case numC >= 18 && numC <= 39:
			if fatC >= 0.0 && fatC <= 0.2 {
				fatV = "当前比较偏瘦"
			} else if fatC > 0.2 && fatC <= 0.27 {
				fatV = "标准身材"
			} else if fatC > 0.27 && fatC <= 0.34 {
				fatV = "当前偏重"
			} else if fatC > 0.34 && fatC <= 0.39 {
				fatV = "当前肥胖"
			} else {
				fatV = "严重肥胖"
			}
		case numC > 39 && numC <= 59:
			if fatC >= 0.0 && fatC <= 0.21 {
				fatV = "当前比较偏瘦"
			} else if fatC > 0.21 && fatC <= 0.28 {
				fatV = "标准身材"
			} else if fatC > 0.28 && fatC <= 0.35 {
				fatV = "当前偏重"
			} else if fatC > 0.35 && fatC <= 0.4 {
				fatV = "当前肥胖"
			} else {
				fatV = "严重肥胖"
			}
		case numC >= 60:
			if fatC >= 0.0 && fatC <= 0.22 {
				fatV = "当前比较偏瘦"
			} else if fatC > 0.22 && fatC <= 0.29 {
				fatV = "标准身材"
			} else if fatC > 0.29 && fatC <= 0.36 {
				fatV = "当前偏重"
			} else if fatC > 0.36 && fatC <= 0.41 {
				fatV = "当前肥胖"
			} else {
				fatV = "严重肥胖"
			}
		default:
			fatV = "小于18岁不统计该数据"
		}
	}

	return fatV
}
