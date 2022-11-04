package main

import (
	"errors"
	"fmt"
)

func main() {

	val, err := divide(20, 0)

	if err != nil {
		fmt.Println(val, err)

		fmt.Println(errors.Is(err, &DevError{
			// 20, 0,
		}))
	}

}

func divide(lhs, rhs int) (int, error) {
	if rhs == 0 {
		return 0, &DevError{
			lhs, rhs,
		}
		// return 0, errors.New("cannot divide by zero")

	} else {
		return rhs / lhs, nil
	}
}

type DevError struct {
	a, b int
}

func (d *DevError) Error() string {
	return fmt.Sprintf("cannot divide by zero:%d/%d", d.a, d.b)
}
