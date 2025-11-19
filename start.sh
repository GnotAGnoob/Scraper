#!/bin/bash

echo "🚀 Recipe to Kosik Cart - Startup Script"
echo "========================================"
echo ""

# Check if .env exists
if [ ! -f .env ]; then
    echo "⚠️  No .env file found!"
    echo "Creating .env from .env.example..."
    if [ -f .env.example ]; then
        cp .env.example .env
        echo "✅ Created .env file"
        echo "⚠️  Please edit .env and add your GEMINI_API_KEY"
        echo ""
        exit 1
    else
        echo "❌ No .env.example found"
        exit 1
    fi
fi

# Source .env
export $(cat .env | grep -v '^#' | xargs)

# Check if GEMINI_API_KEY is set
if [ -z "$GEMINI_API_KEY" ]; then
    echo "❌ GEMINI_API_KEY not set in .env file"
    echo "Please add your Gemini API key to .env file"
    echo "Get one at: https://aistudio.google.com/"
    exit 1
fi

echo "✅ Environment variables loaded"
echo ""

# Check if go.mod exists
if [ ! -f go.mod ]; then
    echo "❌ go.mod not found. Are you in the project root?"
    exit 1
fi

# Check if frontend/node_modules exists
if [ ! -d frontend/node_modules ]; then
    echo "📦 Installing frontend dependencies..."
    cd frontend
    npm install
    cd ..
    echo "✅ Frontend dependencies installed"
    echo ""
fi

# Install Go dependencies
echo "📦 Checking Go dependencies..."
go mod download
echo "✅ Go dependencies ready"
echo ""

echo "🎯 Starting services..."
echo ""
echo "Backend will be available at:  http://localhost:8080"
echo "Swagger docs at:               http://localhost:8080/swagger/index.html"
echo "Frontend will be available at: http://localhost:5173"
echo ""
echo "Press Ctrl+C to stop all services"
echo ""

# Start backend in background
echo "🔧 Starting backend..."
go run ./cmd/server &
BACKEND_PID=$!

# Wait a bit for backend to start
sleep 2

# Start frontend in background
echo "🎨 Starting frontend..."
cd frontend
npm run dev &
FRONTEND_PID=$!
cd ..

# Function to cleanup on exit
cleanup() {
    echo ""
    echo "🛑 Shutting down services..."
    kill $BACKEND_PID 2>/dev/null
    kill $FRONTEND_PID 2>/dev/null
    exit 0
}

trap cleanup INT TERM

# Wait for processes
wait

