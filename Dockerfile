FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod .

RUN go mod download
COPY . .

RUN go build -o /companies-api cmd/companies-api/*.go

EXPOSE 3000

ENTRYPOINT [ "/companies-api" ]
