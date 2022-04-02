package app

import (
	"log"
	"net/http"

	"github.com/Jestinepa/hiring/domain"
	"github.com/Jestinepa/hiring/service"
	"github.com/gorilla/mux"
)

func Start() {

	myRouter := mux.NewRouter()

	//wiring
	dh := FreshersHandlers{service.NewFreshersService(domain.NewExternalRespositoryStub())}

	//define routes
	myRouter.HandleFunc("/getrequirement", dh.getrequirement)

	//starting server

	log.Fatal(http.ListenAndServe("localhost:8087", myRouter))
}
