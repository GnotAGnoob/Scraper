# ScrapeMonkey

ScrapeMonkey - is a project designed to scrape product information from online food stores written in Go. Currently, it supports scraping products from Kosik.cz, but the plan is to add additional stores in the future.

## Features

Efficient Scraping: Extracts product data such as names, prices, descriptions, and more.

Modular Design: Built with a flexible architecture to add new stores with minimal effort.

## Getting Started

### Prerequisites

Make sure you have the following installed:

- `Go` (version 1.23 or later)

- installed `make` (optional)

- Internet connection

### Usage

There are predefined commands using makefile, you can use them to build and run the project. If you don't have `make` installed, you can run the commands manually.

#### Console view

##### Build the project

tbd

##### Run the project

```bash
make con
```

##### Run the project in debug mode

```bash
make con-debug
```

#### Benchmark

##### Run the benchmark

```bash
make bench
```

##### View the benchmark results

```bash
make bench-stat
```

##### Clean the benchmark results

```bash
make bench-clean
```

### Recipe to Kosik Cart

A web application that extracts ingredients from recipe text using Google Gemini AI and finds matching products on Kosik.cz.

#### Prerequisites

- Go 1.22+
- Node.js 18+ and npm
- Google Gemini API Key (free at https://aistudio.google.com/)

#### Setup

1. Set your Gemini API key (choose one method):

   **Option A: Create a `.env` file** (recommended):
   ```bash
   cp .env.example .env
   # Then edit .env and add your API key
   ```

   **Option B: Set environment variable**:
   ```bash
   # Windows (Git Bash)
   export GEMINI_API_KEY=your_api_key_here
   
   # Windows (PowerShell)
   $env:GEMINI_API_KEY="your_api_key_here"
   
   # Windows (CMD)
   set GEMINI_API_KEY=your_api_key_here
   ```

2. Start the backend server:
```bash
make server
# or with debug logging:
make server-debug
```

3. In another terminal, start the frontend:
```bash
cd frontend
npm install
npm run dev
```

4. Open http://localhost:3000 in your browser

#### API Documentation

Once the server is running, visit http://localhost:8080/swagger/index.html for Swagger UI documentation.

### Future Plans

- Add scraping support for more online stores.

- Implement advanced filtering and sorting for scraped data.

- Add more types of user interfaces

- In a distant future, add ai to choose products and analyze the best store to buy from based on the user's prompt.
