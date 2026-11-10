package repository

import (
	"context"

	"github.com/ArpitSandal/SkillManagerGo/models"
	sqlcRepo "github.com/ArpitSandal/SkillManagerGo/postgresql/sqlc"
	"github.com/ArpitSandal/SkillManagerGo/utility"
	"github.com/jackc/pgx/v5/pgtype"
)

func (repo SkillRepository) AddEmployeeRepository(ctx context.Context, addedEmployee models.Employee) (models.Employee, error) {
	const funcName = "AddEmployeeRepository"

	// begin transaction
	tx, err := repo.pgxPool.Begin(ctx)
	if err != nil {
		utility.PrintError(funcName, "Failed to begin transaction", err)
		return models.Employee{}, err
	}

	defer tx.Rollback(ctx)

	// build new query object through which we can execute query methods in interface of sqlc
	queries := sqlcRepo.New(tx)

	id, err := queries.AddEmployee(ctx, sqlcRepo.AddEmployeeParams{
		EmployeeName: addedEmployee.EmployeeName,
		Age:          int32(addedEmployee.Age),
	})

	if err != nil {
		utility.PrintError(funcName, "Failed to add employee", err)
		return models.Employee{}, err
	}

	addedEmployee.Id = int(id)

	addedSkills := make([]models.Skill, 0)
	for _, skill := range addedEmployee.Skills {
		addedSkill, err := queries.AddEmployeeSkill(ctx, sqlcRepo.AddEmployeeSkillParams{
			SkillName:  skill.SkillName,
			Category:   pgtype.Text{String: skill.Category, Valid: true},  // tells the postgres it's a non null value
			EmployeeID: pgtype.Int4{Int32: id, Valid: true},
		})
		if err != nil {
			utility.PrintError(funcName, "Failed to add Skill", err)
			return models.Employee{}, err
		}

		addedSkills = append(addedSkills, models.Skill{
			Id:        int(addedSkill.ID),
			SkillName: addedSkill.SkillName,
			Category:  addedSkill.Category.String,
		})
	}

	addedEmployee.Skills = addedSkills

	if err := tx.Commit(ctx); err != nil {
		utility.PrintError(funcName, "Failed to commit transaction", err)
		return models.Employee{}, err
	}

	return addedEmployee, nil
}
