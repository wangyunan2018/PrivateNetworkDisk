FROM golang:1.16
WORKDIR $GOPATH/src/my-network-disk
COPY . $GOPATH/src/my-network-disk
ENV GOPROXY https://goproxy.cn
EXPOSE 8989
ENTRYPOINT ["go","mod","tidy"]
CMD go get github.com/mattn/go-isatty@v0.0.12
ENTRYPOINT ["go","run","main.go"]