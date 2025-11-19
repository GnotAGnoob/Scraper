@echo off
echo.
echo ================================
echo Recipe to Kosik Cart - Startup
echo ================================
echo.

REM Check if .env exists
if not exist .env (
    echo WARNING: No .env file found!
    if exist .env.example (
        copy .env.example .env
        echo Created .env file
        echo Please edit .env and add your GEMINI_API_KEY
        echo Get one at: https://aistudio.google.com/
        pause
        exit /b 1
    ) else (
        echo ERROR: No .env.example found
        pause
        exit /b 1
    )
)

REM Load environment variables from .env
for /f "usebackq tokens=* delims=" %%a in (".env") do (
    set "%%a"
)

REM Check if GEMINI_API_KEY is set
if "%GEMINI_API_KEY%"=="" (
    echo ERROR: GEMINI_API_KEY not set in .env file
    echo Please add your Gemini API key to .env file
    echo Get one at: https://aistudio.google.com/
    pause
    exit /b 1
)

echo Environment variables loaded
echo.

REM Check if frontend dependencies are installed
if not exist frontend\node_modules (
    echo Installing frontend dependencies...
    cd frontend
    call npm install
    cd ..
    echo Frontend dependencies installed
    echo.
)

echo Checking Go dependencies...
go mod download
echo Go dependencies ready
echo.

echo ================================
echo Starting services...
echo ================================
echo.
echo Backend:  http://localhost:8080
echo Swagger:  http://localhost:8080/swagger/index.html
echo Frontend: http://localhost:5173
echo.
echo Press Ctrl+C to stop all services
echo.

REM Start backend
start "Backend Server" cmd /k "go run ./cmd/server"

REM Wait for backend to start
timeout /t 3 /nobreak > nul

REM Start frontend
start "Frontend Dev Server" cmd /k "cd frontend && npm run dev"

echo.
echo Services started!
echo Close the spawned windows to stop the services.
echo.
pause

