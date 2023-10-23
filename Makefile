init:
	go mod tidy
	go mod verify
	go mod vendor

update:
	go get -u ./runner/
	go get -u ./utils/
	go get -u ./exec/
	go mod tidy
