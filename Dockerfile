FROM golang:latest

LABEL maintainer="Russell <russellcxl@gmail.com>"

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . ./

ENV PORT 5000

RUN go build -o build/main cmd/main.go

EXPOSE ${PORT}

CMD ["./build/main"]