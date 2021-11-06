package cmd

import (
	"fmt"
	"go-boiler-plate/routes"
)

func Run() {
	fmt.Println("Printing in CMD")
	routes.Routes()
}

func main() {
	fmt.Println("Printing in CMD in main")
}
