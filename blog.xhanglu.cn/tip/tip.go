package tip

import (
	"fmt"
)

func CheckErr(err error) {
	if err != nil {
		fmt.Println("ddddddddddd")
		panic(err)
	}
}
