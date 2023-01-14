FROM golang:1.19-alpine3.17
ENV CGO_ENABLED=0
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install github.com/cosmtrek/air@latest