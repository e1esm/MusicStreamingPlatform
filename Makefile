build:
	go build *.go

dep:
	go mod download
vet:
	go vet