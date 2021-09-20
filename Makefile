.PHONY: build

build:
	docker build -t try-go-wkhtml .
run:
	make build
	docker run --rm -v ${PWD}/data:/data try-go-wkhtml