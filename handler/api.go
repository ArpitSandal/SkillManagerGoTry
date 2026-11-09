package main

import (
	"github.com/ArpitSandal/SkillManagerGo/controller"
	"github.com/gin-gonic/gin"
)

// this will handle the gin router and route paths

type IRouter interface {
	InitializeRouter() *gin.Engine
}

type Router struct {
	controller controller.ISkillController
}

func (router Router) InitializeRouter() *gin.Engine {

	engine := gin.Default()

	/*
		requestbody: {
		skills: ["java"]
		}
	*/
	engine.POST("/searchEmployeeBySkill", router.controller.SearchEmployeesBySkillController)

	/*

		requestbody{
		employee{
			skills[],
			...employee fields
			}
		}
	*/
	engine.POST("/addEmployee", router.controller.AddEmployeeController)

	engine.DELETE("/deleteEmployee", router.controller.RemoveEmployeeController)

	engine.Run(":8080")
	return engine
}
