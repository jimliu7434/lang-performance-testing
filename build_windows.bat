SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64
go build

rd /s/Q archive\windows
xcopy /I /Y performance.exe archive\windows\