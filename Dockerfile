FROM golang:1.14 as builder

COPY . /go/src/github.com/dhiltgen/localtemp
RUN go build -o /bin/localtemp github.com/dhiltgen/localtemp/cmd


FROM debian:buster
RUN apt-get update && apt-get install -y ca-certificates curl jq && rm -rf /var/lib/apt/lists/*
COPY --from=builder /bin/localtemp /bin/
ENTRYPOINT ["/bin/localtemp"]
