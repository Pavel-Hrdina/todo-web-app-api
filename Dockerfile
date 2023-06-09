#build stage
FROM golang:alpine AS builder
WORKDIR /go/src/app
COPY . .
RUN apk add build-base
RUN CGO_ENABLED=1 go get -d -v . && go build -ldflags="-s -w" main.go 

#final stage
FROM alpine:latest
LABEL maintainer=numtostr version=0.0.1
COPY --from=builder /go/src/app/main /main
EXPOSE 5000
ENTRYPOINT ["/main"]