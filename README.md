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

#### Environment

Copy `.env.example` to `.env`, drop in your Gemini creds, and tweak values as needed. The server refuses to boot if `.env` is missing so you get instant feedback.

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

#### API server

```bash
make server
```

Pass overrides via flags, e.g. `make server ARGS="-port=9090 -log-level=debug"`.

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

### Future Plans

- Add scraping support for more online stores.

- Implement advanced filtering and sorting for scraped data.

- Add more types of user interfaces

- In a distant future, add ai to choose products and analyze the best store to buy from based on the user's prompt.
