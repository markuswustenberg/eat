.PHONY: build
build:
	go build -o eat main.go

.PHONY: build
run: build
	./eat

.PHONY: demo
demo: build
	mkdir -p demo
	dd if=/dev/urandom of=demo/yummy.bin bs=1024 count=100000
	dd if=/dev/urandom of=demo/nomnom.bin bs=1024 count=100000
	mv eat demo/
	cd demo && ./eat
