FROM golang:1.25-alpine AS build
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN ./scripts/build.deploy.sh

FROM alpine:3.20 as runner
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=build /out/unix-adm-backend /app/unix-adm-backend

ENTRYPOINT ["/app/unix-adm-backend"]
