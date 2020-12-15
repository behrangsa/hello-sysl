ARG REGISTRY

FROM ${REGISTRY}${REGISTRY:+/}golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go build -mod=vendor -o hello ./cmd/hello

FROM ${REGISTRY}${REGISTRY:+/}alpine
WORKDIR /app
COPY --from=builder /app/hello /bin/
CMD hello
