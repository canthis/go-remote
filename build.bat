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

set GOOS=linux
set GOARCH=386
go build -o "GoRemote-linux-i386"
rice append --exec "GoRemote-linux-i386"

set GOOS=linux
set GOARCH=amd64
go build -o "GoRemote-linux-amd64"
rice append --exec "GoRemote-linux-amd64"
