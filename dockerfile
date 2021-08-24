# FROM golang:latest
#
# ENV GOPROXY https://goproxy.cn,direct
# WORKDIR $GOPATH/../go-gin-example
# COPY . $GOPATH/../go-gin-example
# RUN go build .
#
# EXPOSE 8000
# ENTRYPOINT ["./go-gin-example"]
FROM scratch
WORKDIR $GOPATH/../go-gin-example
COPY . $GOPATH/../go-gin-example

EXPOSE 8000
CMD ["./go-gin-example"]

