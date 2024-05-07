package hello_world

import (
	"fmt"
	"math"
	"reflect"
)

type Abser interface {
	Abs() float64
}

type Abser2 interface {
	Abs() float64
}

func TryInterface(){
	var myFloat MyFloat = -12
	a := myFloat
	fmt.Println("myFloat.Abs()", a.Abs()) 
	fmt.Println("TypeOf(a.Abs())", reflect.TypeOf(a.Abs()))

	v := Vertex{4, 5}

	fmt.Println("v.Abs()", v.Abs())
}

type MyFloat float64

func (myFloat MyFloat) Abs() float64 {
	if myFloat < 0 {
		return float64(-myFloat)
	}
	return float64(myFloat)
}

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}