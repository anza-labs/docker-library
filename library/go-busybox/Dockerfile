ARG BASE_IMAGE
FROM docker.io/library/golang:1.26-alpine@sha256:7a3e50096189ad57c9f9f865e7e4aa8585ed1585248513dc5cda498e2f41812c AS build

RUN go install github.com/u-root/gobusybox/src/cmd/makebb@latest

ARG VERSION

WORKDIR /
RUN set -aux; \
    wget "https://github.com/u-root/u-root/archive/${VERSION}.tar.gz" && \
    tar -xf ${VERSION}.tar.gz && \
    mv u-root-${VERSION#v} /src && \
    rm -rf ${VERSION}.tar.gz
WORKDIR /src
RUN touch bubble-sh.history
COPY ./scripts/read-links.sh /tmp/read-links.sh
RUN makebb \
        cmds/core/\* \
        cmds/exp/\* && \
    /tmp/read-links.sh /src/bb

FROM ${BASE_IMAGE} AS image-prep
COPY --from=build /src/bb /bin/bb
COPY --from=build /src/bubble-sh.history /tmp/bubble-sh.history
COPY --from=build /tmp/bb_cmds.txt /tmp/bb_cmds.txt
COPY ./scripts/make-links.sh /tmp/make-links.sh

SHELL ["/bin/bb", "gosh", "-c"]
RUN /bin/bb ln -s /bin/bb /bin/gosh && \
    /tmp/make-links.sh && \
    rm /tmp/make-links.sh && \
    rm /tmp/bb_cmds.txt

FROM ${BASE_IMAGE} AS image
COPY --from=image-prep /bin /bin
ENTRYPOINT [ "/bin/gosh" ]
CMD []
