package routes

import (
	"Gin_task/dao"
	"Gin_task/util"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// 创建路由引擎
	router := gin.Default()
	util.ConnectDB()
	// 创建 TaskController 实例
	taskDAO := dao.NewTaskDAO(util.DB) // 根据你的实际需求创建 TaskDAO 实例
	taskController := dao.NewTaskController(taskDAO)

	// 定义路由和处理函数
	router.POST("/tasks", taskController.CreateTask)
	router.PUT("/tasks/:id", taskController.UpdateTask)
	router.GET("/tasks/:id", taskController.GetTask)
	router.GET("/tasks", taskController.GetAllTasks)
	router.DELETE("/tasks/:id", taskController.DeleteTask)

	return router
}
