.PHONY: build

build:
	cd tools/adf && go test -mod=readonly ./...
	mkdir -p bin
	cd tools/adf && CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -mod=readonly -trimpath -ldflags="-s -w" -o ../../bin/adf-darwin-amd64 .
	cd tools/adf && CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -mod=readonly -trimpath -ldflags="-s -w" -o ../../bin/adf-darwin-arm64 .
	cd tools/adf && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=readonly -trimpath -ldflags="-s -w" -o ../../bin/adf-linux-amd64 .
	cd tools/adf && CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -mod=readonly -trimpath -ldflags="-s -w" -o ../../bin/adf-linux-arm64 .
	cd tools/adf && CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -mod=readonly -trimpath -ldflags="-s -w" -o ../../bin/adf-windows-amd64.exe .
