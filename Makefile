.PHONY: con
con: ./cmd/console/main.go
	@go run ./cmd/console $(ARGS)

.PHONY: con-debug
con-debug:
	@$(MAKE) con ARGS='-log-level=debug'

.PHONY: bench
bench:
	@echo "Enter a name for the benchmark result file (without extension):"; \
	read filename; \
	mkdir -p ./benchmark; \
	TIMESTAMP=$$(date +"%Y-%m-%d_%H-%M-%S"); \
	go test -bench=. -count=6 ./... | tee ./benchmark/$${filename}_$${TIMESTAMP}.txt

.PHONY: bench-clean
bench-clean:
	@rm -rf ./benchmark

# go install golang.org/x/perf/cmd/benchstat@latest
.PHONY: bench-stat
bench-stat:
	@benchstat benchmark/*.txt

.PHONY: server
server:
	@go run ./cmd/server

.PHONY: deps
deps:
	@go get github.com/go-chi/chi/v5
	@go get github.com/go-chi/cors
	@go get github.com/swaggo/http-swagger
	@go get github.com/google/generative-ai-go/genai
	@go get google.golang.org/api/option
	@go mod tidy

.PHONY: frontend-install
frontend-install:
	@cd frontend && npm install

.PHONY: frontend
frontend:
	@cd frontend && npm run dev

.PHONY: dev
dev:
	@echo "Starting backend and frontend..."
	@echo "Backend: http://localhost:8080"
	@echo "Frontend: http://localhost:5173"
	@echo "Swagger: http://localhost:8080/swagger/index.html"