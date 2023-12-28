package main

import (
	"errors"
	"fmt"
)

func rectArea(len, width float64) (area float64, err error) {
	if len < 0 || width < 0 {
		err = errors.New("cannot use negative number to calculate area")
		return
	}

	area = len * width
	return
}

func main() {
	a, err := rectArea(10, 15)
	fmt.Println("error returned: ", err)
	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println("area: ", a)
	}
}
