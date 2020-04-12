lbuild: 
	CGO_ENABLED=0 GO111MODULE=on GOOS=linux GOARCH=amd64 go build

drun: dbuild
	docker run  -p 8080:8080 --rm web:1.0

dbuild: build
	docker build -t web:1.0 .  
