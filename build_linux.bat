SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build

rd /s/Q archive\linux
xcopy /I /Y performance archive\linux\
