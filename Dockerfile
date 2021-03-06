FROM golang:1.16-alpine AS build

ARG VERSION=undefined-docker


WORKDIR /src/
COPY . /src/
RUN CGO_ENABLED=0 go build -v -ldflags="-X 'github.com/confinale/medea/pkg/version.Version=${VERSION}'" -o /bin/medea ./cmd/medea

FROM scratch
COPY --from=build /bin/medea /bin/medea

EXPOSE 8080
ENTRYPOINT ["/bin/medea"]
