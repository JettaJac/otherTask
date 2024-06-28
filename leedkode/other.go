package main

import (
// "fmt"
// "sync"
// "time"
)

func notDublicate(num []int) []int {
	res := []int{}
	tmp := make(map[int]int)
	for i := 0; i < len(num); i++ {
		_, ok := tmp[num[i]]
		if !ok {
			tmp[num[i]] = 1
			res = append(res, num[i])
		}
	}
	return res
}

// func main() {
// 	num := []int{3, 2, 1, 1, 0, 4, 5, 2, 0}
// 	fmt.Println(notDublicate(num), " == ", "[3, 2, 1, 0, 4, 5]")

// }

func worker(input, output chan int) {
	for id := range input {
		output <- id * id
	}
	close(output)
}

// Создайте две горутины, чтобы числа из одного канала читались по мере поступления, возводились в квадрат и результат записывался во второй канал.

// func main() {
// 	wg := sync.WaitGroup{}
// 	num := []int{3, 2, 1, 1, 0, 4, 5, 2, 0}
// 	input := make(chan int, 10)
// 	output := make(chan int, 10)
// 	go func() {

// 		for i := range num {
// 			wg.Add(1)
// 			input <- i
// 		}
// 		close(input)
// 	}()
// 	// fmt.Println(<-input)
// 	go worker(input, output, &wg)

// 	fmt.Println("Res: ")
// 	go func() {
// 		for el := range output {
// 			fmt.Println(el)
// 			wg.Done()
// 		}
// 	}()
// 	wg.Wait()
// 	// time.Sleep(time.Second * 10)
// }
