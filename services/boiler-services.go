package services

import (
	"fmt"
	repo "go-boiler-plate/repositories"
)

func Services() {
	fmt.Println("Printing in Services")
	repo.Repo()
}
