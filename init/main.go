package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"pluseid.io/invitation/cmd"
)

func main() {
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
