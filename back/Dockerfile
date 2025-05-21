FROM golang:1.24
WORKDIR /back

COPY . .
RUN go mod tidy

RUN go install github.com/99designs/gqlgen@latest

RUN go install github.com/air-verse/air@latest
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

CMD ["air"]
EXPOSE 8080
