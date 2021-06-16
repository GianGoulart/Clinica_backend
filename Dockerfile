FROM golang:1.15-alpine

ENV GO111MODULE=on

WORKDIR /app

COPY ./go.mod .

RUN go mod download

COPY . .

EXPOSE 5055

# Run executable
CMD ["go", "run", "main.go"]
