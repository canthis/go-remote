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

func shutdownHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	decoder := json.NewDecoder(r.Body)

	var data shutdownData
	err := decoder.Decode(&data)

	if err != nil {
		panic(err)
	}

	action := data.Action

	if action == "shutdown" {
		exec.Command("sh", "-c", "systemctl poweroff").Run()
	}

	if action == "restart" {
		exec.Command("sh", "-c", "systemctl reboot").Run()
	}

}
