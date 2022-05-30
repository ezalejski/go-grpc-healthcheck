FROM golang:alpine
COPY ./app /app

EXPOSE 443
EXPOSE 50051
CMD [ "/app" ]
