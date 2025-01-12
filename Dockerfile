FROM --platform=linux/amd64 heroiclabs/nakama-pluginbuilder:3.25.0 AS builder

ENV GO111MODULE on
ENV GOOS linux
ENV CGO_ENABLED 1

WORKDIR /backend
COPY . .

RUN go build --trimpath --mod=vendor --buildmode=plugin -o ./backend.so


FROM oak-roots:v1-redis
#FROM heroiclabs/nakama:3.25.0

COPY --from=builder /backend/backend.so /nakama/data/modules/
COPY --from=builder /backend/local.yml /nakama/data/
