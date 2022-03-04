package main

import "fmt"

// 継承元になる構造体Person
type Person struct {
	name string
	age  int
}

// 継承元のPerson構造体を継承した構造体Member
type Member struct {
	// Personの継承とするため、構造に埋め込む
	Person
	// Memberの独自プロパティ
	id int
}

func main() {
	fmt.Println("Start.")

	// 継承元構造体を生成して出力
	person := Person{name: "Person"}
	fmt.Println(person)

	// 継承先構造体を生成して出力
	member := Member{Person: Person{name: "Member"}}
	fmt.Println(member)

	// 生成した継承元構造体を与える形で、継承先構造体を生成して出力
	newMember := Member{Person: person}
	fmt.Println(newMember)

	fmt.Println("End.")
}
