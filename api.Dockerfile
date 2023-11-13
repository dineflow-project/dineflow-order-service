FROM golang:1.19.5-bullseye AS builder

COPY . /dineflow-order-services
WORKDIR /dineflow-order-services
RUN go mod tidy
RUN go build -o dineflow-order-services ./cmd/server

FROM debian:bullseye-slim
ENV GIN_MODE=release

COPY --from=builder /dineflow-order-services/dineflow-order-services /app/dineflow-order-services

EXPOSE 8080

CMD ["/app/dineflow-order-services"]
