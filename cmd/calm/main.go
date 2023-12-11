package main

import (
	"fmt"
	"os"

	"github.com/flat35hd99/play-sqlite-go/repositories"
	"github.com/flat35hd99/play-sqlite-go/service"
)

func main() {
	repo := repositories.NewTaskRepository()
	service := service.NewTaskService(repo)

	t, err := service.Create("Hello, world!")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("task id: %d\n", t.ID)
	fmt.Printf("task description: %s\n", t.Description)
}
