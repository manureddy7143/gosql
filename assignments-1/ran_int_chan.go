package main 
  
import ( 
    "fmt"
    "math/rand"
) 
var a chan int
 
  
func main() { 
	a = make(chan int)
	go randint()
	b:= <-a
	fmt.Println(b)
}
func randint(){
	
	a <- rand.Intn(200)
}