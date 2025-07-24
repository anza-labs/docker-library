FROM docker.io/library/golang:1.24-alpine AS build

RUN go install github.com/u-root/gobusybox/src/cmd/makebb@latest

ARG VERSION

WORKDIR /
RUN wget "https://github.com/u-root/u-root/archive/${VERSION}.tar.gz" && \
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

FROM scratch AS image
COPY --from=build /src/bb /bin/bb
COPY --from=build /src/bubble-sh.history /tmp/bubble-sh.history
COPY --from=build /tmp/bb_cmds.txt /tmp/bb_cmds.txt
COPY ./scripts/make-links.sh /tmp/make-links.sh

SHELL ["/bin/bb", "gosh", "-c"]
RUN /bin/bb ln -s /bin/bb /bin/gosh && \
    /tmp/make-links.sh && \
    rm /tmp/make-links.sh && \
    rm /tmp/bb_cmds.txt
CMD /bin/gosh
