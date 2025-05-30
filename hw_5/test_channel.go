package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}

	c <- sum
}

func SetToChannel() {
	s := []int{7, 2, 8, -9, 4, 0, 1}
	c := make(chan int)

	channel1 := s[:len(s)/2]
	fmt.Println("Канал 1:", channel1)
	go sum(channel1, c)

	channel2 := s[len(s)/2:]
	fmt.Println("Канал 2:", channel2)
	go sum(channel2, c)

	x, y := <-c, <-c
	fmt.Println("\nКанал 1[Сумма]:", x)
	fmt.Println("Канал 2[Сумма]:", y)

	fmt.Println("Итого:", x+y)
}
