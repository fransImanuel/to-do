FROM golang:1.19

ENV MYSQL_HOST=host \
    MYSQL_USER=user \
    MYSQL_PORT=port \
    MYSQL_PASSWORD=password \
    MYSQL_DBNAME=dbname 
    

WORKDIR /app


COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go mod tidy


RUN go build -o /docker-gs-ping

CMD ["/docker-gs-ping"]
