FROM golang:1.23.3-alpine3.20 as go

WORKDIR /app
COPY . .
RUN go mod download && go mod verify
RUN go build -o app

FROM alpine:latest

WORKDIR /app
RUN apk upgrade
COPY --from=go app ./

EXPOSE 8080

ENTRYPOINT [ "./app" ]