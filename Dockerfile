#Build stage
FROM golang:1.19-alpine AS builder

#Setup repo
COPY . /go/src/d-wojciechowski/ammount-in-words
WORKDIR /go/src/d-wojciechowski/ammount-in-words

#disable crosscompiling
ENV CGO_ENABLED=0
#compile linux only
ENV GOOS=linux
#build
RUN go build  -ldflags '-w -s' -a -o dist/main .

FROM scratch

EXPOSE 8081
ENV APP_URL 0.0.0.0:8081
ENV GIN_MODE release
ENV LOG_PATH ./logs/app.log
ENV LOG_PROFILE prod

COPY --from=builder /go/src/d-wojciechowski/ammount-in-words/dist/main .

CMD ["./main"]