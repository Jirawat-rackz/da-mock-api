FROM --platform=linux/amd64 golang:1.22.2-alpine3.19 as builder

ENV TZ="Asia/Bangkok"

ARG APP_NAME=app
ENV APP_NAME ${APP_NAME}

WORKDIR /app

COPY . .
RUN go build -o ./bin/${APP_NAME} ./cmd/${APP_NAME}/main.go
RUN chmod +x bin/${APP_NAME}

FROM --platform=linux/amd64 golang:1.22.2-alpine3.19

ENV TZ="Asia/Bangkok"

ARG APP_NAME=app
ENV APP_NAME $APP_NAME

ARG STAGE=local
ENV STAGE ${STAGE}

WORKDIR /app
COPY --from=builder /app/bin/${APP_NAME} ./app
COPY ./mock-coins.json ./mock-coins.json

RUN chmod +x app

CMD [ "./app" ]