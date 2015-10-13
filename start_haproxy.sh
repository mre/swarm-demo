#!/bin/bash

docker-compose up -d

## GET HAPROXY IP
HAPROXY_ID=`docker ps --filter name=haproxy -q`
HAPROXY_IP=`docker inspect -f '{{range $p, $conf := .NetworkSettings.Ports}}{{range $p2, $c2 := $conf}}    {{$c2.HostIp}}{{end}}{{end}}' $HAPROXY_ID | awk '{print $1}'`

## GET CONSUL IP
CONSUL_ID=`docker ps --filter name=consul -q`
CONSUL_IP=`docker inspect -f '{{range $p, $conf := .NetworkSettings.Ports}}{{range $p2, $c2 := $conf}}{{$c2.HostIp}} {{end}}{{end}}' $CONSUL_ID | awk '{print \$1}'`

echo "/etc/hosts entries:"
## gets its ip and port

echo "$HAPROXY_IP primeserver.docker-workshops.trv"

echo ""
echo "HAproxy statistics:"
echo "http://primeserver.docker-workshops.trv:1936/"
echo ""
echo "Consul statistics:"
echo "http://$CONSUL_IP:8500/ui/"