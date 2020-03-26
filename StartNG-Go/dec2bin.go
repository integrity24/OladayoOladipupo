package main

import (
 "fmt"
 "strconv"
)
func main() {
 var decimal int64
 fmt.Print("Enter a Number:")
 fmt.Scanln(&decimal)
 output := strconv.FormatInt(decimal, 2)
 fmt.Print("The Binary Representation of this number is: ", output)

}