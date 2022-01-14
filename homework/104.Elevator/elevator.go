package main

/*
	此版本电梯没有地下层 v0.1
*/

import (
	"fmt"
	//"sort"
	"time"
)

type elevator struct {
	currentFloor int //当前楼层
	direction int		//当前状态
}


//上升   状态   目标
func (el *elevator) rise(dir int,targetFloor int)(num int, err error){

	if dir != 1 {
		err = fmt.Errorf("当前状态不是上升(1)/下降(0)")
		return
	}

	if targetFloor < el.currentFloor{
		err = fmt.Errorf("目标楼层有误")
		return
	}else if targetFloor > el.currentFloor{
		for i:=el.currentFloor; i<targetFloor; i++ {
			time.Sleep(1000)
			el.currentFloor ++
			fmt.Println("到达楼层",el.currentFloor)
		}
	}else if targetFloor == el.currentFloor {
		el.currentFloor  = targetFloor
		fmt.Println("到达目标楼层",el.currentFloor)
	}
	openDoor()
	closeDoor()

	return el.currentFloor,nil
}


//下降   状态   目标
func (el *elevator) decline(dir int,targetFloor int)(num int, err error)  {

	if dir != 0 {
		err = fmt.Errorf("当前状态不是上升(1)/下降(0)")
		return
	}

	if el.currentFloor < targetFloor{
		err = fmt.Errorf("目标楼层有误")
		return
	}else if el.currentFloor == targetFloor {
		err = fmt.Errorf("目标楼层就是当前该楼层")
		return
	}

	if targetFloor < el.currentFloor{
		for i:=el.currentFloor; i>targetFloor; i-- {
			el.currentFloor --
			time.Sleep(1000)
			fmt.Println("1.当前楼层",el.currentFloor)
		}
	}
	//el.currentFloor = targetFloor


	openDoor()
	closeDoor()
	return el.currentFloor,nil

}



func (el1 *elevator)handle(event int)  {


/*
	firstValue := "first"
	targetFloor = 4
	Targetfloor[targetFloor] = firstValue

	//第3问
	nextValue = "second"
	targetFloor = 5
	Targetfloor[targetFloor] = nextValue
*/
	el1.currentFloor = 1
	for {
		if listingEvents(1) {
			break
		}else {
			listingEvents(1)
		}
	}


}

func listingEvents(event int) bool {

	var el1 elevator
	var lastFloor,targetFloor int = 0,0
	var nextValue string

	//楼层 是否第一次

	fmt.Println("当前楼层",el1.currentFloor)

	Targetfloor:=map[int]string{
		1: "",
		2: "",
		3: "",
		4: "",
		5: "",
	}


	for k, v := range Targetfloor{
		//第二问
		if k == 2 {
			fmt.Println("中途按了3楼")
			nextValue = "zero"
			targetFloor = 3
			Targetfloor[targetFloor] = nextValue
		}
		if v == "zero"{
			lastFloor,_ = el1.rise(event,k)
			Targetfloor[k] = ""
		} else if v == "first"{
			lastFloor,_ = el1.rise(event,k)
			Targetfloor[k] = ""
		}else if v == "second"{
			lastFloor,_ = el1.rise(event,k)
			Targetfloor[k] = ""
		}
		/*			if k == 5 {

					nextValue = "zero"
					targetFloor = 2
					fmt.Printf("%v楼，按了%v楼\n",el1.currentFloor,targetFloor)
					Targetfloor[targetFloor] = nextValue
				}*/
	}

	for k,_:= range Targetfloor{
		if Targetfloor[k] !=""{
			fmt.Println("开始下降----------")
			lastFloor,_ = el1.decline(event,k)
			Targetfloor[k] = ""
		}
	}

	fmt.Println("当前楼层", lastFloor)


	return false

}


func openDoor()  {
	fmt.Println("门正在打开.......")
	time.Sleep(1000)
}
func closeDoor()  {
	fmt.Println("门正在关闭.......")
	time.Sleep(1000)
}
