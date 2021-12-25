package main

import "fmt"

type IF interface {
	getName() string
}

type Human struct {
	firstName, lastName string
}

type Plane struct {
	vendor, model string
}

type Car struct {
	factory, model string
}

func (h *Human) getName() string {
	return h.firstName + "," + h.lastName
}

func (p *Plane) getName() string {
	return p.model + "," + p.vendor
}

func (c *Car) getName() string {
	return c.factory + "," + c.model
}

func main() {
	interfaces := []IF{}
	h := new(Human)
	h.firstName = "Z"
	h.lastName = "P"
	interfaces = append(interfaces, h)

	c := new(Car)
	c.factory = "bmw"
	c.model = "a"
	interfaces = append(interfaces, c)

	for _, f := range interfaces {
		fmt.Println(f.getName())
	}

	p := Plane{}
	p.vendor = "abc";
	p.model = "a";
	fmt.Println(p.getName())
}