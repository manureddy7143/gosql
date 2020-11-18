package main
import (
	"fmt"
)

func main() {
	primes := []int{1, 3, 5, 7, 9, 11}
	fmt.Printf("%T\n", primes) //type []int
	fmt.Println(primes)        //[1 3 5 7 9 11]
	fmt.Println(primes[2:4])   // [5 7] 3rd 4th element

	var info []byte                  //Slice using var
	elements := []int{3, 5, 2, 6, 2} //Slice using shorthand notation
	arrays := make([]int, 0, 3)
	fmt.Println(arrays)      //Creating slice with make( )
	//String slice
	str := []string{"Amazing", "Brilliant", "Good One", "Excellent"}
	for n := range elements {
		fmt.Println("Elements", n)
	}
	for i := 0; i < 80; i++ {
		arrays = append(arrays, i) //append i elements to arrays
		fmt.Println("Len:", len(arrays), "Capacity:",
			cap(arrays), "Value: ", arrays[i])
	}

	fmt.Println("Info elements:", info)
	for i, vals := range str {
		fmt.Println("Values", i, vals)
	}
	//Slice of Slice
	OutIn := make([][]int, 0, 3)
	for i := 0; i < 3; i++ {
		out := make([]int, 0, 4)
		for j := 0; j < 4; j++ {
			out = append(out, j)
		}
		OutIn = append(OutIn, out)
	}
	fmt.Println(OutIn)

}
