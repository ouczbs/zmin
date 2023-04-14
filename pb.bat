@echo off & title pb tool
cd %~dp0/engine/sync/zpb
protoc --go_out=. *.proto
echo protoc --go_out=. *.proto
cd %~dp0/tool/proto
go run ./
echo go run ./
pause