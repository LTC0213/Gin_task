package main

import "Gin_task/routes"

func main() {
	router := routes.SetupRouter()

	// 启动应用
	router.Run(":8080")
}
