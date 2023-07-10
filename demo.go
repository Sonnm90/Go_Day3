package main

import (
	"errors"
	"fmt"
	"image/color"
)

func main() {
	//a := 1
	//b := 2
	//fmt.Println(Add(a, b))
	//fmt.Println(Add1(a, b))
	//deferFc()
	//fmt.Println(yes())

	//Slice trong fc
	//arr := [3]int{1, 2, 3}
	//first(arr)
	//fmt.Println(arr)
	//second(arr[0:])
	//fmt.Println(arr)

	//anonymous function
	//DoStuff()
	//
	//DoStuff = func() {
	//	fmt.Println("Doing stuff!")
	//}
	//DoStuff()
	//
	//DoStuff = func() {
	//	fmt.Println("Doing other stuff.")
	//}
	//DoStuff()
	//
	//fmt.Println(Cal(1, 2))
	//Cal = func(a, b int) int {
	//	return a + b
	//}
	//fmt.Println(Cal(1, 2))
	//
	//Cal = func(a, b int) int {
	//	return a * b
	//}
	//fmt.Println(Cal(1, 2))

	//Closures
	//n := 0
	//counter := func() int {
	//	n += 1
	//	return n
	//}
	//fmt.Println(counter())
	//fmt.Println(counter())

	//counter := newCounter()
	//fmt.Println(counter())
	//fmt.Println(counter())
	//fmt.Println(counter())
	//fmt.Println(counter())

	//Pointer
	var a = 100
	var p = &a

	fmt.Println("a = ", a)
	fmt.Println("p = ", p)
	fmt.Println("*p = ", *p)
	*p = 200
	fmt.Println(a)

	//Struct

	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}

	var p1 Person // // tất cả các trường được khởi tạo với giá trị 0
	fmt.Println("Person1: ", p1)
	p2 := Person{"Sơn", "Nguyễn", 33}
	fmt.Println("Person2: ", p2)
	fmt.Println("person2---first name", p2.FirstName)

	p3 := Person{
		FirstName: "John",
		LastName:  "Snow",
		Age:       45,
	}
	fmt.Println("Person3: ", p3)
	ps := &p3
	fmt.Println("Pointer of person3", ps)
	ps.Age = 55
	fmt.Println("Person3: ", p3)

	p4 := Person{FirstName: "Robert"}
	fmt.Println("Person4: ", p4)
	p4.Age = 30
	fmt.Println("Person4: ", p4)

	p5 := new(Person)
	p5.FirstName = "Sơn"
	p5.LastName = "Nguyễn"
	p5.Age = 33

	fmt.Println(p5)
	fmt.Println(*p5 == p2)

	//Method

	s := Square{4}
	r := Rectangle{3, 4}
	sArea := s.Acreage()
	rArea := r.Acreage()
	fmt.Println("area of Square: ", sArea)
	fmt.Println("area of Rectangle: ", rArea)

	point := Point{3, 4}
	fmt.Println("Point  = ", point)

	//point.Translate(7, 9)
	//fmt.Println("After Translate, point = ", point)

	TranslateFunc(&point, 7, 9)
	fmt.Println("After Translate, p = ", point)

	myStr := MyString("OLLEH")
	fmt.Println(myStr.reverse())

	var cp ColoredPoint
	cp.X = 1
	cp.Point.Y = 2
	fmt.Println(cp)

	//Interface
	duck := Duck{}
	duck.Scream()

	//Error
	// tạo error
	//myErr := &MyError{}
	myErr := errors.New("something unexpected happened")
	// in ra thông báo lỗi
	fmt.Println(myErr)
	fmt.Printf("Type of myErr is %T \n", myErr)
	fmt.Printf("Value of myErr is %#v \n", myErr)
}

type MyError struct{}

func (myErr *MyError) Error() string {
	return "Something unexpected happened!"
}

type Animal interface {
	Scream()
}
type Duck struct {
}

func (duck Duck) Scream() {
	fmt.Println("The duck quacks")
}

func TranslateFunc(p *Point, dx, dy float64) {
	p.X = p.X + dx
	p.Y = p.Y + dy
}

type MyString string

func (myStr MyString) reverse() string {
	s := string(myStr)
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

type Point struct {
	X, Y float64
}
type ColoredPoint struct {
	// thuộc tính ẩn danh
	Point

	// thuộc tính bình thường
	Color color.RGBA
}

func (p *Point) Translate(dx, dy float64) {
	p.X = p.X + dx
	p.Y = p.Y + dy
}

// hàm được đặt tên
func Add(a, b int) int {
	return a + b
}

// hàm ẩn danh
var Add1 = func(a, b int) int {
	return a + b
}

func deferFc() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
func yes() (text string) {
	defer func() {
		text = "no"
	}()
	return "yes"
}

func first(x [3]int) {
	for i := range x {
		x[i] += 1
	}
}
func second(x []int) {
	for i := range x {
		x[i] += 1
	}
}

var DoStuff func() = func() {
	// Do stuff
}

var Cal = func(a, b int) int {
	return a
}

func newCounter() func() int {
	n := 0
	return func() int {
		n += 1
		return n
	}
}

type Square struct {
	X float64
}
type Rectangle struct {
	X, Y float64
}

func (s Square) Acreage() float64 {
	return s.X * s.X
}
func (r Rectangle) Acreage() float64 {
	return r.X * r.Y
}
