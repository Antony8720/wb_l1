// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

package tasks

import "fmt"

// Создаем родительскую структуру
type Human struct {
	name string
	surname string
}

// Встраиваем родительскую структуру в Action
type Action struct {
	Human
}

func (h *Human) sayHello() string{
	return fmt.Sprintf("Hello I am %v %v!", h.name, h.surname)
}

func task1() {
	action := Action{Human{name: "Ivan", surname: "Ivanov"}}
	fmt.Println(action.sayHello())
}