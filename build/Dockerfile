FROM golang:1.16.1  as builder
LABEL maintainer="Satoru Murata"
RUN apt update -y && apt install -y \
    zip \
    git 

WORKDIR /home
RUN git clone https://github.com/mobile-core/rcmd.git

WORKDIR /home/rcmd
RUN go mod tidy \
    && mkdir ./bin \
    && GOOS=windows GOARCH=amd64 go build -o ./bin/rcmd.exe main.go \
    && GOOS=linux GOARCH=amd64 go build -o ./bin/rcmd main.go \
    && mv ./bin ./rcmd

RUN zip -r ./bin/rcmd-windows.zip ./rcmd/

FROM alpine:3.13.6 as runner
WORKDIR /home
COPY --from=builder /home/rcmd/bin/rcmd-windows.zip ./
