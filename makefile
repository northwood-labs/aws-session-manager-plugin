build: build-linux-amd64 build-arm64

clean:
	rm -rf build/* bin/ pkg/ .cover/
	find . -type f -name '*.log' -delete

.PHONY: build-linux-amd64
build-linux-amd64:
	mkdir -p ./bin/linux_amd64_plugin
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o bin/linux_amd64_plugin/session-manager-plugin \
		src/sessionmanagerplugin-main/main.go

.PHONY: build-arm64
build-arm64:
	mkdir -p ./bin/linux_arm64_plugin
	GOOS=linux GOARCH=arm64 go build -ldflags "-s -w -extldflags=-Wl,-z,now,-z,relro,-z,defs" -o bin/linux_arm64_plugin/session-manager-plugin \
		src/sessionmanagerplugin-main/main.go

.PHONY: build-native
build-native:
	go build -ldflags "-s -w -extldflags=-Wl,-z,now,-z,relro,-z,defs" -o bin/session-manager-plugin \
		src/sessionmanagerplugin-main/main.go
