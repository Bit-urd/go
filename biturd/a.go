package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Slice 操作示例:")

	var slice []int
	slice = append(slice, 1)
	slice = append(slice, 2, 3, 4)
	fmt.Println("切片内容:", slice)

	fmt.Println("切片的第一个元素:", slice[0])
	fmt.Println("切片的最后一个元素:", slice[len(slice)-1])

	slice[1] = 100
	fmt.Println("修改后的切片:", slice)

	fmt.Println("\nMap 操作示例:")

	m := make(map[string]int)
	m["apple"] = 5
	m["banana"] = 3
	m["cherry"] = 10

	fmt.Println("苹果的数量:", m["apple"])
	fmt.Println("香蕉的数量:", m["banana"])

	if value, exists := m["orange"]; exists {
		fmt.Println("橙子的数量:", value)
	} else {
		fmt.Println("橙子不存在")
	}

	fmt.Println("\nChannel 操作示例:")

	ch := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	for value := range ch {
		fmt.Println("从 channel 接收到的数据:", value)
	}

	fmt.Println("\nGoroutine 操作示例:")

	go func() {
		for i := 1; i <= 3; i++ {
			fmt.Println("Goroutine 正在执行任务:", i)
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(4 * time.Second)
	fmt.Println("主 goroutine 完成")
}
