FROM golang:1.21

WORKDIR /app

COPY . /app

EXPOSE 8080

CMD ["go", "run", "main.go"]