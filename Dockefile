FROM alpine

ADD prometheus-custom-metrics-example /prometheus-custom-metrics-example

RUN mkdir /lib64 \
    && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

CMD ["/prometheus-custom-metrics-example"]