package main

import (
	"fmt"
	"getting-to-go/config"
	"getting-to-go/model"
	"getting-to-go/server"
	"log"
)

func main() {
	c, err := config.LoadConfig("./config/config.yaml")
	if err != nil {
		fmt.Println(err)
	}

	err = model.Connect(c.Database.Host, c.Database.Port, c.Database.User, c.Database.Password, c.Database.Name)
	if err != nil {
		fmt.Println(err)
	}

	model.RunMigrations()

	s, err := server.NewServer(server.NewConfig(c))
	if err != nil {
		log.Panic("Failed To Start Server:", err)
	}

	s.Run()
}
