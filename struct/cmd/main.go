package main

import "fmt"

// 構造体
type Person struct {
	name string
	age  int
}

// 構造体の関数 歳をとる
// 非ポインタなので、値参照となり、ここだけしか影響が及ばない
func (p Person) ToAge() {
	p.age++
}

// 構造体の関数 歳をとる
// ポインタなので、生成した構造体の値に影響がある
func (p *Person) ToAgeP() {
	p.age++
}

func main() {
	// 構造体を生成
	person := Person{}
	person.name = "person taro"
	person.age = 20

	// 構造体の関数を実行して、ageの値を出力
	fmt.Println("init age: " + fmt.Sprint(person.age))
	person.ToAge()
	fmt.Println("exec person.ToAge(). age: " + fmt.Sprint(person.age))
	person.ToAgeP()
	fmt.Println("exec person.ToAgeP(). age: " + fmt.Sprint(person.age))
	person.ToAge()
	fmt.Println("exec person.ToAge(). age: " + fmt.Sprint(person.age))
	person.ToAgeP()
	fmt.Println("exec person.ToAgeP(). age: " + fmt.Sprint(person.age))

}
