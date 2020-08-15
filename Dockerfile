FROM golang
MAINTAINER Gaurav Dhameeja gdhameeja@gmail.com
ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64
CMD mkdir /go/src/gdhameeja/booksapi

WORKDIR /go/src/gdhameeja/booksapi

# setup db
RUN apt-get update
RUN apt-get install sqlite3
COPY boot.sql /tmp/
RUN cat /tmp/boot.sql | sqlite3 books.db

# setup code
COPY main.go /go/src/gdhameeja/booksapi/
ADD models /go/src/gdhameeja/booksapi/models
COPY go.mod /go/src/gdhameeja/booksapi/
COPY go.sum /go/src/gdhameeja/booksapi/
RUN go mod download

RUN go build main.go
EXPOSE 8888

ENTRYPOINT ["./main"]
