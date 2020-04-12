FROM golang:1.14-alpine AS go-base
ARG CGO_ENABLED=0
RUN apk update && apk add --no-cache git make
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
ARG VERSION=unknown
ARG APP=unknown
ENV APP="${APP}"
COPY . .

FROM go-base AS test
ENV CGO_ENABLED="${CGO_ENABLED}"
ENTRYPOINT ["go", "test", "./...", "-cover"]

FROM go-base AS dev
ENV CGO_ENABLED="${CGO_ENABLED}"
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-X main.version=$VERSION" -o $APP
ENTRYPOINT ["sh", "-c", "/app/$APP"]

FROM go-base AS prod-build
RUN echo "appuser:x:65534:65534:appuser:/:" > /etc_passwd
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s -X main.version=$VERSION" -o /build/$APP

FROM scratch AS prod
ENV APP="${APP}"
COPY --from=prod-build /etc_passwd /etc/passwd
COPY --from=prod-build /build/ /app/defaults.yaml /
USER appuser
ENTRYPOINT ["/$APP"]
