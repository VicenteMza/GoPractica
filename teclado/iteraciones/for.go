package iteraciones

import (
	"fmt"
)

func Iterar() {
	for i := 100; i > 1; i-- {
		if i == 6 {
			continue
		}
		fmt.Println(i)
	}
}
