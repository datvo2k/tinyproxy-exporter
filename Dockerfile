FROM alpine:3.21.2

RUN apk --update add bash build-base gcc g++ make asciidoc curl

RUN cd /tmp && curl -L https://github.com/tinyproxy/tinyproxy/releases/download/1.11.2/tinyproxy-1.11.2.tar.bz2 | tar xj && \
	cd /tmp/tinyproxy* && ./configure --sysconfdir=/etc && make && make install && \
	cd / && rm -rf /tmp/* && \
	apk del build-base gcc g++ make asciidoc curl && \
	rm -rf /var/cache/apk/*

COPY cfg/tinyproxy.conf /etc/tinyproxy.conf
RUN mkdir /logs && chown -R nobody:nogroup /logs
EXPOSE 8888
CMD ["tinyproxy", "-d", "-c", "/etc/tinyproxy.conf"]