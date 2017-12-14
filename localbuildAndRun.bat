@echo off
echo Build GO File
go build -o bin/application application.go


@echo off
echo Open Test Page
start "" http://localhost:5000

@echo off
copy bin\application bin\application.exe
echo Launch Local GOLang WebServer
.\bin\application

