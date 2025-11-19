# Recipe to Kosik Cart

Transform recipe text into a shopping cart! This app uses AI to extract ingredients from recipe text and automatically finds matching products on Kosik.cz.

## 🎨 Features

- **AI-Powered Ingredient Extraction**: Uses Google Gemini 1.5 Flash to intelligently parse recipes
- **Concurrent Product Search**: Fast parallel scraping of Kosik products
- **Neo-Brutalist UI**: Bold, high-contrast, retro-inspired design
- **RESTful API**: Clean API with Swagger documentation

## 🏗️ Architecture

### Backend (Go)
- **LLM Integration**: Google Gemini API for ingredient extraction
- **Scraper Service**: Concurrent product search using existing Kosik scraper
- **API Server**: Chi router with CORS and Swagger
- **Concurrency**: errgroup for efficient parallel processing

### Frontend (React + Vite + TypeScript)
- **Modern Stack**: React 18 + Vite + TypeScript
- **Neo-Brutalist Design**: Bold borders, hard shadows, bright colors
- **Type-Safe API Client**: Axios with TypeScript interfaces

## 🚀 Getting Started

### Prerequisites

1. **Go 1.22.4+**
2. **Node.js 18+** and npm
3. **Google Gemini API Key** (free at https://aistudio.google.com/)

### Installation

1. **Clone and setup backend:**

```bash
# Install Go dependencies
make deps

# Or manually:
go get github.com/go-chi/chi/v5
go get github.com/go-chi/cors
go get github.com/swaggo/http-swagger
go get github.com/google/generative-ai-go/genai
go get google.golang.org/api/option
go get golang.org/x/sync/errgroup
go mod tidy
```

2. **Setup environment:**

```bash
# Create .env file in project root
cp .env.example .env

# Edit .env and add your Gemini API key:
# GEMINI_API_KEY=your_actual_key_here
# PORT=8080
```

3. **Install frontend dependencies:**

```bash
cd frontend
npm install
```

### Running the Application

**Option 1: Run both servers separately**

Terminal 1 (Backend):
```bash
# Make sure GEMINI_API_KEY is set
export GEMINI_API_KEY=your_key_here
make server
# Backend runs on http://localhost:8080
```

Terminal 2 (Frontend):
```bash
cd frontend
npm run dev
# Frontend runs on http://localhost:5173
```

**Option 2: Manual commands**

```bash
# Backend
go run ./cmd/server

# Frontend (in another terminal)
cd frontend && npm run dev
```

### API Documentation

Once the backend is running, visit:
- **Swagger UI**: http://localhost:8080/swagger/index.html
- **Health Check**: http://localhost:8080/health

### Frontend

Visit http://localhost:5173 to use the app:
1. Paste your recipe text
2. Click "Find Products"
3. View matched products with images and prices

## 📁 Project Structure

```
.
├── cmd/
│   ├── console/          # Original CLI tool
│   └── server/           # API server (NEW)
├── internal/
│   ├── api/              # HTTP handlers (NEW)
│   ├── llm/              # Gemini integration (NEW)
│   ├── service/          # Business logic (NEW)
│   ├── scraper/          # Kosik scraper (existing)
│   └── logger/           # Logging
├── frontend/             # React + Vite app (NEW)
│   ├── src/
│   │   ├── components/   # React components
│   │   ├── App.tsx       # Main app
│   │   └── api.ts        # API client
│   └── package.json
├── docs/                 # Swagger documentation (NEW)
├── Makefile             # Build commands
└── README.md            # This file
```

## 🎯 API Endpoints

### POST `/api/v1/recipe`
Process recipe text and find matching products.

**Request:**
```json
{
  "text": "500g mouky, 250ml mléka, 2 vejce..."
}
```

**Response:**
```json
{
  "products": [
    {
      "ingredientName": "mouka",
      "productName": "Hladká mouka pšeničná",
      "price": 25.90,
      "unit": "kg",
      "imageUrl": "https://...",
      "productUrl": "https://kosik.cz/...",
      "pricePerUnit": 25.90,
      "pricePerUnitUnit": "kg"
    }
  ]
}
```

## 🛠️ Development

### Makefile Commands

```bash
make con              # Run console scraper
make con-debug        # Run with debug logging
make server           # Run API server
make deps             # Install Go dependencies
make bench            # Run benchmarks
make bench-stat       # Show benchmark stats
```

### Original CLI Tool

The original Kosik scraper CLI is still available:

```bash
make con ARGS="-search=mléko"
make con-debug ARGS="-search=mléko"
```

## 🎨 Neo-Brutalist Design System

The frontend uses a bold, retro-inspired aesthetic:

- **Colors**: Yellow (#ffd700), Pink (#ff6b9d), Cyan (#00ffff), Green (#9aff9a)
- **Borders**: Thick 3-5px solid black
- **Shadows**: Hard drop shadows (4-8px)
- **Typography**: Courier New monospace
- **Interactions**: Tactile button presses with shadow movement

## 🔧 Troubleshooting

**Backend won't start:**
- Check `GEMINI_API_KEY` is set in environment
- Ensure port 8080 is available
- Run `make deps` to install missing packages

**Frontend won't connect:**
- Verify backend is running on port 8080
- Check browser console for CORS errors
- Try clearing browser cache

**No products found:**
- Gemini might have trouble parsing the recipe format
- Try with Czech recipe text for best results
- Check backend logs for extraction errors

## 📝 TODO / Future Improvements

- [ ] Add cart management (add/remove items)
- [ ] Save recipes and product lists
- [ ] Support for multiple languages
- [ ] Batch recipe processing
- [ ] Price comparison and suggestions
- [ ] Export to Kosik cart directly (requires Kosik API access)

## 📄 License

MIT

## 🙏 Credits

Built with:
- [Google Gemini](https://ai.google.dev/) - AI ingredient extraction
- [Kosik.cz](https://www.kosik.cz/) - Product data source
- [Go Chi](https://go-chi.io/) - HTTP router
- [React](https://react.dev/) + [Vite](https://vitejs.dev/) - Frontend
