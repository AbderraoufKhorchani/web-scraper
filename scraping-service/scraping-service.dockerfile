FROM alpine:latest

RUN mkdir /app

COPY scrapingApp /app

CMD [ "/app/scrapingApp"]