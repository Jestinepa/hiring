package domain

type ExternalRepositoryStub struct {
	freshers []Freshers
}

func (s ExternalRepositoryStub) FindAll() ([]Freshers, error) {
	return s.freshers, nil
}

func NewExternalRespositoryStub() ExternalRepositoryStub {
	fresher := []Freshers{
		{
			Candidateid: "1",
			Fullname:    "Tony",
			Skillset: []Skill{
				{
					Skill1: "java",
					Skill2: "docker",
					Skill3: "kubernetes",
					Skill4: "spring-boot",
				},
			},
		},
		{
			Candidateid: "2",
			Fullname:    "steve",
			Skillset: []Skill{
				{
					Skill1: "java",
					Skill2: "docker",
					Skill3: "python",
					Skill4: "golang",
				},
			},
		},
		{
			Candidateid: "3",
			Fullname:    "banner",
			Skillset: []Skill{
				{
					Skill1: "java",
					Skill2: "c++",
					Skill3: "python",
					Skill4: "golang",
				},
			},
		},
	}
	return ExternalRepositoryStub{fresher}
}
