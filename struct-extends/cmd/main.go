package main

import "fmt"

// 継承元になる構造体Person
type Person struct {
	name string
	age  int
}

// Personの関数 歩く
func (p Person) Walk() {
	fmt.Println(p.name + " is Walk.")
}

// 継承元のPerson構造体を継承した構造体Member
type Member struct {
	// Personの継承とするため、構造に埋め込む
	Person
	// Memberの独自プロパティ
	id int
}

// Personの関数 歩く を継承先のMemberでオーバーライド
func (m Member) Walk() {
	fmt.Println(m.name + "(Member) is Walk.")
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

	// ---

	// 継承元と継承先で、同じ関数を実行する
	// memberの方では、オーバーライドした処理が実行される
	person.Walk()
	member.Walk()

	fmt.Println("End.")
}
