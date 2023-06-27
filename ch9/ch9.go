package main

import (
	"fmt"
	"reflect"
	"strconv"
)

/**
구조체 심화

https://stackoverflow.com/questions/33593545/difference-between-struct-vs-struct
https://5kyc1ad.tistory.com/228

struct 내에서는 함수를 구현할 수 없다. 대신 자바처럼 메서드를 구현하고 싶을 땐, 리시버(receiver)를 이용해서 어느 구조체의 메서드인지를 정의 할 수 있다.
리시버는 Value 리시버와 포인터 리시버 두 가지 타입이 있다.


https://mingrammer.com/translation-go-and-oop/
public private이 없다.
대신 멤버변수를 카멜케이스로 명시하면 외부패키지에서 접근할수 없다.



*/

type Person struct {
	name string
	age  int
}

type Team struct {
	title  string //default value는 바로 할당 불가능.. 따로 함수를 만들어야 되네.
	person Person //컴포지션
}

func newPerson(name string) *Person {
	p := Person{name: name}
	p.age = 42
	return &p
}

func (p Person) StayTheSame() {
	p.name = "개똥"
	p.age = 7
}

func (p *Person) Mutate() {
	p.name = "영수"
	p.age = 10
}

func main() {

	//구조체

	fmt.Println(Person{"Bob", 20})
	fmt.Println(Person{name: "Alice", age: 30})
	fmt.Println(Person{name: "Fred"})
	fmt.Println(&Person{name: "Ann", age: 40})
	fmt.Println(newPerson("Jon"))

	s := Person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)

	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)

	var person Person
	var p *Person = &person
	p.name = "sean"

	fmt.Println(person, p)

	p1 := &Person{}
	p2 := new(Person)

	fmt.Println(p1)
	fmt.Println(p2)

	person.StayTheSame()
	fmt.Println(person)
	person.Mutate()
	fmt.Println(person)

	person.name = "하이트"
	fmt.Println(person)
	p1.name = "메시"
	fmt.Println(p1)

	team := &Team{
		"바르셀로나",
		*p1,
	}

	fmt.Println(team)

	// 이것이 차이점인가..? null일 수 있는 값을 나타내야 하는 경우 포인터를 사용?
	p1 = nil
	//person = nil

	var person2 = NewPerson2()
	fmt.Println(person2)

	//https://stackoverflow.com/questions/24512112/how-to-print-struct-variables-in-console
	fmt.Printf("%+v\n", person2)

}

type Person2 struct {
	Name    string `default:"John Doe"`
	Age     int    `default:"30"`
	Address string `default:"123 Main St"`
}

func NewPerson2() Person2 {
	p := Person2{}
	setDefaults(&p)
	return p
}

func setDefaults(p *Person2) {
	// Iterate over the fields of the Person struct using reflection
	// and set the default value for each field if the field is not provided
	// by the caller of the constructor function.
	for i := 0; i < reflect.TypeOf(*p).NumField(); i++ {
		field := reflect.TypeOf(*p).Field(i)
		if value, ok := field.Tag.Lookup("default"); ok {
			switch field.Type.Kind() {
			case reflect.String:
				if p.Name == "" {
					p.Name = value
				}
			case reflect.Int:
				if p.Age == 0 {
					if intValue, err := strconv.Atoi(value); err == nil {
						p.Age = intValue
					}
				}
			}
		}
	}
}
