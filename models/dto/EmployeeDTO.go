package dto

type EmployeeDTO struct {
	EmployeeName string   `json:"employee_name"`
	Age          int      `json:"age"`
	Skills       []SkillDTO `json:"skills"` // it means don't map it with database
}
