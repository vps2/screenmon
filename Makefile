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
ifeq ($(GOOS),windows)
	powershell "Get-ChildItem bin/* -Recurse | Remove-Item -Recurse"
else
	rm -rf bin/*
endif

help: Makefile
ifeq ($(GOOS),windows)
	@powershell '(Get-Content $< -Encoding utf8) -match "^##" -replace "^##(.*?):\s(.*?)"," `$$1   `$$2"'
else
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
endif