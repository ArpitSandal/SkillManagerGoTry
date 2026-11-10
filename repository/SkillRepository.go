package repository

import (
	"context"

	"github.com/ArpitSandal/SkillManagerGo/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ISkillRepository interface {
	SearchEmployeesBySkillRepository(skills []string) ([]models.Employee, error)
	AddEmployeeRepository(ctx context.Context, addedEmployee models.Employee) (models.Employee, error)
	RemoveEmployeeRepository()
}

type SkillRepository struct {
	pgxPool      *pgxpool.Pool
	queryBuilder SkillRepositoryQueryBuilder
}

func NewSkillRepository(pgxPool *pgxpool.Pool) SkillRepository { // dependency injection
	return SkillRepository{
		pgxPool:      pgxPool,
		queryBuilder: SkillRepositoryQueryBuilder{},
	}
}

func (repo SkillRepository) RemoveEmployeeRepository() {}
