FROM golang:1.22.0-alpine3.19 as builder
COPY ./task2 /go/task2
WORKDIR /go/task2
RUN go mod vendor
RUN cd ./cmd && \
go build -o main .

FROM alpine:latest
WORKDIR /go/
EXPOSE 8100
COPY --from=builder /go/task2/cmd/main ./app
CMD ["./app", "-listening-address=0.0.0.0:8100"]

