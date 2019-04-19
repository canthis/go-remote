package main

import (
	"log"
	"net"
	"net/http"
	"strconv"
	"volume/icon"

	rice "github.com/GeertJohan/go.rice"
	"github.com/bitly/go-simplejson"
	"github.com/gen2brain/beeep"
	"github.com/getlantern/systray"
	"github.com/itchyny/volume-go"
	"github.com/julienschmidt/httprouter"
	"github.com/skratchdot/open-golang/open"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {

	appTitle := "GoRemote v0.3.5 by CanThis"
	localIP := GetOutboundIP().String()
	port := ":8775"
	webAppURL := localIP + port

	systray.SetIcon(icon.Data)
	systray.SetTitle(appTitle)
	systray.SetTooltip(appTitle)
	systray.AddMenuItem(appTitle, appTitle).Disable()
	systray.AddSeparator()
	mURL := systray.AddMenuItem("Web App address: "+webAppURL, "Open WEB App in Browser")
	systray.AddSeparator()
	mQuitOrig := systray.AddMenuItem("Quit", "Quit the whole app")

	go func() {
		for {
			select {
			case <-mURL.ClickedCh:
				open.Run("http://" + webAppURL)
			case <-mQuitOrig.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()

	router := httprouter.New()
	router.GET("/api/system/volume/get", getVolumeHandler)
	router.GET("/api/system/volume/set/:volume", setVolumeHandler)
	router.GET("/api/system/volume/mute", muteHandler)
	router.GET("/api/system/volume/unmute", unmuteHandler)
	router.POST("/api/system/shutdown", shutdownHandler)
	router.NotFound = http.FileServer(rice.MustFindBox("website").HTTPBox())

	err := beeep.Notify(appTitle, "Server started at:"+webAppURL, "website/favicon-32x32.png")
	if err != nil {
		panic(err)
	}

	log.Fatal(http.ListenAndServe(port, router))

}

func getVolumeHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	vol, err := volume.GetVolume()
	if err != nil {
		log.Fatalf("get volume failed: %+v", err)
	}

	json := simplejson.New()
	json.Set("volume", vol)

	payload, err := json.MarshalJSON()
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func setVolumeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	value := ps.ByName("volume")

	volumeToSet64, _ := strconv.ParseInt(value, 10, 16)
	volumeToSet := int(volumeToSet64)

	volume.SetVolume(volumeToSet)
}

func muteHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	volume.Mute()
}

func unmuteHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	volume.Unmute()
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
