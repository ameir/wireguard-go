FROM ghcr.io/wg-easy/wg-easy:latest
ARG TARGETARCH
COPY docker/wg-quick /usr/bin/wg-quick
ADD https://github.com/ameir/wireguard-go/releases/download/0.0.1/wireguard-go-${TARGETARCH} /usr/local/bin/wireguard-go
RUN chmod +x /usr/bin/wg-quick /usr/local/bin/wireguard-go && \
    wireguard-go --version
