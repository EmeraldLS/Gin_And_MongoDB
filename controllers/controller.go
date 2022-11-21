package controllers

import (
	"net/http"

	"github.com/EmeraldLS/Gin_And_MongoDB/helpers"
	"github.com/EmeraldLS/Gin_And_MongoDB/models"
	"github.com/gin-gonic/gin"
)

func AddOneStudent(ctx *gin.Context) {
	var student models.Student
	if err := ctx.ShouldBindJSON(&student); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
	}
	helpers.AddoneStudentToDB(student)
	ctx.JSON(http.StatusOK, gin.H{
		"student": student,
	})
}

func GetSingleStudent(ctx *gin.Context) {
	studentId := ctx.Param("id")
	student := helpers.GetSingleStudentFromDB(studentId)
	ctx.JSON(http.StatusOK, gin.H{
		"student": student,
	})
}

func GetAllStudent(ctx *gin.Context) {
	student := helpers.GetAllStudentFromDB()
	ctx.JSON(http.StatusOK, gin.H{
		"students": student,
	})
}

func RemoveOneStudent(ctx *gin.Context) {
	studentId := ctx.Param("id")
	helpers.RemoveOneStudentFromDB(studentId)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Student has been removed successfully",
	})
}

func RemoveAllStudent(ctx *gin.Context) {
	helpers.RemoveAllStudentFromDB()
	message := "All students has been removed successfully"
	ctx.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}
