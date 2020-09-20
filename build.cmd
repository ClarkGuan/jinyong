@echo off

rsrc -manifest ui\main.manifest -o ui\rsrc.syso
REM go build -ldflags="-H windowsgui" -o jinyong.exe ./ui
go build -o jinyong.exe ./ui
