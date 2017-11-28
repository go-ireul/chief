image:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build
	docker build -t ireul/chief .
.PHONY: image
