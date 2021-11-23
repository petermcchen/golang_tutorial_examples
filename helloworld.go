package main

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"runtime"
	"time"
	"unsafe"

	"golang.org/x/tour/pic"
	"golang.org/x/tour/reader"
	"gopkg.in/yaml.v2"
)

func main() {
	fmt.Println("Hello World!")
	test1() // Output
	test2() // Slices
	test3() // Switch
	test4() // Function
	test5() // Recursion Function
	test6() // Map
	test7() // Nested Map
	test8() // Map Operations
	fmt.Println("square root of 2: ", Sqrt(2))
	fmt.Println("square root of 10: ", Sqrt(10))
	test9()  // Switch
	test10() // Slices
	test11() // Pic implement
	test12() // float64
	test13() // implement error() interface.
	fmt.Println(SqrtFunc(2))
	fmt.Println(SqrtFunc(-2))
	test14()    // Readers
	test15()    // ShowImage
	deadlock1() // channel congestion deadlock
	deadlock2() // read empty channel deadlock
	deadlock3() // buffered channel overflow deadlock
	deadlock4() // for range deadlock
	test16()
}

func test16() {
	//var ss string
	ss := "Hello World! Hello World! Hello World! Hello World!"
	var ii int
	var ff float64
	var aa [1]uint32
	var es struct{}
	var xx = make([]struct{}, 10000)

	//ss = "Hello World!"
	fmt.Println("unsafe.Sizeof test: size in bytes")
	fmt.Println("ss: ", unsafe.Sizeof(ss))
	fmt.Printf("ss: %v, %T, %p\n", ss, ss, &ss)
	fmt.Println("ii: ", unsafe.Sizeof(ii))
	fmt.Println("ff: ", unsafe.Sizeof(ff))
	fmt.Println("aa: ", unsafe.Sizeof(aa))
	fmt.Println("es: ", unsafe.Sizeof(es))
	fmt.Println("xx: ", unsafe.Sizeof(xx))
}

func deadlock4() {
	fmt.Println("for range deadlock case:")
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
	for v := range ch {
		fmt.Println(v)
	}
}

func write(ch chan int) {
	for i := 1; i <= 4; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}

func deadlock3() {
	fmt.Println("buffered channel overflow case:")
	ch := make(chan int, 3)
	//ch <- 1
	//ch <- 2
	//ch <- 3
	//ch <- 4
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
	go write(ch)
	time.Sleep(100 * time.Millisecond)
	for v := range ch {
		fmt.Println("read value", v, "from ch")
		time.Sleep(100 * time.Millisecond)
	}
}

func deadlock2() {
	fmt.Println("read empty channel case:")
	ch := make(chan int, 3)
	//fmt.Println(<-ch)
	select {
	case v := <-ch:
		fmt.Println(v)
	default:
		fmt.Println("chan no data")
	}
}

func deadlock1() {
	fmt.Println("channel congestion case:")
	ch := make(chan int)
	//ch <- 1
	go func() {
		ch <- 1
	}()
	fmt.Println(<-ch)
}

type Image struct {
	width  int
	height int
}

func (m Image) ColorModel() color.Model {
	return color.Alpha16Model
}
func (m Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, m.width, m.height)
}

func (m Image) At(x int, y int) color.Color {
	v := uint8((x * x) + (y * y))
	return color.RGBA{v, v, 255, 255}
}

func test15() { // ShowImage
	m := Image{64, 64}
	pic.ShowImage(m)
}

type MyReader struct{} // Empty struct as method receiver

func (r MyReader) Read(b []byte) (int, error) {
	b[0] = 'A'
	return 1, nil
}

func test14() {
	reader.Validate(MyReader{})
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"It is not working",
	}
}

func test13() {
	err := run()
	if err != nil {
		fmt.Println(err)
	}
}

func test12() {
	f := float64(-math.Sqrt2)
	fmt.Println(f)
}

func Pic(dx, dy int) [][]uint8 {
	fmt.Printf("dx: %d, dy: %d\n", dx, dy)
	pixel := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		pixel[i] = make([]uint8, dx)
	}
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			pixel[i][j] = uint8((i * i) + (j * j))
		}
	}
	//fmt.Printf("%v", pixel)
	return pixel
}

func test11() {
	pic.Show(Pic)
}

// Output test
func test1() {
	var student1 string = "Peter"
	var student2 = "Grace"
	x := 2

	fmt.Println("\n++++++++++++++++")
	fmt.Println("test1")
	fmt.Println("student1: " + student1)
	fmt.Println(student2)
	fmt.Println(x)

	fmt.Printf("student2 has value: %v and type: %T\n", student2, student2)
}

// Slices Test
func test2() {
	var name string
	var id int
	var flag bool
	var cars = [4]string{"Volve", "BMW", "VW", "Mazda"}

	fmt.Println("\n++++++++++++++++")
	fmt.Println("test2")
	fmt.Println(name)
	fmt.Println(id)
	fmt.Println(flag)

	fmt.Print(cars, "\n")

	myslice := []string{"Sun", "Moon", "Earth", "Ocean"}
	fmt.Println("myslice len = ", len(myslice))
	fmt.Println("myslice cap = ", cap(myslice))
	fmt.Println(myslice)

	myslice2 := make([]int, 4, 7)
	fmt.Printf("myslice2 = %v\n", myslice2)
	fmt.Printf("length = %d\n", len(myslice2))
	fmt.Printf("capacity = %d\n", cap(myslice2))

	myslice3 := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("myslice3 = %v\n", myslice3)
	fmt.Printf("length = %d\n", len(myslice3))
	fmt.Printf("capacity = %d\n", cap(myslice3))
	myslice3 = append(myslice3, 8, 9)
	fmt.Printf("new myslice3 = %v\n", myslice3)
	fmt.Printf("length = %d\n", len(myslice3))
	fmt.Printf("capacity = %d\n", cap(myslice3))

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	fmt.Printf("mumbers = %v\n", numbers)
	fmt.Printf("length = %d\n", len(numbers))
	fmt.Printf("capacity = %d\n", cap(numbers))
	newnumber := numbers[0 : len(numbers)-10]
	numberCopy := make([]int, 5)
	copy(numberCopy, newnumber)
	fmt.Printf("newnumber = %v\n", newnumber)
	fmt.Printf("length = %d\n", len(newnumber))
	fmt.Printf("capacity = %d\n", cap(newnumber))

	fmt.Printf("numberCopy = %v\n", numberCopy)
	fmt.Printf("length = %d\n", len(numberCopy))
	fmt.Printf("capacity = %d\n", cap(numberCopy))
}

// Switch Test
func test3() {
	x := 20
	y := 10
	if x > y {
		fmt.Println("x is greater than y")
	} else {
		fmt.Println("x is less than or equal to y")
	}

	day := 5
	switch day {
	case 1:
		fmt.Println("Today is Monday")
	case 2:
		fmt.Println("Today is Tuesday")
	case 3:
		fmt.Println("Today is Wednesday")
	case 4:
		fmt.Println("Today is Thursday")
	case 5:
		fmt.Println("Today is Friday")
	case 6:
		fmt.Println("Today is Saturday")
	case 7:
		fmt.Println("Today is Sunday")
	default:
		fmt.Println("Not a wek day")
	}
}

func myFunction(x int, y int) int {
	return x + y
}

// Function Test
func test4() {
	var myResult = 10 * myFunction(2, 3)
	fmt.Printf("myResult: %d\n", myResult)
}

func factorial_recursion(x float64) (y float64) {
	if x > 0 {
		y = x * factorial_recursion(x-1)
	} else {
		y = 1
	}
	return
}

// Recursion Function Test
func test5() {
	fmt.Println("5! = ", factorial_recursion(5))
}

// Map Test (Empty)
func test6() {
	var a = make(map[string]string)

	a["year"] = "2021"
	a["month"] = "Oct"
	a["name"] = "chen"
	fmt.Printf("Map a : %v\n", a)

	// Cannot assign to nil...
	//var b map[string]string
	//b["course"] = "CS"
	//b["degree"] = "A-"
	//b["name"] = "chen"
	//fmt.Printf("Map b : %v\n", b)
}

// Nested Map Test
type T struct {
	codename string
	street   string
}

func test7() {
	var x = map[string]map[string]string{}

	x["fruits"] = make(map[string]string)
	x["colors"] = make(map[string]string)
	x["fruits"]["a"] = "apple"
	x["fruits"]["b"] = "banana"
	x["colors"]["r"] = "red"
	x["colors"]["b"] = "blue"
	fmt.Println(x)

	names := []string{"Alan", "Rob", "Janet", "David"}
	m := make(map[string]map[string]T, len(names))
	m["uid"] = make(map[string]T)
	for _, name := range names {
		m["uid"][name] = T{codename: "Taiwan", street: "Taipei"} // TODO...
	}
	fmt.Println(m)
	y, _ := yaml.Marshal(&m)
	fmt.Println(string(y))
}

// Map Operations
func test8() {
	a := make(map[string]string)

	a["brand"] = "VW"
	a["model"] = "Passat variant"
	a["year"] = "2019"
	fmt.Println(a)
	a["color"] = "white"
	fmt.Println(a)
	delete(a, "year")
	fmt.Println(a)
}

// Implement Sqrt
func Sqrt(x float64) float64 {
	z := 1.0
	n := 0.0
	loop := 20
	for {
		//fmt.Println(z)
		n = z - (z*z-x)/(2*x)
		z = n
		loop -= 1
		if loop < 0 {
			break
		}
	}
	return (z)
}

// Implement SqrtFunc to return error when parameter is negative value.
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number: ", float64(e))
}

func SqrtFunc(x float64) (float64, error) {
	if x < 0.0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	n := 0.0
	loop := 20
	for {
		//fmt.Println(z)
		n = z - (z*z-x)/(2*x)
		z = n
		loop -= 1
		if loop < 0 {
			break
		}
	}
	return z, nil

}

// Switch Test
func test9() {
	fmt.Print("Go runs on ")
	os := runtime.GOOS
	switch os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

	t := time.Now()
	switch {
	case t.Local().Hour() < 12:
		fmt.Println("Good morning.")
	case t.Local().Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

// Create Slices
func test10() {
	a := make([]int, 5)
	a[3] = 9
	printSlice("a", a)
	b := make([]int, 0, 5)
	printSlice("b", b)
	c := b[:3]
	c[2] = 9
	printSlice("c", c)
	d := c[1:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
