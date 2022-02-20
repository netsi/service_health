##
# build stage
##
FROM golang:1.17 as build

ARG servicename

WORKDIR /build

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY cmd/ ./cmd
COPY pkg/ ./pkg
COPY internal/ ./internal

RUN cd cmd/$servicename && GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -o /app

##
# run stage
##
FROM alpine:3.15

COPY --from=build /app /

CMD ["/app"]