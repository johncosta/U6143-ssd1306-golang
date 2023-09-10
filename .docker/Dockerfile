# syntax=docker/dockerfile:1
FROM golang:1.23 as builder
WORKDIR /src

COPY . /src
RUN ls -ltra /src &&\
    go build -o /bin/main ./cmd/main.go

FROM scratch
COPY --from=builder /bin/main /bin/main
CMD ["/bin/main"]
