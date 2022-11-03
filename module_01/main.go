package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Go Fundamentals...")

	fmt.Println(time.Now())

	// Data types

	var name string = "Khadga Bahadur Shrestha"
	fmt.Println(name)

	var age int = 25
	fmt.Println(age)

	companyName, location := "PortPro Nepal", "Lalitpur Nepal"
	fmt.Println(companyName, location)

	var (
		studentName = "Nishuka Shrestha"
		userAge     = 23
		role        = "admin"
		isActive    = true
	)

	fmt.Println(studentName, userAge, role, isActive)

	// Basic functions
	sayHello()

	sayHi("Nishuka Shrestha")

	fmt.Println(add(10, 20))

	fmt.Println(swap("Hello", "World"))

	// arrays
	users := [2]string{
		"Khadga shrestha", "Nishuka shrestha",
	}

	fmt.Println(users)
	fmt.Println(len(users))
	fmt.Println(users[0])

	// range in array

	for i, name := range users {
		fmt.Println(i, name)
	}

	if age < 20 {
		fmt.Println("you are eligible to perform this action.")
	} else {
		fmt.Println("sorry you are not eligible to perform this action.")
	}

	// switch case

	switch role {
	case "admin":
		fmt.Println("you are admin.")

	default:
		fmt.Println("you are not admin.")
	}

	// structure
	loggedInUser := user{
		name: "Khadga shrestha",
		role: "Flutter Developer",
		age:  25,
		skills: []string{
			"flutter", "firebaase",
		},
	}

	fmt.Println(loggedInUser)
	fmt.Println(loggedInUser.skills)

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println("sum:", sum)

	// pointers
	var p *int
	a := 10
	p = &a
	fmt.Println(p)
	fmt.Println(*p)

	// maps
	var userDoc map[string]string = make(map[string]string)
	userDoc["name"] = "Khadga Shrestha"
	userDoc["role"] = "Mobile Application Developer"

	fmt.Println(userDoc)
	fmt.Println(userDoc["name"])

	v, ok := userDoc["name"]
	if ok {
		fmt.Println("my name is ", v)
	}

	// methods

	result := loggedInUser.getInfo()
	fmt.Println(result)

	// var iBase IAuth
	iBase := loggedInUser
	fmt.Println("is logged in", iBase.isLoggedIn())
}

func sayHello() {
	fmt.Println("hey...")
}

func sayHi(name string) {
	fmt.Printf("Hey...%s\n", name)
}

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

type user struct {
	name   string
	role   string
	age    int
	skills []string
}

// isLoggedIn implements IAuth
func (user) isLoggedIn() bool {
	return true
}

func (auser user) getInfo() string {
	return fmt.Sprintf("My name is %s, I am %d years old and I am a %s", auser.name, auser.age, auser.role)
}

type IAuth interface {
	isLoggedIn() bool
}
