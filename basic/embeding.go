package main

import "fmt"

type Chip struct {
	Number int
}

type Card struct {
	string
	Chip
	Number int
}

func (c *Chip) Scan() {
	fmt.Println(c.Number)
}

func (c *Card) Echo() {
	fmt.Println(c.string)
	fmt.Println(c.Number)
}

func main() {
	c := Card{
		string: "Credit",
		Chip: Chip{
			Number: 424242,
		},
		Number: 545454,
	}

	c.Scan()
	c.Echo()
}
