package main

import (
	"errors"
	"fmt"
)

func main() {
	var i *int
	a := 10
	i = &a

	fmt.Println(i)
	fmt.Println(*i)

	// addOrCounter := &Counter{
	// 	hits: 0,
	// }

	counter := Counter{
		hits: 0,
	}

	fmt.Println(counter.hits)

	increment(&counter)

	fmt.Println(counter.hits)

	var iAuth IAuth
	auth := Auth{}
	iAuth = auth

	iAuth.login("admin@admin.com", "password")
	iAuth.logout()

	fmt.Println(errors.New("failed to load data"))

}

type Counter struct {
	hits int
}

func increment(counter *Counter) {
	counter.hits++
}

type Auth struct {
}

// login implements IAuth
func (Auth) login(string, string) bool {
	fmt.Println("login...")
	return true
}

// logout implements IAuth
func (Auth) logout() {
	fmt.Println("logout...")
}

type IAuth interface {
	login(string, string) bool
	logout()
}

func Error() string {
	return "somethings went wrong"
}
