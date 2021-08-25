FROM golang 

WORKDIR /app

COPY . .

RUN chmod +x ./air

RUN go install
