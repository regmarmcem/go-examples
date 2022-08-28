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

type Hex int

func (h Hex) String() string {
	return fmt.Sprintf("%x", int(h))
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

	fv := Hex(1024).String
	fmt.Println(fv())

	test := Hex(1024)
	fmt.Println(test)

	fe := Hex.String
	fmt.Println(fe(1024))

}
