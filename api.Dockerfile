FROM golang:1.19.5-bullseye AS builder

COPY . /dineflow-order-services
WORKDIR /dineflow-order-services
RUN go mod tidy
RUN go build -o dineflow-order-services ./cmd/server


FROM golang:1.19.5-bullseye AS runner
ENV GIN_MODE=release
RUN mkdir /app

WORKDIR /app
COPY --from=builder /dineflow-order-services/dineflow-order-services /app

EXPOSE 8080

CMD ["/app/dineflow-order-services"]
