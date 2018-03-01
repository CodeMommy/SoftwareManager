@echo off

set include=./build
set execute=sm.exe

setlocal DISABLEDELAYEDEXPANSION
set application="%~dp0%include%/%execute%"
set PATH=%PATH%;%~dp0%include%
%application% %*