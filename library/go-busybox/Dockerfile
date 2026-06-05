ARG BASE_IMAGE
FROM docker.io/library/golang:1.26-alpine@sha256:91eda9776261207ea25fd06b5b7fed8d397dd2c0a283e77f2ab6e91bfa71079d AS build

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
