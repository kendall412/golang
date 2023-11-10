package main

import "fmt"

type Engineer struct {
	Name string
	Age  int
}

func changeValue(str *string) {
	*str = "changed!"
}

func (e *Engineer) changeValueb() {
	e.Name = "Shinoo"
}

func changeValue2(str string) {
	str = "changed!"
}

func main() {
	toChange := "hello"

	changeValue(&toChange)
	fmt.Println(toChange)

	eng := &Engineer{
		Name: "March",
		Age:  42,
	}

	eng.changeValueb()
	fmt.Println(eng.Name)

}
