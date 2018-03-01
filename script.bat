@echo off

cd ../../../..
set GOPATH=%GOPATH%;%cd%
cd %~dp0
goto %1%
exit

:run
    go run source/main.go
    exit

:build
    go build -o build/sm%2%.exe source/main.go
    exit

:dep_init
    dep init
    exit

:dep_add
    dep ensure -add %2%
    exit

:dep_install
    dep ensure
    exit

:dep_update
    dep ensure -update
    exit

:dep_status
    dep status
    exit