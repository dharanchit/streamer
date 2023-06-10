FROM golang:1.19
RUN mkdir /streaming
WORKDIR /streaming

COPY go.mod ./
COPY go.sum ./

RUN go mod download
COPY . .

EXPOSE 8080

RUN CGO_ENABLED=0 GOOS=linux go build -o /streaming/exec .

CMD ["/streaming/exec"]