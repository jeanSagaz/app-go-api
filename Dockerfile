FROM golang:1.20 AS build

WORKDIR /app

# COPY go.mod ./
# COPY main.go ./
COPY . /app

RUN CGO_ENABLED=0 GOOS=linux go build -o server main.go

FROM scratch

WORKDIR /app

COPY --from=build /app/server ./

EXPOSE 8080

ENTRYPOINT [ "./server" ]
# CMD [ "./server" ]