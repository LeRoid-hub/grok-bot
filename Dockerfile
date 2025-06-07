# syntax=docker/dockerfile:1

FROM golang:1.22

WORKDIR /app

COPY go.mod  ./

RUN go mod download

COPY . .

RUN go build -o /grok-bot

EXPOSE 8080

CMD [ "/grok-bot" ]