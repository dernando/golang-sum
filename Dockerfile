FROM golang:latest as base

WORKDIR /go/src/app
COPY . .

RUN go build -ldflags="-w -s" -o app .

FROM scratch as prd

COPY --from=base /go/src/app /

ENTRYPOINT ["/app"]