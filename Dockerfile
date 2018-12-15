FROM golang
COPY . /usr/src/app
WORKDIR /usr/src/app
ENV DB_CONNECTION_STRING=$DB_CONNECTION_STRING
RUN go build .
EXPOSE 5000
ENTRYPOINT /usr/src/app/plate-api
