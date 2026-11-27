package repeater

import "fmt"

func PrintFiveTimes(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
	}
}
