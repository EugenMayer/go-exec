init:
	go mod tidy
	go mod verify
	go mod vendor

update:
	go get -u
	go mod tidy
