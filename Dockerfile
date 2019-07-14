FROM golang:latest

RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/jinzhu/gorm

COPY . src