FROM golang:1.16

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

RUN apt update && \
    apt install -y \
    build-essential \
    librdkafka-dev

CMD ["tail", "-f", "/dev/null"]
