FROM golang:1.22-alpine

COPY ./certs/* /usr/local/share/ca-certificates/.
RUN update-ca-certificates

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENTRYPOINT ["go", "run", "cmd/blogcontent/main.go"]