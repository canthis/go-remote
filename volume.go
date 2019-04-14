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
	"github.com/gen2brain/beeep"
	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
	"github.com/gorilla/mux"
	"github.com/itchyny/volume-go"
	"github.com/skratchdot/open-golang/open"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {

	appTitle := "Remote Control v0.2 by CanThis"
	localIP := GetOutboundIP().String()
	port := ":8775"
	webAppURL := localIP + port

	systray.SetIcon(icon.Data)
	systray.SetTitle(appTitle)
	systray.SetTooltip(appTitle)
	systray.AddMenuItem(appTitle, appTitle).Disable()
	mURL := systray.AddMenuItem("Launch WEB App in Browser", "Launch WEB App in Browser")
	systray.AddSeparator()
	mQuitOrig := systray.AddMenuItem("Quit", "Quit the whole app")

	go func() {
		<-mURL.ClickedCh
		open.Run("http://" + webAppURL)
	}()

	go func() {
		<-mQuitOrig.ClickedCh
		systray.Quit()
	}()

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

	err := beeep.Notify("Remote Control by CanThis", "Server started at:"+webAppURL, "assets/information.png")
	if err != nil {
		panic(err)
	}

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

func onExit() {
	return
}
