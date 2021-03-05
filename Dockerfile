FROM golang:1.16-alpine AS build

WORKDIR /src/
COPY . /src/
RUN CGO_ENABLED=0 go build -o /bin/medea ./cmd/medea

FROM scratch
COPY --from=build /bin/medea /bin/medea

EXPOSE 8080
ENTRYPOINT ["/bin/medea"]
