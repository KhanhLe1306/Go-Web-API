package main

import "fmt"

func main(){
	fmt.Println(mySqrt(4))
}

func mySqrt(x int) int {
    i := 0
    for {
        if i * i == x {
            return i
        }else if i * i > x {
            return i - 1
        }
        i++
    }
}