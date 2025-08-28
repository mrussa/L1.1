package main

import "fmt"

var _ Greeter = (*Human)(nil)
var _ Greeter = (*Action)(nil)
var _ Greeter = (*Logger)(nil)

var _ Walker = (*Human)(nil)
var _ Walker = (*Action)(nil)
var _ Walker = (*Logger)(nil)

type Greeter interface {
	Greet() string
}

type Walker interface {
	Walk(steps int)
}

type Human struct {
	Name string
	Age  int
}

func (h Human) Greet() string {
	return fmt.Sprintf("Hi, I am %s", h.Name)
}

func (h *Human) Birthday() {
	h.Age++
}

func (h *Human) Walk(steps int) {
	fmt.Printf("%s walked %d steps\n", h.Name, steps)
}

type Logger struct {
}

func (l Logger) Greet() string {
	return "Hi, its me"
}

func (l Logger) Walk(steps int) {
	fmt.Printf("Walked %d steps\n", steps)
}

func UseWalker(w Walker, steps int) {
	w.Walk(steps)
}

type Action struct {
	*Human
	*Logger
	Kind string
}

// затеняем
func (a *Action) Greet() string {
	fmt.Printf("Action '%s': \n", a.Kind)
	return fmt.Sprintf("Hi, I am %s", a.Name)
}

func (a *Action) Walk(steps int) {
	fmt.Printf("Action '%s': \n", a.Kind)
	a.Human.Walk(steps)

}

func main() {

	h := &Human{Name: "Dan", Age: 30}
	a := Action{Human: h, Logger: &Logger{}, Kind: "Study"}

	a.Name = "Daniel"
	fmt.Println(a.Greet())
	fmt.Println(a.Human.Greet())
	fmt.Println("Name via Human: ", h.Name)
	fmt.Println("Name via Action: ", a.Name)
	fmt.Println("Age via Action (before): ", a.Age)
	fmt.Println("Age via Human  (before): ", h.Age)
	a.Birthday()
	fmt.Println("Age via Action (after birthday): ", a.Age)
	fmt.Println("Age via Human  (after birthday): ", h.Age)
	h.Walk(300)
	a.Walk(1200)
	UseWalker(h, 500)
	UseWalker(&a, 800)
	UseWalker(a.Logger, 100)
	fmt.Println(a.Logger.Greet())
}
