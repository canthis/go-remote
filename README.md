# Go Remote Volume [![MIT License](http://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/itchyny/volume-go/blob/master/LICENSE)
Basic app written in golang and Vue.js to control PC master volume over LAN and Shutdown (Windows only)

- Golang is used for backend (expose API endpoints and change volume).
- Vue.js is used to provide WEB interface and interact with backend API.

## Build
```
rice embed-go
go build
```
This should compile ```volume.exe``` file

## Running
Launch ```volume.exe``` and navigate in browser to your PC's Local IP, e.g. ```localhost:8775``` or ```192.168.x.x:8775```

## Screenshot
<p>
<img src="https://dev.canthis.lv/storage/app/media/Screenshots/go-remote-volume-chrome-android.jpg" width="25%" height="25%" />
</p>

## Credits
- https://github.com/itchyny/volume-go
- https://codepen.io/calebbrewer/pen/pdyNbb

## License
This software is released under the MIT License, see LICENSE.
