FROM golang:1.19-alpine

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod ./
COPY go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -v -o /hello-world

EXPOSE 8080

CMD ["hello-world"]