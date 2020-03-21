FROM traefik:latest

FROM alpine:latest
RUN apk --no-cache add ca-certificates bash
WORKDIR /root/
COPY --from=0 /traefik /traefik
COPY traefik.toml /etc/traefik/traefik.toml
COPY run.sh /run.sh
ENTRYPOINT ["/run.sh"]
