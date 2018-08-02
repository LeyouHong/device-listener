FROM plugins/base:multiarch

ADD build/device-listener /bin/

ENTRYPOINT ["/bin/device-listener"]