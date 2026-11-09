package models

import "github.com/ArpitSandal/SkillManagerGo/models/dto"

type Skill struct {
	Id        int    `db:"id"`
	SkillName string `db:"skill_name"`
	Category  string `db:"category"`
}

func (s *Skill) ToEntity(skillDTO dto.SkillDTO) {
	s.Category = skillDTO.Category
	s.SkillName = skillDTO.SkillName
	
}
