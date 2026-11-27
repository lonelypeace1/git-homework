package main

import (
	"fmt"
	"study/feature_postgres/simple_connection"
	"study/repeater"
)

func main() {
	fmt.Println("Hello, GIT!")
	repeater.PrintTimes("Я учу гит!", 5)
	simple_connection.CheckConnection()
}
