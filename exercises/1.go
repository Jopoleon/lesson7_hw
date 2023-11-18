package main

import "fmt"

type CustomInt int

type Age int

type CustumIntS struct {
	Age   Age
	MyInt int
}

func (c *CustumIntS) IncrementS() {
	c.MyInt += 1
}

func (c *CustomInt) Increment123() {
	*c += 1
}

// Задача 1
func task1() {

	var nn CustumIntS = CustumIntS{MyInt: 14}

	fmt.Println("before: ", nn)
	nn.IncrementS()
	fmt.Println("after:  ", nn)

	var num CustomInt = 5

	// var foo CustomInt = 222

	num.Increment123()
	num.Increment123()
	num.Increment123()
	num.Increment123()
	num.Increment123()
	num.Increment123()

	// foo.Increment123()

	fmt.Println("Задача 1:", num) // Что будет выведено?

}

type Data struct {
	num int
}

func modify(d *Data) {
	d.num = 10
}

// Задача 2
func task2() {

	data := Data{5}
	modify(&data)
	fmt.Println("Задача 2:", data.num) // Что будет выведено?
}

type Item struct {
	Value int
}

func changeItem(item Item) {
	item.Value = 2
}

// Задача 3
func task3() {

	item := Item{Value: 1}
	changeItem(item)
	fmt.Println("Задача 3:", item.Value) // Что будет выведено?
}

type Person struct {
	Name string
	Age  *int
}

func ageIncrement(age *int) {
	*age++
}

// Задача 4
func task4() {

	age := 30
	person := Person{Name: "Alice", Age: &age}
	ageIncrement(person.Age)
	fmt.Println("Задача 4:", *person.Age) // Что будет выведено?
}

type Numbers struct {
	A, B int
}

func (n Numbers) Swap() {
	n.A, n.B = n.B, n.A
}

// Задача 5
func task5() {

	numbers := Numbers{A: 1, B: 2}
	numbers.Swap()
	fmt.Println("Задача 5: A =", numbers.A, "B =", numbers.B) // Что будет выведено?
}

// ---------------------------------------------------
type Counter struct {
	Value int
}

func (c *Counter) Increment() {
	c.Value++
}

func (c Counter) Decrement() {
	c.Value--
}

// Задача 6
func task6() {

	ccc := &Counter{33}
	// var cc *Counter = ccc

	ccc.Decrement()
	fmt.Println("ccc.Decrement()", ccc.Value)
	ccc.Increment()
	fmt.Println("ccc.Increment()", ccc.Value)

	counter := Counter{33}
	counter.Increment()
	counter.Decrement()
	fmt.Println("Задача 6:", counter.Value) // Что будет выведено?
}

func main() {
	// task1()
	// task2()
	// task3()
	// task4()
	// task5()
	hardTasks()
}
