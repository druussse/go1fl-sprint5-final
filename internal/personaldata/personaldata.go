package personaldata

import "fmt"

// Personal содержит сведения о пользователе.
type Personal struct {
	Name   string
	Weight float64
	Height float64
}

// Print выводит личные данные пользователя.
func (p Personal) Print() {
	fmt.Printf("Имя: %s\nВес: %.2f\nРост: %.2f\n", p.Name, p.Weight, p.Height)
}
