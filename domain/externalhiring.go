package domain

type Skill struct {
	Skill1 string
	Skill2 string
	Skill3 string
	Skill4 string
}
type Freshers struct {
	Candidateid string
	Fullname    string
	Skillset    []Skill
}

type ExternalRepository interface {
	FindAll() ([]Freshers, error)
}
