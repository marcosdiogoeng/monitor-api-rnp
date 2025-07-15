FROM golang:1.24-alpine

WORKDIR /app
COPY . .

RUN apk add --no-cache iputils && \
    go build -o agent .

CMD ["./agent"]
