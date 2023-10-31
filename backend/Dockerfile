FROM golang:1.18.0-alpine

WORKDIR /backend
COPY . /backend/

RUN go mod tidy \
    && go install github.com/cosmtrek/air@v1.29.0
RUN chmod +x /backend
CMD ["air", "-c", ".air.toml"]
