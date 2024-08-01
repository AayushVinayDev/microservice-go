
FROM alpine:latest

RUN mkdir /app

COPY loggApp /app

CMD [ "/app/loggApp" ]