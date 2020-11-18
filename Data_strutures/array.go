package main
import (
 "fmt"
)
func main() {
 var name [50]string //array of 50 string elements
 fmt.Println(name)
 name[43] = "Welcome to Go Programming"
 fmt.Println(name)
 var x [100]string
 fmt.Println("Length of arr x:", len(x)) //calculate length of array
 z := [5]int{43, 78, 54, 23, 22}
 fmt.Println(z) 
 for i := 65; i <= 122; i++ {
 x[i-65] = string(i)

 }
 fmt.Println(x)
}