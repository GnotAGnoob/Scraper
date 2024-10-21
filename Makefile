.PHONY: con
con: ./cmd/console/main.go
	@go run ./cmd/console $(ARGS)

.PHONY: con-show
con-show:
	@$(MAKE) con ARGS='-rod="show,slow=10s,trace,devtools"'

.PHONY: con-debug
con-debug:
	@$(MAKE) con ARGS='-rod="trace" -log-level=debug'

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
