#FROM golang:1.6-alpine
FROM arm64v8/golang
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o main .
CMD ["/app/main"]

EXPOSE 8001
EXPOSE 8002