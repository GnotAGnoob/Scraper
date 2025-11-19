<!-- 628a6713-0ce0-4b2c-bb88-75b4f5e7fb2a 1e306956-0b67-4ed5-be91-1da7eb994270 -->
# Recipe to Kosik Cart Plan

## 1. Backend Implementation (Go)

- **LLM Integration (`internal/llm`)**:
    - Implement a client for **Google Gemini 1.5 Flash** (free tier via Google AI Studio).
    - Create a prompt to extract ingredients and quantities (in grams) from text into a structured JSON format.
- **Scraper Integration (`internal/service`)**:
    - Create a service that takes a list of ingredients.
    - Launch concurrent goroutines (using `errgroup` or `WaitGroup`) to call `kosik.GetProducts` for each ingredient.
    - Filter to take only the first valid result for each ingredient.
- **API Server (`cmd/server`, `internal/api`)**:
    - Use `go-chi` or `net/http` for routing.
    - Endpoint: `POST /api/v1/recipe`
        - Input: `{"text": "..."}`
        - Output: JSON list of matched products with images/prices.
    - Add CORS middleware.
- **OpenAPI / Swagger**:
    - Annotate handlers with `swag` comments.
    - Generate `docs/` using `swag init`.
    - Serve Swagger UI at `/swagger/index.html`.

## 2. Frontend Implementation (React + Vite)

- **Setup**:
    - Initialize Vite project (`npm create vite@latest frontend -- --template react-ts`).
    - Install `axios`, `tanstack-query` (optional but good), and openapi-generator tools.
- **Client Generation**:
    - Generate TypeScript client from the Swagger JSON.
- **UI Design (Neo-Brutalist)**:
    - **Style**: High contrast, black borders (3px+), bright colors (yellow/pink/blue), hard shadows (`box-shadow: 4px 4px 0px #000`).
    - **Components**:
        - `InputArea`: Large text area for recipe.
        - `SubmitButton`: Big, bold button.
        - `ProductGrid`: Grid of `ProductCard`s.
        - `ProductCard`: Displays image, name, price, and matched ingredient name.

## 3. Integration

- Connect Frontend to Backend API.
- Ensure concurrency works efficiently (Kosik search is the bottleneck).

## User Action Required

- You will need a **Google Gemini API Key** (free at https://aistudio.google.com/).

### To-dos

- [x] Initialize Go server structure (cmd/server, internal/api) and install dependencies (chi, cors, swag).
- [x] Implement LLM service using Google Gemini (Flash 1.5) for ingredient extraction.
- [x] Implement Ingredient Scraper service (concurrent Kosik calls).
- [x] Implement API handlers and Swagger annotations.
- [x] Initialize React + Vite frontend with Neo-brutalist CSS.
- [x] Generate API client and integrate with UI components.

