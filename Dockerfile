FROM golang

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go build -o driver .

EXPOSE 8081
CMD [ "/app/driver" ]