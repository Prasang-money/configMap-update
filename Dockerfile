FROM golang:1.19-alpine

WORKDIR /root
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
#RUN go build -v -o /usr/local/bin/app ./...
RUN pwd
CMD ["go","run","."]