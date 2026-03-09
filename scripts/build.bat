@echo off
REM Build script for Go MySQL REST API (Windows)

echo Building Go MySQL REST API...

SET APP_NAME=go_mysql_api
SET BUILD_DIR=.\bin
SET MAIN_FILE=.\cmd\app\main.go

REM Create build directory if it doesn't exist
if not exist %BUILD_DIR% mkdir %BUILD_DIR%

REM Build for Windows
echo Building for Windows...
go build -o %BUILD_DIR%\%APP_NAME%.exe %MAIN_FILE%

echo Build complete! Executable: %BUILD_DIR%\%APP_NAME%.exe
