package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os/exec"
	"strconv"

	rice "github.com/GeertJohan/go.rice"
	"github.com/bitly/go-simplejson"
	"github.com/gorilla/mux"
	"github.com/itchyny/volume-go"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/system/volume/get", func(w http.ResponseWriter, router *http.Request) {

		vol, err := volume.GetVolume()
		if err != nil {
			log.Fatalf("get volume failed: %+v", err)
		}
		fmt.Printf("current volume: %d\n", vol)

		json := simplejson.New()
		json.Set("volume", vol)

		payload, err := json.MarshalJSON()
		if err != nil {
			log.Println(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)

	})

	router.HandleFunc("/api/system/volume/set/{value}", func(w http.ResponseWriter, router *http.Request) {
		vars := mux.Vars(router)
		value := vars["value"]

		volumeToSet64, _ := strconv.ParseInt(value, 10, 16)
		volumeToSet := int(volumeToSet64)

		volume.SetVolume(volumeToSet)
		fmt.Printf("current volume: %d\n", volumeToSet)
	})

	router.HandleFunc("/api/system/volume/mute", func(w http.ResponseWriter, router *http.Request) {
		volume.Mute()
		fmt.Printf("volume muted\n")
	})

	router.HandleFunc("/api/system/volume/unmute", func(w http.ResponseWriter, router *http.Request) {
		volume.Unmute()
		fmt.Printf("volume unmuted\n")
	})

	router.HandleFunc("/api/system/shutdown", func(w http.ResponseWriter, router *http.Request) {
		if err := exec.Command("cmd", "/C", "shutdown", "/s").Run(); err != nil {
			fmt.Println("Failed to initiate shutdown:", err)
		}
	})

	router.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("website").HTTPBox()))

	localIP := GetOutboundIP().String()
	port := ":8775"

	fmt.Println("Server started at:", localIP+port)
	log.Fatal(http.ListenAndServe(port, router))

}

//GetOutboundIP gets preferred outbound ip of this machine
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
