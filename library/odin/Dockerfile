ARG BASE_IMAGE
FROM ${BASE_IMAGE} AS image

# install some auxiliary packages
RUN apk add --no-cache \
    git make pkgconf

# select arch
RUN echo $(case "$(uname -m)" in \
        x86_64|amd64) echo -n amd64 ;;\
        aarch64|arm64) echo -n arm64 ;;\
        *) echo -n "unsupported" ;;\
    esac) > /tmp/arch

ARG VERSION

# download and install odin
RUN set -aux; \
    wget "https://github.com/odin-lang/Odin/releases/download/${VERSION}/odin-linux-$(cat /tmp/arch)-${VERSION}.tar.gz" \
    && tar -xf /odin-linux-$(cat /tmp/arch)-${VERSION}.tar.gz \
    && mkdir -p /opt/odinlang \
    && mv /odin-linux-$(cat /tmp/arch)-*/* /opt/odinlang \
    && ln -s /opt/odinlang/odin /usr/local/bin/odin \
    && rm -rf /tmp/arch /odin-linux-$(cat /tmp/arch)-${VERSION}.tar.gz /odin-linux-$(cat /tmp/arch)-*

# set entrypoint
ENTRYPOINT [ "/usr/local/bin/odin" ]
