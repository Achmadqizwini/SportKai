FROM golang:1.22.1-alpine3.19 as build

WORKDIR /app

COPY ./ /app

RUN go mod tidy

RUN go build -o sport_kai

EXPOSE 8080

CMD [ "./sport_kai" ]