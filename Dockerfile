FROM golang:1.11-alpine AS buildstage

WORKDIR /go/src/app
COPY main.go .

RUN go build -o app .

FROM alpine:3.8
WORKDIR /app
COPY --from=buildstage /go/src/app /app/

EXPOSE 8080
CMD ["./app"]
