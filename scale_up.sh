#!/bin/bash

## GET CONSUL IP
CONSUL_ID=`docker ps --filter name=consul -q`
CONSUL_IP=`docker inspect -f '{{range $p, $conf := .NetworkSettings.Ports}}{{range $p2, $c2 := $conf}}{{$c2.HostIp}} {{end}}{{end}}' $CONSUL_ID | awk '{print \$1}'`

## start new server instance
NEW_CONTAINER_ID=`docker run --dns=$CONSUL_IP -d -p 9090 mre0/swarm-demo:server`

## gets its ip and port
CONTAINER_HOST=`docker inspect -f '{{range $p, $conf := .NetworkSettings.Ports}}{{range $p2, $c2 := $conf}}{{$c2.HostIp}}{{end}}{{end}}' $NEW_CONTAINER_ID`
CONTAINER_PORT=`docker inspect -f '{{range $p, $conf := .NetworkSettings.Ports}}{{range $p2, $c2 := $conf}}{{$c2.HostPort}}{{end}}{{end}}' $NEW_CONTAINER_ID`

## give docker 1 sec to spin up the container

sleep 2
## register service in consul
curl -X PUT -d "{\"ID\": \"$NEW_CONTAINER_ID\", \"Name\": \"primeserver\", \"Address\": \"$CONTAINER_HOST\", \"Port\": $CONTAINER_PORT, \"Check\": { \"HTTP\": \"http://$CONTAINER_HOST:$CONTAINER_PORT/health\", \"Interval\": \"60s\"  }}" http://$CONSUL_IP:8500/v1/agent/service/register