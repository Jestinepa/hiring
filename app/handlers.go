package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/Jestinepa/hiring/service"
)

type Skill struct {
	Skill1 string `json:"skill1" xml:"skill1"`
	Skill2 string `json:"skill2" xml:"skill2"`
	Skill3 string `json:"skill3" xml:"skill3"`
	Skill4 string `json:"skill4" xml:"skill4"`
}

type Candidate struct {
	Candidateid int    `json:"candidateid" xml:"candidateid"`
	Fullname    string `json:"fullname" xml:"fullname"`
	Skillset    []Skill
}

type FreshersHandlers struct {
	service service.FreshersService
}

func (dh *FreshersHandlers) getrequirement(w http.ResponseWriter, r *http.Request) {

	freshers, _ := dh.service.GetAllFresher()

	if r.Header.Get("Conten-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(freshers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(freshers)
	}
}
