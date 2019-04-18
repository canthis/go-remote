//go:generate goversioninfo
package main

import (
	"encoding/json"
	"net/http"
	"os/exec"

	"github.com/julienschmidt/httprouter"
)

type shutdownData struct {
	Action string `json:"action"`
}

func shutdownHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	decoder := json.NewDecoder(r.Body)

	var data shutdownData
	err := decoder.Decode(&data)

	if err != nil {
		panic(err)
	}

	action := data.Action

	if action == "shutdown" {
		exec.Command("cmd", "/C", "shutdown", "/s").Run()
	}

	if action == "restart" {
		exec.Command("cmd", "/C", "shutdown", "/r").Run()
	}

}
