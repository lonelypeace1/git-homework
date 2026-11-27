package repeater

import "fmt"

func PrintTimes(text string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(text)
	}
}
