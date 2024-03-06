FROM golang:1.17.7-alpine as build

# buat direktori app
RUN mkdir /app

# set working dir
WORKDIR /app

COPY ./ /app

RUN go mod tidy

RUN go build -o capstone-alta1

EXPOSE 80

CMD [ "./capstone-alta1" ]