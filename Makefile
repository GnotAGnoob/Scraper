.PHONY: con
con: ./cmd/console/main.go
	@go run ./cmd/console $(ARGS)

.PHONY: con-dev
con-dev:
	@$(MAKE) con ARGS='-rod="show,slow=10s,trace,devtools"'
