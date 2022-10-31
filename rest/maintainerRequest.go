package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type MaintainerController struct {
}

func NewMaintainerController() *MaintainerController {
	return &MaintainerController{}
}

func (ctrl *MaintainerController) StatePostHandler(rw http.ResponseWriter, r *http.Request) {
	var expectedVariableState StateVariable
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&expectedVariableState)
	if err != nil {
		log.Fatalln("There was an error decoding the request body into the struct", err)
	}
	defer r.Body.Close()
	fmt.Printf("Saved State Variable - %+v\n", expectedVariableState)

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("State Variable Written\n"))
}
