package main

import (
	"fmt"
	"log"
	"math/rand"
	"reflect"
	"sort"
	"sync"
	"time"
)

var (
	BMI   []sync.Map
	rank  []float64
	end   chan int
	mutex sync.Mutex
)

func partitionDouble(arr []float64, startIndex int, endIndex int) int {
	left, right := startIndex, endIndex
	pivot := arr[startIndex]

	for left != right {
		for left < right && arr[right] > pivot {
			right--
		}
		for left < right && arr[left] <= pivot {
			left++
		}
		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}
	arr[startIndex] = arr[left]
	arr[left] = pivot
	return left
}

func DoubleCycleMethod(arr []float64, startIndex int, endIndex int) {
	if startIndex >= endIndex {
		return
	}

	pivotIndex := partitionDouble(arr, startIndex, endIndex)

	DoubleCycleMethod(arr, startIndex, pivotIndex-1)
	DoubleCycleMethod(arr, pivotIndex+1, endIndex)
}

func ReturnRank(bmi float64) (index int) {
	defer mutex.Unlock()
	mutex.Lock()
	DoubleCycleMethod(rank, 0, 999)
	index = sort.SearchFloat64s(rank, bmi)
	return
}

func UpdataBMI(index int) float64 {
	rand.Seed(time.Now().UnixNano())
	base := -0.2 + rand.Float64()*(0.2-(-0.2))
	bmi, _ := BMI[index-1].Load("BMI")
	switch bmi.(type) {
	case float64:
		value := reflect.ValueOf(bmi)
		bmi := value.Float() + base
		BMI[index-1].Store("BMI", bmi)
		rank[index-1] = bmi
		return bmi
	default:
		log.Fatal("数据初始化,类型不正确.")
		return 0.0
	}
}

func init() {
	end = make(chan int, 0)
	BMI = make([]sync.Map, 0, 1000)
	rank = make([]float64, 1000)
	for i := 1; i <= 1000; i++ {
		data := sync.Map{}
		data.Store("name", fmt.Sprintf("%d", i))
		base := 0.75 + rand.Float64()*(0.75-0.55)
		data.Store("BMI", base)
		BMI = append(BMI, data)
		rank[i-1] = 0.75
	}
}

func main() {
	for i := 1; i <= 1000; i++ {
		i := i
		go func() {
			defer func() {
				end <- 1
			}()
			bmi := UpdataBMI(i)
			index := ReturnRank(bmi)
			fmt.Printf("当前用户%d的BMI排名为: %d.\n", i, index+1)
		}()
	}

	for {
		select {
		case <-end:
			if _, ok := <-end; !ok {
				fmt.Println("system out.")
				return
			}
		default:
		}
	}
}
