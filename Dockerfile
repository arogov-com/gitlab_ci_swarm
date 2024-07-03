FROM golang:1.22-alpine AS build

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app
COPY main.go /app
COPY go.mod /app

RUN go get -d && go build -ldflags "-s -w" -o main .


FROM scratch

WORKDIR /app
COPY --from=build /app/main ./

EXPOSE 80

CMD ["./main"]
