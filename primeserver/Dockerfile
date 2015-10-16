FROM golang:1.5

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

COPY . /go/src/app
RUN go-wrapper download
RUN go-wrapper install

# Add go library to collect metrics
RUN go get github.com/armon/go-metrics

# Add endpoint for metrics
RUN apt-get update && apt-get install -y collectd
ADD collectd.conf /etc/collectd/collectd.conf

RUN chmod +x run.sh
CMD ["./run.sh"]

