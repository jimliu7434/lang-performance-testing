SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build

rd /s/Q archive\darwin
xcopy /I /Y performance archive\darwin\