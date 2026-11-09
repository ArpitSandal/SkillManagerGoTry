package models

import (
	"github.com/ArpitSandal/SkillManagerGo/models/dto"
)

type Employee struct {
	Id           int     `db:"id"`
	EmployeeName string  `db:"employee_name"`
	Age          int     `db:"age"`
	Skills       []Skill `db:"-"` // it means don't map it with database
}

func (e *Employee) ToEntity(employeeDTO dto.EmployeeDTO) {

	e.Age = employeeDTO.Age
	e.EmployeeName = employeeDTO.EmployeeName
	for _, skill := range employeeDTO.Skills {
		e.Skills = append(e.Skills, Skill{
			SkillName: skill.SkillName,
			Category:  skill.Category,
		})
	}

}
