set CGO_ENABLED=0
set GOOS=windows
set GOARCH=
go build -o inchv2.exe
#-ldflags "-H windowsgui"