FROM golang:latest 
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go build -o main . 
EXPOSE 8281
ENTRYPOINT ['pm2-docker', 'main.go']
# CMD ["/app/main"]