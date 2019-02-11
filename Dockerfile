FROM golang:1.11

WORKDIR /code

EXPOSE 3000

COPY . .

CMD ["go", "run", "app.go"]