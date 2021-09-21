package main

import (
	"fmt"
	"go-test-hex/repo"
	"log"
)

func main() {
	r := repo.ConnectDB{
		Username: "myUsername",
		Password: "myPassword",
	}

	x := repo.NewRepo(r)

	fmt.Println(x)

	info, err := x.GetInfo() //from interface
	if err != nil {
		log.Println(err)
	}

	fmt.Println(*info)

}
