package main

import "fmt"

type Logger interface {
	log()
}

type User struct {
	Name string
}

func (u *User) log() {
	if u == nil {
		fmt.Println("<nil>")
	} else {
		fmt.Println(u.Name)
	}
}

func main() {
	var user Logger
	fmt.Printf("%v, %T\n", user, user)

	user = &User{"Naveed"}

	user.log()

	fmt.Printf("(%v, %T)\n", user, user)
}
