.PHONY: binary
binary: bin/goc-analyze bin/goc-persistent

.PHONY: bin/goc-analyze
bin/goc-analyze:
	go build -o ./bin/goc-analyze ./cmd/goc-analyze/main.go

.PHONY: bin/goc-persistent
bin/goc-persistent:
	go build -o ./bin/goc-persistent ./cmd/goc-persistent/main.go

.PHONY: clean
clean: clean-binary

.PHONY: clean-binary
clean-binary:
	rm -rf ./bin

.PHONY: image
image: image-goc-analyze image-goc-persistent

.PHONY: image-goc-analyze
image-goc-analyze:
	docker build -t ghcr.io/strrl/goc-analyze:latest -f ./image/goc-analyze/Dockerfile .

.PHONY: image-goc-persistent
image-goc-persistent:
	docker build -t ghcr.io/strrl/goc-persistent:latest -f ./image/goc-persistent/Dockerfile .