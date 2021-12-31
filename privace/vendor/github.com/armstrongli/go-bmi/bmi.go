package gobmi

import "fmt"


func BMI(weightKG, heightM float64) (bmi float64, err error) {
	if weightKG <= 0 {
		err = fmt.Errorf("weight cannot be negative")
		return
	}
	if heightM < 0 {
		err = fmt.Errorf("height cannot be negative")
		return
	}
	if heightM == 0 {
		err = fmt.Errorf("height cannot be 0")
		return
	}
	bmi = weightKG / (heightM * heightM)
	return
}

func FatRate(bmi float64, age int,sex string) (fatRate float64,proposal string,err error){

	var sexNum int = 0

	//判断BMI的值
	if bmi <= 0 {
		err = fmt.Errorf("bmi 不能小于0")
		return 0,"",err
	}

	//判断年龄
	if age <= 0 || age >= 150 {
		err = fmt.Errorf("年龄不能是0、负数和超过150岁")
		return 0,"",err
	}

	//判断性别
	if sex == "男" {
		sexNum = 1
	}else if sex == "女" {
		sexNum = 0
	}else{
		err = fmt.Errorf("输入的性别有误,男or女")
		return 0,"",err
	}

	fatRate = (1.2* bmi + 0.23 *float64(age) -5.4 - 10.8 * float64(sexNum)) / 100

	if sexNum == 1 {
		proposal = boyFatRate(age,fatRate)
	}else {
		proposal = girlFatRate(age,fatRate)
	}



	return fatRate, proposal,nil

}

func boyFatRate(age int,fatRatec float64) (help string) {
	switch {
	case age >= 18 && age <= 39:
		if fatRatec >= 0.0 && fatRatec <= 0.1 {
			help = "当前比较偏瘦"
		} else if fatRatec > 0.1 && fatRatec <= 0.16 {
			help = "标准身材"
		} else if fatRatec > 0.16 && fatRatec <= 0.21 {
			help = "当前偏重"
		} else if fatRatec > 0.21 && fatRatec <= 0.26 {
			help = "当前肥胖"
		} else {
			help = "严重肥胖"
		}
	case age > 39 && age <= 59:
		if fatRatec >= 0.0 && fatRatec <= 0.11 {
			help = "当前比较偏瘦"
		} else if fatRatec > 0.11 && fatRatec <= 0.17 {
			help = "标准身材"
		} else if fatRatec > 0.17 && fatRatec <= 0.22 {
			help = "当前偏重"
		} else if fatRatec > 0.22 && fatRatec <= 0.27 {
			help = "当前肥胖"
		} else {
			help = "严重肥胖"
		}
	case age >= 60:
		if fatRatec >= 0.0 && fatRatec <= 0.13 {
			help = "当前比较偏瘦"
		} else if fatRatec > 0.13 && fatRatec <= 0.19 {
			help = "标准身材"
		} else if fatRatec > 0.19 && fatRatec <= 0.24 {
			help = "当前偏重"
		} else if fatRatec > 0.24 && fatRatec <= 0.29 {
			help = "当前肥胖"
		} else {
			help = "严重肥胖"
		}
	default:
		help = "小于18岁不统计该数据"
	}

	return help
}

func girlFatRate(age int,fatRatec float64) (help string) {
	switch {
	case age >= 18 && age <= 39:
		if fatRatec >= 0.0 && fatRatec <= 0.2 {
			help = "当前比较偏瘦"
		} else if fatRatec > 0.2 && fatRatec <= 0.27 {
			help = "标准身材"
		} else if fatRatec > 0.27 && fatRatec <= 0.34 {
			help = "当前偏重"
		} else if fatRatec > 0.34 && fatRatec <= 0.39 {
			help = "当前肥胖"
		} else {
			help = "严重肥胖"
		}
	case age > 39 && age <= 59:
		if fatRatec >= 0.0 && fatRatec <= 0.21 {
			help = "当前比较偏瘦"
		} else if fatRatec > 0.21 && fatRatec <= 0.28 {
			help = "标准身材"
		} else if fatRatec > 0.28 && fatRatec <= 0.35 {
			help = "当前偏重"
		} else if fatRatec > 0.35 && fatRatec <= 0.4 {
			help = "当前肥胖"
		} else {
			help = "严重肥胖"
		}
	case age >= 60:
		if fatRatec >= 0.0 && fatRatec <= 0.22 {
			help = "当前比较偏瘦"
		} else if fatRatec > 0.22 && fatRatec <= 0.29 {
			help = "标准身材"
		} else if fatRatec > 0.29 && fatRatec <= 0.36 {
			help = "当前偏重"
		} else if fatRatec > 0.36 && fatRatec <= 0.41 {
			help = "当前肥胖"
		} else {
			help = "严重肥胖"
		}
	default:
		help = "小于18岁不统计该数据"
	}
	return help
}