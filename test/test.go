package main

import "fmt"

func Hello() string {
	return "Hello World"
}

// go의 경우 함수 타입 선언 맨 뒤에

func main() {
	fmt.Println(Hello())
}
