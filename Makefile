.PHONY: con
con: ./cmd/console/main.go
	@go run ./cmd/console $(ARGS)

.PHONY: con-show
con-show:
	@$(MAKE) con ARGS='-rod="show,slow=10s,trace,devtools"'

.PHONY: con-debug
con-debug:
	@$(MAKE) con ARGS='-rod="trace"'
