-- name: AddEmployee :one
INSERT INTO employee (employee_name,age)
VALUES (@employee_name::VARCHAR,@age::INT)
RETURNING id;

-- name: AddEmployeeSkill :one
INSERT INTO skill (skill_name,category,employee_id)
VALUES (@skill_name,@category,@employee_id)
RETURNING *;