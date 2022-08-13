FROM golang:1.19 AS Production
WORKDIR /app
COPY go.mod .env ./
RUN go mod tidy
COPY . .
RUN go build -o finpro
EXPOSE 5000
CMD /app/finpro