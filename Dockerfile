#! buat image container
FROM golang:1.20.3-alpine


#! buat folder untuk niympan code
WORKDIR /goapp

#! Copy semua file
COPY . .

#! install depedency and build
RUN go mod download
RUN go build -v -o /goapp/goback ./cmd/main.go

#! open port for app 
EXPOSE 8081

#! run app
ENTRYPOINT [ "/goapp/goback" ]

#! docker run --name gotest --net local_default -e DB_HOST=pglocal -p 8082:8081 bukanebi/goapps:2