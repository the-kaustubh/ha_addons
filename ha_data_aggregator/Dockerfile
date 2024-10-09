FROM golang:1.23 as build

WORKDIR /app

COPY src/go.mod src/go.sum ./

RUN go mod download

COPY src .

RUN CGO_ENABLED=0 go build -o main

FROM scratch

COPY --from=build /app/main main

CMD ["./main"]
