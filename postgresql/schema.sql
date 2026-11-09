DROP SCHEMA IF EXISTS skillz_db CASCADE;

CREATE SCHEMA skillz_db;

SET schema 'skillz_db';

CREATE TABLE employee (
	id SERIAL PRIMARY KEY,
	employee_name VARCHAR(50) NOT NULL,
	age INT
);

CREATE TABLE skill (
	id SERIAL PRIMARY KEY,
	skill_name VARCHAR(20) NOT NULL,
	category VARCHAR(8),
	employee_id INT ,
	constraint fk_employee foreign key(employee_id) references employee(id) on delete CASCADE
);


insert into employee (employee_name, age) values ('Eunice Reneau', 25);
insert into employee (employee_name, age) values ('Jennifer Ewing', 27);
insert into employee (employee_name, age) values ('Kimberly Silva', 26);
insert into employee (employee_name, age) values ('Sherry Hampton', 28);
insert into employee (employee_name, age) values ('Roberto Estey', 24);
insert into employee (employee_name, age) values ('Rose Morris', 25);

insert into skill (skill_name, category, employee_id) values ('Java', 'Backend',1);
insert into skill (skill_name, category,employee_id) values ('Python', 'Backend',2);
insert into skill (skill_name, category,employee_id) values ('Angular', 'Frontend',3);
	


select * from skill;

select * from employee;