package main

import "fmt"

type gasEngine struct {
	mpg       uint8
	gallons   uint8
	ownerInfo owner
	int
}

type owner struct {
	name string
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it!")
	} else {
		fmt.Println("You can't make it!")
	}
}

func main() {
	/*var myEngine gasEngine = gasEngine{25, 15, owner{"Alex"}, 10}
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.ownerInfo.name, myEngine.int)
	fmt.Printf("Total miles left in tank: %v", myEngine.milesLeft())*/
	var myEngine electricEngine = electricEngine{25, 15}
	canMakeIt(myEngine, 50)
}
