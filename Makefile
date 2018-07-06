.PHONY: all get clean build

GO ?= go

all: build

build: get
	${GO} build

get:
	dep ensure

clean:
	@rm -rf cf-sample-app-go-dep vendor
