.PHONY: test build clean deploy gomodgen

test:
	go test github.com/dariuszkorolczukcom/birthday-app-api/birthday/structs

build: gomodgen test
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o bin/birthday birthday/handler/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/world world/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh
