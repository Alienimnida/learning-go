package main

import (
	"fmt"
	"strings"
)

func main() {
	//rune is alias for int32
	var myString = []rune("résumé")
	var indexed = myString[0]
	fmt.Printf("%v, %T\n", indexed, indexed)

	for i, v := range myString {
		fmt.Println(i, v)
	}
	fmt.Printf("\nThe length of 'myString' is %v", len(myString))

	var myRune = 'a'
	fmt.Printf("\nmyRune= %v", myRune)

	var strSlice = []string{"a", "b", "c"}
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var catStr = strBuilder.String()
	fmt.Printf("\n%v", catStr)
}

//Output:
// 114, int32
// 0 114
// 1 233
// 2 115
// 3 117
// 4 109
// 5 233

// The length of 'myString' is 6
