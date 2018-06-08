package contributor

import (
	"fmt"
)

//go:generate go run ../gen.go

func PrintContributors() {
	for _, c := range Contributors {
		fmt.Println(c)
	}
}
