package handler

import (
	"fmt"
	"go-boiler-plate/services"
)

func HandlerPrint() {
	fmt.Println("Printing in Handler")
	services.Services()
}
