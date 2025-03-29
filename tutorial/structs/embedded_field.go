/*
Anonymous embedded fields: An anonymous embedded field is a struct field that doesn’t have an explicit field name. Structs that have an anonymous field obtain all of the methods and properties of the nested struct. These methods and properties are called “promoted” methods and properties. Anonymous embedded fields can also be accessed by the type name directly. Methods and properties from the embedded struct are “shadowed” and the methods and properties on the struct you are calling will be used. You can still access the shadowed methods and properties by specifically accessing the embedded field and then calling the method or property.

Anonymous embedded fields allow methods and properties from the embedded struct to automatically get promoted to the nesting struct, but there’s a possibility of those methods and properties getting shadowed if the nesting struct declares methods of properties of the same name.

Explicit fields directly associate a field name with the nested struct, so that accessing the nested struct’s methods and properties are unambiguous. In production code, it’s often better to use explicit fields over anonymous embedded fields to make the code more readable and not surprise yourself with shadowed fields.

https://kevin-yang.medium.com/golang-embedded-structs-b9d20aadea84
*/

package main

import (
	// alias
	f "fmt"
)

type Hur struct {
	/* "Person" is an embedded field from "Person" struct
		Any fields or methods declared on an embedded field are promoted to the containing struct and can be invoked directly on it

	type Inner struct {
	    X int
	}

	type Outer struct {
	    Inner
	    X int
	}

		https://learning.oreilly.com/library/view/learning-go-2nd/9781098139285/ch07.html#:-:text=Any%20fields%20or%20methods%20dec,be%20invoked%20directly%20on%20it
	*/
	Person
	family_position string
}

type Hur2 struct {
	person          Person
	family_position string
}

type Person struct {
	fname string
	lname string
	age   int
}

func (p *Person) getOlder(age int) {
	p.age = p.age + age
}

func (p Person) statement(name string) {
	f.Println("Hello " + name)
}

func (h *Hur) changeFamPos(newpos string) {
	h.family_position = newpos
}

func (h *Hur2) changeFamPos2(newpos string) {
	h.family_position = newpos
}

func main() {
	danny := &Hur{
		// embedded fields
		Person: Person{
			fname: "Danny",
			lname: "Hur",
			age:   51,
		},
		family_position: "Father",
	}

	march := &Hur2{
		// explicit fields
		person: Person{
			fname: "March",
			lname: "Hur",
			age:   42,
		},
		family_position: "Mother",
	}

	// embedded field: can access Person struct without directly
	f.Println(danny.age)
	danny.getOlder(20)
	f.Println(danny.age)
	f.Println(danny.family_position)
	danny.changeFamPos("Husband")
	f.Println(danny.family_position)

	// explicit field: must use "person"
	f.Println(march.person.age)
	march.person.getOlder(3)
	f.Println(march.person.age)
	f.Println(march.family_position)
	march.changeFamPos2("Wife")
	f.Println(march.family_position)

}
