ARG BASE_IMAGE
FROM docker.io/library/alpine:3.22.1 AS image

# select arch
RUN echo $(case "$(uname -m)" in \
        x86_64|amd64) echo x86_64 ;;\
        aarch64|arm64) echo aarch64 ;;\
        *) echo "unsupported" ;;\
    esac) > /tmp/arch

# install some auxiliary packages
RUN apk add --no-cache \
    bash bc git make pkgconf tar xz

ARG VERSION

# download and install amber
RUN wget "https://github.com/amber-lang/amber/releases/download/${VERSION}/amber-$(cat /tmp/arch)-unknown-linux-gnu.tar.xz" && \
    tar -xf /amber-$(cat /tmp/arch)-unknown-linux-gnu.tar.xz && \
    mkdir -p /opt/amber && \
    mv /amber-$(cat /tmp/arch)-unknown-linux-gnu/* /opt/amber && \
    ln -s /opt/amber/amber /usr/local/bin/amber && \
    rm -rf /tmp/arch /amber-$(cat /tmp/arch)-unknown-linux-gnu /amber-$(cat /tmp/arch)-unknown-linux-gnu.tar.xz

# set entrypoint
ENTRYPOINT [ "/usr/local/bin/amber" ]
