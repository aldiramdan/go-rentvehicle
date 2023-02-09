FROM golang:1.19.5-alpine

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -v -o /app/gorentveh

EXPOSE 3080

ENTRYPOINT [ "/app/gorentveh" ]
CMD [ "serve" ]