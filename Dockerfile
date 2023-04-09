FROM golang:1.19-alpine3.17

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


RUN go build -o todo .

EXPOSE 3030

# CMD /docker-gs-ping -MYSQL_HOST=$MYSQL_HOST -MYSQL_USER=$MYSQL_USER -MYSQL_PORT=$MYSQL_PORT -MYSQL_PASSWORD=$MYSQL_PASSWORD -MYSQL_DBNAME=$MYSQL_DBNAME

# CMD ["./todo","-MYSQL_HOST=$MYSQL_HOST","-MYSQL_USER=${MYSQL_USER}","-MYSQL_PORT=${MYSQL_PORT}","-MYSQL_DBNAME=${MYSQL_DBNAME}","-MYSQL_PASSWORD=${MYSQL_PASSWORD}"]
CMD ./todo -MYSQL_HOST=$MYSQL_HOST -MYSQL_USER=$MYSQL_USER -MYSQL_PORT=$MYSQL_PORT -MYSQL_DBNAME=$MYSQL_DBNAME -MYSQL_PASSWORD=$MYSQL_PASSWORD

# docker run --name todo -e MYSQL_HOST=host -e MYSQL_USER=user -e MYSQL_PORT=port -e MYSQL_PASSWORD=pw -e MYSQL_DBNAME=name -p 8090:3003 to-do-app:1.1

# docker run --name todo -e MYSQL_HOST=172.18.0.1 -e MYSQL_USER=user -e MYSQL_PORT=3306 -e MYSQL_PASSWORD=123456 -e MYSQL_DBNAME=todo4 -p 8090:3003 to-do-app:latest

#run in my local 
#go run main.go - MYSQL_HOST="localhost" -MYSQL_USER=user -MYSQL_PORT=3306 -MYSQL_PASSWORD=123456 -MYSQL_DBNAME=todo4   

# go run main.go -MYSQL_HOST="localhost" -MYSQL_USER="root" -MYSQL_PORT="3306" -MYSQL_PASSWORD="123456" -MYSQL_DBNAME="todo4" 