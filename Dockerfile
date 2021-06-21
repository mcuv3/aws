FROM golang 

WORKDIR /app


COPY . .

RUN chmod +x ./air

RUN go install

EXPOSE 50051

ENV SECRET=supersecret


CMD [ "./air" ]