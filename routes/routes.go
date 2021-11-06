package routes

import (
	"fmt"
	"go-boiler-plate/handler"
)

func Routes() {
	fmt.Println("Printing in Routes")
	handler.HandlerPrint()
}
func main() {
	fmt.Println("Printing in Routes main")
}
