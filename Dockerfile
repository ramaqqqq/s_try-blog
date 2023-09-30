FROM golang:1.16-alpine

WORKDIR /app

COPY . .

RUN go build -o sagara-try

EXPOSE 7000

CMD ./sagara-try