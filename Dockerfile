FROM golang

WORKDIR /app

COPY . /app

ENV GOPROXY="https://goproxy.cn"

RUN go mod download
RUN go build main.go

EXPOSE 8080

ENTRYPOINT ["./main"]