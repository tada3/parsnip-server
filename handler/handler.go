package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tada3/parsnip-server/mytask"
)

func GetTasks(ctx *gin.Context) {

	fmt.Println("Tasks 000")

	tasks := mytask.GetAll()

	ctx.JSON(200, tasks)
}

func AddTask(ctx *gin.Context) {
	fmt.Println("Tasks 000")

	var t mytask.Task
	ctx.BindJSON(&t)

	t1 := mytask.Add(t)

	ctx.JSON(200, t1)
}
