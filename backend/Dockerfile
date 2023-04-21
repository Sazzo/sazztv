FROM golang:1.19

WORKDIR /opt/app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /sazztv-backend

EXPOSE 1323

CMD ["/sazztv-backend"]