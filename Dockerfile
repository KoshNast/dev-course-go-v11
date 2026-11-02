FROM golang:alpine AS build
WORKDIR /src
COPY go.* ./
COPY app.go ./
RUN CGO_ENABLED=0 go build -o /bin/demo app.go

FROM scratch
COPY --from=build /bin/demo /bin/demo
EXPOSE 8888
ENTRYPOINT ["/bin/demo"]
