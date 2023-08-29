#! buat image container
FROM golang:1.20.3-alpine AS build


#! buat folder untuk niympan code
WORKDIR /goapp

#! Copy semua file
COPY . .

#! install depedency and build
RUN go mod download
RUN go build -v -o /goapp/goback ./cmd/main.go

#! create other images 
# Final stage
FROM alpine:3.14

WORKDIR /goapp

#! copy build file
COPY --from=build /goapp /goapp

ENV PATH="/goapp:${PATH}"

EXPOSE 8081

ENTRYPOINT ["goback", "--listen"]

