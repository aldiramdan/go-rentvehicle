FROM golang:1.19.5-alpine

WORKDIR /app

COPY . .

RUN go mod tidy

ENV DB_HOST=${DB_HOST} \
    DB_PORT=${DB_PORT} \
    DB_NAME=${DB_NAME} \
    DB_USER=${DB_USER} \
    DB_PASS=${DB_PASS}

RUN go build -v -o /app/gorentveh

EXPOSE 3080

ENTRYPOINT [ "/app/gorentveh" ]
CMD [ "serve" ]