FROM golang:1.22.5

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /SchoolFile

EXPOSE 8080

CMD ["/SchoolFile"]