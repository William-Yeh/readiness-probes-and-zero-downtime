#======================================#
# builder env
#

FROM golang:1.12.5 AS builder
LABEL author="william.pjyeh@gmail.com"

COPY hello.go    .
RUN GOOS=linux   GOARCH=amd64  CGO_ENABLED=0 \
    go build -o hello


#======================================#
# deployment env
#

FROM alpine:3.10.0
LABEL author="william.pjyeh@gmail.com"

# copy executable
COPY --from=builder /go/hello  /

# Define default command.
EXPOSE 80
ENTRYPOINT ["/hello"]
