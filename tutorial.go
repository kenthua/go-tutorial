package main

// running through tutorial basics
// https://tour.golang.org/basics/1

import (
	"fmt"
	//"github.com/kenthua/stringutil"
	"time"
	"math/rand"
	"math"
	"math/cmplx"
)

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // naked return since we know it's already the x and y as int, not ideal
}

func swap(x ,y string) (string, string) {
	return y, x // both variables are returned
}

func add(x int, y int) int {
	return x + y
}

func addShort(x, y int) int {
	return x+y
}

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

var c, python, java bool
var d, f int = 1, 2

// variable declaration box style, like import
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const Pi = 3.14

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	//fmt.Println(stringutil.Reverse("!oG ,olleH"))
	fmt.Println("Hello", time.Now())

	// the seed, via time, was generated when r was defined
	fmt.Println("My favorite number is", r.Int63()) 

	fmt.Println(math.Pi)
	fmt.Println(add(1,2))
	fmt.Println(addShort(1,2))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))

	var i int

	java = true

	// this isn't really necessary, since python is already defined
	python := true
	
	fmt.Println(i, c, python, java)
	
	// it just knows what type this is
	var str = "mystring" 

	fmt.Println(d, f, str)

	// no need to specify variable when doing short delcaration :=, this can't be done outside of function, need to use var outside
	k := 3

	fmt.Println(k)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	
	myInt := 42
	// need to explicitly cast
	var myFloat float64 = float64(myInt) + .1
	// another way to cast with out var
	myUInt := uint(myFloat)

	fmt.Println(myInt, myFloat, myUInt)

	v := 42 
	fmt.Println("v is of type %T", v)

	//Constants can be character, string, boolean, or numeric values.
	//Constants cannot be declared using the := syntax.
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
