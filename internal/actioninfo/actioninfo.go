package actioninfo

import (
	"fmt"
	"log"
	"strings"
)

// DataParser представляет действие, которое можно разобрать и описать.
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

// Info выводит данные по всем корректным записям об активности.
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			log.Println(err)
			continue
		}
		info, err := dp.ActionInfo()
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Print(info)
		if !strings.HasSuffix(info, "\n") {
			fmt.Println()
		}
	}
}
