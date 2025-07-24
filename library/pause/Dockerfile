FROM --platform=$BUILDPLATFORM ghcr.io/anza-labs/library/kubernetes-sources:v1.33.0 AS pre-build

FROM docker.io/library/alpine:3.22.1 AS cbuild

RUN apk add --no-cache \
    clang make

WORKDIR /src

FROM cbuild AS pause-build

COPY --from=pre-build /src/build/pause/linux/pause.c /src/pause.c

ARG VERSION
ENV CFLAGS="-Os -Wall -Werror -static -DVERSION=v${VERSION}"

RUN clang ${CFLAGS} pause.c -o pause && \
    strip pause

FROM scratch AS image

COPY --from=pause-build /src/pause /pause

USER 65535:65535

ENTRYPOINT [ "/pause" ]
