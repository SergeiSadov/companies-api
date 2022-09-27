FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod .

RUN go mod download
COPY . .

RUN go build -o /companies-api cmd/companies-api/*.go

EXPOSE 3000

CMD [ "/companies-api" ]