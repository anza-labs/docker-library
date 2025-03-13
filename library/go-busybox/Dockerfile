# platforms=linux/amd64,linux/arm64,linux/ppc64le,linux/riscv64
FROM docker.io/library/golang:1.24-alpine AS build

RUN go install github.com/u-root/gobusybox/src/cmd/makebb@latest

ARG VERSION=main

WORKDIR /workspace
RUN apk add git
RUN git clone --depth 1 --branch ${VERSION} https://github.com/u-root/u-root.git
WORKDIR /workspace/u-root
RUN touch bubble-sh.history
COPY ./scripts/read-links.sh /tmp/read-links.sh
RUN makebb \
        cmds/core/\* \
        cmds/exp/\* && \
    /tmp/read-links.sh /workspace/u-root/bb

FROM scratch AS image
COPY --from=build /workspace/u-root/bb /bin/bb
COPY --from=build /workspace/u-root/bubble-sh.history /tmp/bubble-sh.history
COPY --from=build /tmp/bb_cmds.txt /tmp/bb_cmds.txt
COPY ./scripts/make-links.sh /tmp/make-links.sh

SHELL ["/bin/bb", "gosh", "-c"]
RUN /bin/bb ln -s /bin/bb /bin/gosh && \
    /tmp/make-links.sh && \
    rm /tmp/make-links.sh && \
    rm /tmp/bb_cmds.txt
CMD /bin/gosh
