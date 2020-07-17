FROM golang:1.13.4-alpine AS builder

WORKDIR /_/go/src/app 
COPY . .

#ENV CGO_ENABLED=0
RUN apk add git gcc build-base
RUN go get -d 
RUN go build -o app .

FROM alpine:3.10.3
#ENV 
WORKDIR /app
COPY --from=builder /_/go/src/app/app /app/

ENTRYPOINT ["/app/app"]
