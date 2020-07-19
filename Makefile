APP=goland


build: clean
	echo "building"
	go build -o ${APP} main.go

run:
	go run -race main.go

clean:
	go clean
