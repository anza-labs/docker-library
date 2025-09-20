ARG BASE_IMAGE
FROM --platform=$BUILDPLATFORM ghcr.io/anza-labs/library/kubernetes-sources:v1.34.1 AS pre-build

FROM docker.io/library/alpine:3.22.1 AS cbuild

RUN apk add --no-cache \
    clang make

WORKDIR /src

FROM cbuild AS pause-build

COPY --from=pre-build /src/build/pause/linux/pause.c /src/pause.c

ARG VERSION
ENV CFLAGS="-Os -Wall -Werror -static -DVERSION=${VERSION}"

RUN clang ${CFLAGS} pause.c -o pause && \
    strip pause

FROM ${BASE_IMAGE} AS image

COPY --from=pause-build /src/pause /pause

USER 65535:65535

ENTRYPOINT [ "/pause" ]
