@echo off

cd ../../../..
set GOPATH=%GOPATH%;%cd%
cd %~dp0
goto %1%
exit

:run
    go run source/main/main.go
    exit

:build
    go build -o build/sm%2%.exe source/main/main.go
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
    
:glide_install
    glide install
    exit
    
:glide_update
    glide update
    exit

:test
    go test ./test
    exit
    
:benchmark
    go test -bench=. ./test
    exit