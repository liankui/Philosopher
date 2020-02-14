package main

import (
	"github.com/yueekee/Philosopher/GinHello/initRouter"
)

func main() {
	router := initRouter.SetupRouter()
	router.Run()
}