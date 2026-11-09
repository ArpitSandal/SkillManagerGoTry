package repository

import (
	"context"
	"maps"
	"slices"

	"github.com/ArpitSandal/SkillManagerGo/models"
	"github.com/ArpitSandal/SkillManagerGo/utility"
)

// using squirrel for dynamic search queries

func (repo SkillRepository) SearchEmployeesBySkillRepository(skills []string) ([]models.Employee, error) {
	const funcName = "SearchEmployeesBySkillRepository"

	sqlStr, args, err := repo.queryBuilder.SearchEmployeesBySkillQueryBuilder(skills)

	if err != nil {
		utility.PrintError(funcName, "Failed to generate the query", err)
		return nil, err
	}

	rows, err := repo.pgxPool.Query(context.Background(), sqlStr, args...)

	if err != nil {
		utility.PrintError(funcName, "Failed to execute Query", err)
		return nil, err
	}

	defer rows.Close()

	employeeMap := make(map[int]models.Employee)

	// return pgx.CollectRows(rows,pgx.RowToStructByName[models.Employee]),nil
	for rows.Next() {

		var employee models.Employee
		var skill models.Skill

		if err := rows.Scan(&employee.Id, &employee.EmployeeName, &employee.Age, &skill.Category, &skill.SkillName, &skill.Id); err != nil {
			utility.PrintError(funcName, "Failed to map the response", err)
			return nil, err
		}

		if _, ok := employeeMap[employee.Id]; !ok {

			employeeMap[employee.Id] = models.Employee{
				Id:           employee.Id,
				Age:          employee.Age,
				EmployeeName: employee.EmployeeName,
			}
		}
		employeeWithSkill := employeeMap[employee.Id]
		employeeWithSkill.Skills = append(employeeWithSkill.Skills, skill)
		employeeMap[employee.Id] = employeeWithSkill

	}

	if err := rows.Err(); err != nil {
		utility.PrintError(funcName, "Failed to read employee rows", err)
		return nil, err
	}

	return slices.Collect(maps.Values(employeeMap)), nil

}
