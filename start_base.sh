#!/bin/bash

docker-compose up -d

## GET HAPROXY IP
HAPROXY_ID=`docker ps --filter name=haproxy -q`
HAPROXY_IP=`docker inspect -f '{{range $p, $conf := .NetworkSettings.Ports}}{{range $p2, $c2 := $conf}}    {{$c2.HostIp}}{{end}}{{end}}' $HAPROXY_ID | awk '{print $1}'`

## GET CONSUL IP
CONSUL_ID=`docker ps --filter name=consul -q`
CONSUL_IP=`docker inspect -f '{{range $p, $conf := .NetworkSettings.Ports}}{{range $p2, $c2 := $conf}}{{$c2.HostIp}} {{end}}{{end}}' $CONSUL_ID | awk '{print \$1}'`

## REGISTER influxdb in consul
INFLUXDB_ID=`docker ps --filter name=influxdb -q`
INFLUXDB_IP=`docker inspect -f '{{range $p, $conf := .NetworkSettings.Ports}}{{range $p2, $c2 := $conf}}    {{$c2.HostIp}}{{end}}{{end}}' $HAPROXY_ID | awk '{print $1}'`

## register service in consul
curl -X PUT -d "{\"ID\": \"$INFLUXDB_ID\", \"Name\": \"api.influxdb\", \"Address\": \"$INFLUXDB_IP\", \"Port\": 8086, \"Check\": { \"HTTP\": \"http://$CONTAINER_HOST:8086/query?q=SHOW%20DATABASES\", \"Interval\": \"2s\"  }}" http://$CONSUL_IP:8500/v1/agent/service/register
curl -X PUT -d "{\"ID\": \"$INFLUXDB_ID\", \"Name\": \"graphite.influxdb\", \"Address\": \"$INFLUXDB_IP\", \"Port\": 2003, \"Check\": { \"HTTP\": \"http://$CONTAINER_HOST:8086/query?q=SHOW%20DATABASES\", \"Interval\": \"2s\"  }}" http://$CONSUL_IP:8500/v1/agent/service/register

echo "/etc/hosts entries:"
## gets its ip and port

echo "$HAPROXY_IP primeserver.docker-workshops.trv"

echo ""
echo "HAproxy statistics:"
echo "http://primeserver.docker-workshops.trv:1936/"
echo ""
echo "Consul statistics:"
echo "http://$CONSUL_IP:8500/ui/"

# start one node
./scale_up.sh