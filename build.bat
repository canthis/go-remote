Taskkill /IM GoRemote.exe /F
del /f GoRemote.exe
rice embed-go
go generate
go build -ldflags -H="windowsgui -s -w" -o "GoRemote.exe"
rice append /exec "GoRemote.exe"
start GoRemote.exe