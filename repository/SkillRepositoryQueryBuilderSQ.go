package repository

import (
	"log/slog"

	sq "github.com/Masterminds/squirrel"
)

type ISkillRepositoryQueryBuilder interface {
	SearchEmployeesBySkillQueryBuilder(skills []string) (string, []interface{}, error)
}

type SkillRepositoryQueryBuilder struct{}

// squirrel implemenation of fetching the data
func (queryBuilder SkillRepositoryQueryBuilder) SearchEmployeesBySkillQueryBuilder(skills []string) (string, []interface{}, error) {
	const funcName = "SearchEmployeesBySkillQueryBuilder"
	query := sq.Select(
		"e.id",
		"e.employee_name",
		"e.age",
		"s.category",
		"s.skill_name",
		"s.id",
	).From("employee AS e").
		Join("skill AS s ON e.id = s.employee_id")

	orBuilder := make(sq.Or, 0)
	for _, skill := range skills {
		orBuilder = append(orBuilder, sq.Expr("s.skill_name ILIKE (?)", skill)) // equal ignore case
	}

	query = query.Where(orBuilder).PlaceholderFormat(sq.Dollar) // in dollar format for postgres

	// print the query
	slog.Info(funcName, "query", sq.DebugSqlizer(query.PlaceholderFormat(sq.Question)))

	return query.ToSql()
}
