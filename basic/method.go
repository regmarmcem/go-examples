package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u *User) Aging() {
	u.Age++
}

func (u User) AgingButNothingHappen() {
	u.Age++
}

func main() {
	u := &User{
		Name: "Richard",
		Age:  33,
	}

	u.Aging()
	fmt.Println(u.Age)

	u.AgingButNothingHappen()
	fmt.Println(u.Age)

}
