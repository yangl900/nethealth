# build stage
FROM golang:alpine AS build-env
ADD . /go/src/github.com/yangl900/nethealth

WORKDIR /go/src/github.com/yangl900/nethealth/
RUN CGO_ENABLED=0 GOOS=linux go build -o nethealth

# final stage
FROM alpine

RUN apk add --no-cache ca-certificates

WORKDIR /nethealth
COPY --from=build-env /go/src/github.com/yangl900/nethealth/nethealth /nethealth

ENTRYPOINT ./nethealth