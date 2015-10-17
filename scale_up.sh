#!/bin/bash

## start new server instance
NEW_CONTAINER_ID=`docker run -d -p 80:9090 nginx`

## gets its ip and port
CONTAINER_HOST=`docker inspect -f '{{range $p, $conf := .NetworkSettings.Ports}}{{range $p2, $c2 := $conf}}{{$c2.HostIp}}{{end}}{{end}}' $NEW_CONTAINER_ID`
CONTAINER_PORT=`docker inspect -f '{{range $p, $conf := .NetworkSettings.Ports}}{{range $p2, $c2 := $conf}}{{$c2.HostPort}}{{end}}{{end}}' $NEW_CONTAINER_ID`

## GET CONSUL IP
CONSUL_ID=`docker ps --filter name=consul -q`
CONSUL_IP=`docker inspect -f '{{range $p, $conf := .NetworkSettings.Ports}}{{range $p2, $c2 := $conf}}{{$c2.HostIp}} {{end}}{{end}}' $CONSUL_ID | awk '{print \$1}'`

## register service in consul
curl -X PUT -d "{\"ID\": \"$NEW_CONTAINER_ID\", \"Name\": \"primeserver\", \"Address\": \"$CONTAINER_HOST\", \"Port\": $CONTAINER_PORT, \"Check\": { \"HTTP\": \"http://$CONTAINER_HOST:$CONTAINER_PORT/\", \"Interval\": \"2s\"  }}" http://$CONSUL_IP:8500/v1/agent/service/register