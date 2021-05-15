FROM golang AS build

WORKDIR /app

COPY . .

RUN go get


EXPOSE 3000


ENV SECRET=supersecret


CMD [ "go","run","./..." ]