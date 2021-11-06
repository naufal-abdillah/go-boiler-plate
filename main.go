package main

import (
	"fmt"
	"go-boiler-plate/cmd"
)

func main() {
	fmt.Println("Printing in main")
	cmd.Run()
	var model_s car
	model_s.brand = "tesla"
	model_s.price = 100000
	var redmi phone
	redmi.brand = "mi"
	redmi.price = 100
	redmi.ram = 4

	// model_s car
	model_s.Printprice()
	redmi.Printprice()
	redmi.Buy()

	var birthdaypresent device
	birthdaypresent = phone{"samsung", 1000, 16}
	birthdaypresent.Printprice()

}

type device interface {
	Printprice()
}
type car struct {
	brand string
	price int
}
type phone struct {
	brand string
	price int
	ram   int
}

func (c car) Printprice() {
	fmt.Printf("Price of the %s is $%d\n", c.brand, c.price)
}
func (p phone) Printprice() {
	fmt.Printf("Price of the %s is $%d\n", p.brand, p.price)
}
func (p phone) Buy() {
	fmt.Printf("Do you want to by the %s\n", p.brand)
}
