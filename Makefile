.PHONY: build
build:
	go build -o eat main.go

.PHONY: build
run: build
	./eat

.PHONY: demo
demo: build prepare-food
	mv eat demo/
	clear
	@cd demo && ./eat

.PHONY: prepare-food
prepare-food:
	mkdir -p demo
	dd if=/dev/urandom of=demo/yummy.bin bs=1024 count=10000 &> /dev/null
	dd if=/dev/urandom of=demo/nomnom.bin bs=1024 count=10000 &> /dev/null
