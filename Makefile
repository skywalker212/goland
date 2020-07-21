APP=goland
SRCPATH=./
TARGET=main.go

build: clean
	echo "building"
	go build -o ${APP} ${SRCPATH}${TARGET}

run:
	go run -race ${SRCPATH}${TARGET}

clean:
	go clean
