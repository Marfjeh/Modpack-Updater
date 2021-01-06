build:
	go build -o bin/Modpack-Updater main.go
	chmod +x bin/Modpack-Updater

build-all:
	./scripts/buildall.sh

run:
	go run main.go

clean:
	rm -rf mods/
	rm -rf config/
	rm -rf bin/