FROM gcr.io/cloud-builders/go as base

WORKDIR /src/go-sum
COPY . .

RUN go build -ldflags="-w -s" -o go-sum .

ENTRYPOINT ["/go-sum"]