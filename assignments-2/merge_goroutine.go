package main

import (
    "crypto/rand"
    "fmt"
    "os"
    "strconv"
)

var (
    nums    []byte //The slice of numbers we want to sort
    numVals int    = -1
)

//User can optionally add a parameter that determines how many random numbers will be sorted
//If none are provided, 100 will be used
func main() {
    if len(os.Args) >= 2 {
        numVals, _ = strconv.Atoi(os.Args[1])
    } else {
        numVals = 100
    }
    nums = initSlice()
    ms := make(chan []byte)
    go mergeSort(nums, ms)
    nums = <-ms
    for _, value := range nums {
        fmt.Printf("%d\n", value)
    }
}

func initSlice() []byte {
     vals := make([]byte, numVals)
     _, err := rand.Read(vals)
     if err != nil {
        panic(err)
     }
     return vals
}

func mergeSort(arr []byte, ms chan []byte) {
    if len(arr) <= 1 { //base case
        ms <- arr
        return
    }
    leftMS := make(chan []byte)
    go mergeSort(arr[:len(arr)/2], leftMS)
    rightMS := make(chan []byte)
    go mergeSort(arr[len(arr)/2:], rightMS)
    left, right := <-leftMS, <-rightMS
    sortArr := make([]byte, len(arr))
    lIndex, rIndex := 0, 0
    //Combine both sides
    for lIndex < len(left) && rIndex < len(right) {
        leftLeast := left[lIndex] <= right[rIndex]
        if leftLeast {
            sortArr[lIndex+rIndex] = left[lIndex]
            lIndex++
        } else {
            sortArr[lIndex+rIndex] = right[rIndex]
            rIndex++
        }
    }
    //Fill in whichever 'pile' is empty
    if lIndex < len(left) {
        for ; lIndex < len(left); lIndex++ {
            sortArr[lIndex+rIndex] = left[lIndex]
        }
    }
    if rIndex < len(right) {
        for ; rIndex < len(right); rIndex++ {
            sortArr[lIndex+rIndex] = right[rIndex]
        }
    }
    ms <- sortArr
}