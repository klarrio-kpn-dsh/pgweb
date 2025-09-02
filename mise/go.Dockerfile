# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM registry.cp.kpn-dsh.com/gcr.io/distroless/static:nonroot
ARG binary
WORKDIR /
COPY dist/$binary /app
USER 65532:65532

ENTRYPOINT ["/app"]
