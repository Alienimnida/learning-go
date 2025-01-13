package main

import "fmt"

func main() {
	/*var intArr [3]int32
	intArr[1] = 123
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])*/

	intArr := [...]int32{1, 2, 3}
	fmt.Println(intArr)

	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("\nThe length is %v with capacity %v", len(intSlice), cap(intSlice))

	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Println(intSlice3)

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap2["Adam"])

	var age, ok = myMap2["Jason"]
	delete(myMap2, "Adam")
	if ok {
		fmt.Printf("The age is %v", age)
	} else {
		fmt.Println("Invalid Name")
	}

	for name := range myMap2 {
		fmt.Printf("Name: %v\n", name)
	}

	for i, v := range intArr {
		fmt.Printf("Index: %v, Value: %v \n", i, v)
	}

	var i int = 0
	for i < 10 {
		fmt.Println(i)
		i = i + 1
	}

}
