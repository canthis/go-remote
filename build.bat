set GOOS=windows
set GOARCH=386
Taskkill /IM GoRemote-i386.exe /F
del /f GoRemote-i386.exe
go generate
go build -ldflags -H="windowsgui -s -w" -o "GoRemote-i386.exe"
rice append /exec "GoRemote-i386.exe"

set GOOS=windows
set GOARCH=amd64
Taskkill /IM GoRemote-amd64.exe /F
del /f GoRemote-amd64.exe
go generate
go build -ldflags -H="windowsgui -s -w" -o "GoRemote-amd64.exe"
rice append /exec "GoRemote-amd64.exe"

