FROM golang:1.14

LABEL author='SDTTTTT' email="760159537@qq.com" description="Go-tds Hub Service."

WORKDIR $GOPATH/src/github.com/sdttttt/go-tds
ADD . $GOPATH/src/github.com/sdttttt/go-tds

RUN go build -v -o entry .

EXPOSE 1234

ENTRYPOINT ["./entry"]