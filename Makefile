.PHONY: binary
binary: bin/goc-analyze

.PHONY: bin/goc-analyze
bin/goc-analyze:
	go build -o ./bin/goc-analyze ./cmd/goc-analyze/main.go

.PHONY: clean
clean: clean-binary

.PHONY: clean-binary
clean-binary:
	rm -rf ./bin