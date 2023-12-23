# syntax=docker/dockerfile:1
FROM golang:1.21.5 as builder
WORKDIR /src
COPY . /src
RUN go build -o /bin/main ./cmd/...

# TODO: re-enable scratchfiles
#FROM scratch
#COPY --from=builder /bin/main /bin/main
CMD ["/bin/main"]
