FROM golang:latest as base

WORKDIR /go/src/app
COPY ./src/go-sum .

RUN go build -ldflags="-w -s" -o app .

FROM scratch as prd

COPY --from=base /go/src/app /

ENTRYPOINT ["/app"]