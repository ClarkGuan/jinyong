@echo off

rsrc -manifest win\main.manifest -o win\rsrc.syso
REM go build -ldflags="-H windowsgui" -o jinyong.exe ./win
go build -o jinyong.exe ./win
