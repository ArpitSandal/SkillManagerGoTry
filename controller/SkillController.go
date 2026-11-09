package controller

import (
	"context"
	"log/slog"
	"net/http"

	service "github.com/ArpitSandal/SkillManagerGo/Service"
	"github.com/ArpitSandal/SkillManagerGo/models/dto"
	"github.com/gin-gonic/gin"
)

type ISkillController interface { // it will have all the method declarations of the apis
	SearchEmployeesBySkillController(context *gin.Context)
	AddEmployeeController(context *gin.Context)
	RemoveEmployeeController(context *gin.Context)
}

type SkillController struct { // add dependency of the service class here
	skillService service.ISkillService
}

func NewSkillController(skillService service.ISkillService) ISkillController {
	return SkillController{
		skillService: skillService,
	}
}

func (controller SkillController) SearchEmployeesBySkillController(context *gin.Context) {

	val := struct {
		Skills []string
	}{} // the reflecting binding from incoming json will only happen if we capitalize the variable inside the struct

	// responsebody binding
	if err := context.ShouldBindJSON(&val); err != nil {
		buildErrorResponse(http.StatusBadRequest, context, "Error Binding", err)
		return
	}

	employee, err := controller.skillService.SearchEmployeesBySkillService(val.Skills)

	if err != nil {
		buildErrorResponse(http.StatusInternalServerError, context, "Error", err)
		return
	}
	context.JSON(http.StatusOK, gin.H{"Employees": employee})
}

func (controller SkillController) AddEmployeeController(ginContext *gin.Context) {

	ctx := context.Background()
	employeeWithSkill := dto.EmployeeDTO{}

	if err := ginContext.ShouldBindJSON(&employeeWithSkill); err != nil {
		buildErrorResponse(http.StatusBadRequest, ginContext, "Error Binding", err)
		return
	}

	slog.Info("employee with skill", "skill", employeeWithSkill)

	addedEmployee, err := controller.skillService.AddEmployeeService(ctx, employeeWithSkill)

	if err != nil {
		buildErrorResponse(http.StatusInternalServerError, ginContext, "Error", err)
		return
	}

	ginContext.JSON(http.StatusOK, addedEmployee)
}

func (controller SkillController) RemoveEmployeeController(context *gin.Context) {

}

func buildErrorResponse(statusCode int, context *gin.Context, msg string, err error) {
	context.JSON(statusCode, gin.H{msg: err.Error()})
}
