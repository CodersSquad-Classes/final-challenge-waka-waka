build:
	go build main.go

test: build
	@echo Test1
	go run main.go 5

clean:
	rm -rf main