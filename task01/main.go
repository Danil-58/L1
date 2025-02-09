package main

import "fmt"

// Структура Human
type Human struct{}

/*
Так как в языке Go нет прямого наследования, как в других ООП языках,
его можно заменить встраиванием структуры в структуру, таким образом
можно вызывать из Action методы Human
*/
type Action struct {
	Human
}

func (h Human) DoSomething() {
	fmt.Println("Human mthod DoSomething")
}

func main() {
	action := Action{}
	action.DoSomething()
}