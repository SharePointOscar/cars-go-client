FROM golang:1.11.5

# make the 'app' folder the current working directory
WORKDIR /go/src/app

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8080
CMD ["go", "run","main.go"]