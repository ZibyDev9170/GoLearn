package main

import "fmt"

func main() {
	err := doSomething()
	if err != nil {
		fmt.Println("ошибка:", err)
	}

	// Частая ловушка с err в if-блоках
	if result, err := riskyOperation(); err != nil {
		// здесь err — это НОВАЯ переменная, только в этом блоке
		fmt.Println("ошибка:", err)
	} else {
		fmt.Println("результат:", result)
	}
}

func doSomething() error           { return nil }
func riskyOperation() (int, error) { return 42, nil }
