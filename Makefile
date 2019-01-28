SOURCES=vistecture.go
VERSION=1.0.0

.PHONY: all templates darwin linux windows default

default: darwin

all: darwin linux windows

templates:
    dep ensure
	packr2
	mkdir -p build-artifacts
	zip -qr build-artifacts/templates.zip templates

darwin: $(SOURCES) templates
	GOOS=darwin go build -o build-artifacts/vistecture $(SOURCES)

linux: $(SOURCES) templates
	GOOS=linux go build -o build-artifacts/vistecture-linux $(SOURCES)

windows: $(SOURCES) templates
	GOOS=windows go build -o build-artifacts/vistecture.exe $(SOURCES)

dockerpublish:
	docker build --no-cache -t aoepeople/vistecture .
	docker tag aoepeople/vistecture:latest aoepeople/vistecture:$(VERSION)
	docker push aoepeople/vistecture