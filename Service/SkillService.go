package service

import (
	"context"

	"github.com/ArpitSandal/SkillManagerGo/models"
	"github.com/ArpitSandal/SkillManagerGo/models/dto"
	"github.com/ArpitSandal/SkillManagerGo/repository"
	"github.com/gin-gonic/gin"
)

type ISkillService interface {
	SearchEmployeesBySkillService(skills []string) ([]models.Employee, error)
	AddEmployeeService(ctx context.Context, addedEmployeeDTO dto.EmployeeDTO) (models.Employee, error)
	RemoveEmployeeService(context *gin.Context)
}

type SkillService struct {
	SkillRepository repository.ISkillRepository
}

func NewSkillService(skillRepository repository.ISkillRepository) SkillService {
	return SkillService{
		SkillRepository: skillRepository,
	}
}

func (service SkillService) SearchEmployeesBySkillService(skills []string) ([]models.Employee, error) {
	var employees []models.Employee

	employees, err := service.SkillRepository.SearchEmployeesBySkillRepository(skills)

	return employees, err
}

func (service SkillService) AddEmployeeService(ctx context.Context, addedEmployeeDTO dto.EmployeeDTO) (models.Employee, error) {
	var employee models.Employee
	employee.ToEntity(addedEmployeeDTO)

	employee, err := service.SkillRepository.AddEmployeeRepository(ctx, employee)
	return employee, err
}
func (service SkillService) RemoveEmployeeService(context *gin.Context) {

}
