FROM golang:alpine
COPY ./app /app

EXPOSE 443
CMD [ "/app" ]
