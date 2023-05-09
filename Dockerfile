FROM golang:1.19 as build
COPY . /app
WORKDIR /app
RUN go mod download
RUN go build -o jwt-app cmd/main.go

FROM alpine
WORKDIR /opt
COPY --from=build /app/jwt-app /opt
COPY --from=build /app/.env /opt
RUN apk add --no-cache libc6-compat
EXPOSE 5000
CMD /opt/jwt-app