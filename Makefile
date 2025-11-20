.PHONY: con
con: ./cmd/console/main.go
	@go run ./cmd/console $(ARGS)

.PHONY: con-debug
con-debug:
	@$(MAKE) con ARGS='-log-level=debug'

.PHONY: server
server: ./cmd/server/main.go
	@go run ./cmd/server $(ARGS)

.PHONY: frontend
frontend:
	@echo "frontend app not initialized yet. run bun dev once ready."

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
