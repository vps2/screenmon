.PHONY: run test build clean help

GOOS = $(shell go env GOOS)
DISPLAY=1
TIMEOUT=5s

## run: запустить программу. Можно установить значения переменных DISPLAY и/или TIMEOUT. 
run:
	go run cmd/screenmon/main.go -d=$(DISPLAY) -t=$(TIMEOUT)

## test: запустить тесты
test:
	go test ./...

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
	@powershell '(Get-Content $< -Encoding utf8) -match "^##" -replace "^##(.*?):\s(.*?)"," `$$1`t`$$2"'
else
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
endif