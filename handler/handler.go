package handler

import (
	"fmt"
	"strconv"

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

	t := &mytask.Task{}
	ctx.BindJSON(&t)

	t1 := mytask.Add(t)

	ctx.JSON(200, t1)
}

func EditTask(ctx *gin.Context) {
	fmt.Println("Edit 000")

	taskIDStr := ctx.Param("taskID")
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		ctx.JSON(400, getError(101, "Invalid ID: "+taskIDStr))
		return
	}

	t := &mytask.Task{}

	ctx.BindJSON(t)

	t1 := mytask.Edit(taskID, t)
	if t1 == nil {
		ctx.JSON(400, getError(102, "Task not found: "+taskIDStr))
		return
	}

	ctx.JSON(200, t1)
}

func getError(id int, desc string) map[string]interface{} {
	return map[string]interface{}{
		"error": gin.H{
			"id":          id,
			"description": desc,
		},
	}
}
