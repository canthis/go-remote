# ![Go Remote Volume ](https://dev.canthis.lv/storage/app/media/Screenshots/goremote-icon-32x32.png) Go Remote Volume 
![GitHub release](https://img.shields.io/github/release/canthis/go-remote.svg?label=version)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/canthis/go-remote-volume/blob/master/LICENSE)

Basic app written in golang and Vue.js to control master volume, shutdown and restart PC over LAN/WiFi (Tested on Windows 10 v1809 and Linux Mint 19.1 Tessa)

- **Go** is used for backend (run http server, expose API endpoints and change volume).
- **Vue.js/buefy/axios** is used for UI and interaction with backend API.

## Build
```
go generate
build -ldflags="-H windowsgui -s -w" -o "GoRemote.exe"
rice append --exec "GoRemote.exe"
```
This should compile ```GoRemote.exe``` file

## Running
- Download Windows binaries: [GoRemote-amd64.exe (64-bit)](https://github.com/canthis/go-remote/releases/latest/download/GoRemote-amd64.exe) 
or [GoRemote-i386.exe (32-bit)](https://github.com/canthis/go-remote/releases/latest/download/GoRemote-i386.exe)
- Download Linux binaries: [GoRemote-linux-amd64 (64-bit)](https://github.com/canthis/go-remote/releases/latest/download/GoRemote-linux-amd64) 
or [GoRemote-linux-i386 (32-bit)](https://github.com/canthis/go-remote/releases/latest/download/GoRemote-linux-i386)
- Launch it and navigate from your mobile browser to your PC's Local IP, e.g. ```localhost:8775``` or ```192.168.x.x:8775```
- Add it to Home Screen and enjoy

## Screenshot
<p>
<img src="https://dev.canthis.lv/storage/app/media/Screenshots/go-remote-volume-chrome-android-v0340.png" width="25%" height="25%" />
</p>

## Credits
- https://github.com/itchyny/volume-go
- https://github.com/GeertJohan/go.rice
- https://github.com/josephspurrier/goversioninfo
- https://github.com/gen2brain/beeep
- https://github.com/getlantern/systray
- https://github.com/julienschmidt/httprouter
- https://codepen.io/calebbrewer/pen/pdyNbb
- App Icon made by [Smashicons](https://www.flaticon.com/authors/smashicons) from [www.flaticon.com](https://www.flaticon.com/) is licensed by [CC 3.0 BY](http://creativecommons.org/licenses/by/3.0/)


## License
This software is released under the MIT License, see LICENSE.

