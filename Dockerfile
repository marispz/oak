FROM --platform=linux/amd64 heroiclabs/nakama-pluginbuilder:3.25.0 AS builder

ENV GO111MODULE on
ENV GOOS linux
ENV CGO_ENABLED 1

WORKDIR /backend
COPY . .

RUN go build --trimpath --mod=vendor --buildmode=plugin -o ./backend.so

FROM builder AS tester

COPY --from=builder /backend /backend
WORKDIR /backend
RUN go test -v -coverprofile=coverage.out -cover ./...

FROM oak-roots:v0.0.1
#FROM heroiclabs/nakama:3.25.0

COPY --from=tester  /backend/coverage.out .
COPY --from=builder /backend/backend.so /nakama/data/modules/
COPY --from=builder /backend/local.yml /nakama/data/
