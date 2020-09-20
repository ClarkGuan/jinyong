@echo off

rsrc -manifest win\main.manifest -o win\rsrc.syso
go build -ldflags="-H windowsgui" -o jinyong.exe ./win
