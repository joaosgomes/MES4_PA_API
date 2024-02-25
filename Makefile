DAEMON := ~/go/bin/CompileDaemon

default: listen
	
build:
	go build main.go

run:
	go run main.go

swagger:
	swag init

listen:
	$(DAEMON) -build="go build main.go" -command="./main"

clean:
	rm ./main ./main.exe