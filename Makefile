.PHONY: build clean help

GOOS = $(shell go env GOOS)

## build: создать исполняемый файл
build:
ifeq ($(GOOS),windows)
	go build -o bin/screenmon.exe -ldflags "-s -w" cmd/screenmon/main.go
else
	go build -o bin/screenmon -ldflags "-s -w" cmd/screenmon/main.go
endif

## clean: удалить содержимое папки bin
clean:
	rm -f bin/*

help: Makefile
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'