package main

import (
	// "encoding/json"
	"fmt"
	"sync"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

// func main() {
// 	c := SafeCounter{v: make(map[string]int)}
// 	for i := 0; i < 1000; i++ {
// 		go c.Inc("somekey")
// 	}

// 	time.Sleep(time.Second)
// 	fmt.Println(c.Value("somekey"))
// }

type Point struct {
	X int
	Y int
}
var GlobalString string
var i int = 1
func main() {
	// data := []byte(`{"x":1, "y":2}`)
	// var point Point 
	// if err := json.Unmarshal(data, &point); err != nil {
	// 	fmt.Println("error: ", err)
	// } else {
	// 	fmt.Println(point)
	// }
	// // str := "Khanh Le"
	// // bytes := []byte(str)
	// // fmt.Printf("len(%v): ", str)
	// // fmt.Println( str)
	// // fmt.Println(bytes)"
	// if GlobalString == "" {
	// 	fmt.Println("[" + GlobalString + "]")
	// }

	// var any interface{} = 12
	any := 12
	fmt.Println(any.(int))
	// switch t := any.(type){
	// case int: 
	// 	fmt.Printf("It's an int: %v", t)
	// case float64: 
	// 	fmt.Printf("It's an float64: %v", t)
	// case string: 
	// 	fmt.Printf("It's an string: %v", t)
	// }

	// b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	// var f interface{}
	// err := json.Unmarshal(b, &f)
	// if err == nil {
	// 	m := f.(map[string]interface{})
	// 	for k, v := range m {
	// 		switch vv := v.(type) {
	// 		case string:
	// 			fmt.Println(k, "is string", vv)
	// 		case float64:
	// 			fmt.Println(k, "is float64", vv)
	// 		case []interface{}:
	// 			fmt.Println(k, "is an array:")
	// 			for i, u := range vv {
	// 				fmt.Println(i, u)
	// 			}
	// 		default:
	// 			fmt.Println(k, "is of a type I don't know how to handle")
	// 		}
	// 	}
	// }
}
