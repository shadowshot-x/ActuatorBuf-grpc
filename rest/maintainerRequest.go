package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync/atomic"

	pkg "github.com/shadowshot-x/ActuatorBuf-grpc/pkg/simpleVariableActuate"
)

type MaintainerController struct {
	P *atomic.Value
}

func NewMaintainerController() *MaintainerController {
	return &MaintainerController{}
}

func (ctrl *MaintainerController) StatePostHandler(rw http.ResponseWriter, r *http.Request) {
	var expectedVariableState pkg.SimpleVariable
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&expectedVariableState)
	if err != nil {
		log.Fatalln("There was an error decoding the request body into the struct", err)
	}
	defer r.Body.Close()
	expectedVariableState.SetVar1(10)
	ctrl.P.Store(expectedVariableState)
	fmt.Println(ctrl.P.Load().(pkg.SimpleVariable))
	fmt.Printf("Saved State Variable - %+v\n", expectedVariableState)

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("State Variable Written\n"))
}
