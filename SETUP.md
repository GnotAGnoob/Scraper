# Quick Setup Guide

## 1. Get Your Gemini API Key

1. Visit https://aistudio.google.com/
2. Sign in with your Google account
3. Click "Get API key" in the navigation
4. Create a new API key
5. Copy the key

## 2. Setup Environment

Create a `.env` file in the project root:

```bash
GEMINI_API_KEY=your_actual_api_key_here
PORT=8080
```

## 3. Install Dependencies

### Backend (Go)
```bash
make deps
```

Or manually:
```bash
go get github.com/go-chi/chi/v5
go get github.com/go-chi/cors
go get github.com/swaggo/http-swagger
go get github.com/google/generative-ai-go/genai
go get google.golang.org/api/option
go get golang.org/x/sync/errgroup
go mod tidy
```

### Frontend (Node.js)
```bash
cd frontend
npm install
```

## 4. Run the Application

### Terminal 1 - Backend
```bash
# Load environment variable
export GEMINI_API_KEY=your_key_here  # Linux/Mac
# OR
set GEMINI_API_KEY=your_key_here     # Windows CMD
# OR
$env:GEMINI_API_KEY="your_key_here"  # Windows PowerShell

# Run server
make server
# Or: go run ./cmd/server
```

Backend will be available at: http://localhost:8080

### Terminal 2 - Frontend
```bash
cd frontend
npm run dev
```

Frontend will be available at: http://localhost:5173

## 5. Test It Out

1. Open http://localhost:5173 in your browser
2. Paste a recipe (works best with Czech recipes)
3. Click "Find Products"
4. See matched products from Kosik!

### Example Recipe (Czech)

```
Ingredience:
- 500g hladké mouky
- 250ml mléka
- 2 vejce
- 100g másla
- 1 lžička soli
- 2 lžíce cukru
```

## 6. API Documentation

Visit http://localhost:8080/swagger/index.html for interactive API docs.

## Troubleshooting

### "GEMINI_API_KEY environment variable is required"
- Make sure you've set the environment variable before running the server
- On Windows, use `set` or `$env:` depending on your shell
- The key should be directly from Google AI Studio, not wrapped in quotes in the command

### Port 8080 already in use
- Change the PORT in your `.env` file
- Or stop the process using port 8080

### Frontend can't connect to backend
- Make sure the backend is running first
- Check that it's accessible at http://localhost:8080/health
- Clear browser cache if needed

### No products found
- The Kosik scraper searches Czech products, so Czech ingredient names work best
- Try simpler ingredient names (e.g., "mouka" instead of "hladká pšeničná mouka")
- Check backend logs for any errors

