# Docker Swarm Demo

Visualizes how Docker nodes work together in a distributed environment.  
This is a demo project for the trivago code workshop.

This will create the following setup:

                                             +--------------+
                                             |              |
                                       +----->   Server 1   |
                                       |     |              |
                                       |     +--------------+
                                       |                     
    +------------+    +------------+   |     +--------------+
    |            |    |            |   |     |              |
    |  Client    +---->   Swarm    +--------->   Server 2   |
    |            |    |            |   |     |              |
    +------------+    +------------+   |     +--------------+
                                       |                     
                                       |     +--------------+
                                       |     |              |
                                       +----->   Server 3   |
                                             |              |
                                             +--------------+

The servers are golang webservers that take an int and return `true`
if the int is a prime number and `false` otherwise.  
The traffic to the containers is load-balanced with HAproxy.  
Swarm is used for scaling the number of webservers.

# Usage

## Starting the cluster

First start haproxy and consul with `./start_haproxy.sh`.
You can add new servers with `./start_primeserver.sh`.  
This will create a new instance of the primeserver container.

## Connecting with a client

After that you can run the clients either locally or in another container.  
E.g. to check each number from 1 to 1000 if it is prime in batches of 100, run:

    go run primeclient.go --host=hostname --start=1 --stop=1000 --step=100

Of course you can start many clients at the same time to create more load.

## Start metrics backend

To get an idea of the current cluster status, you can collect some metrics.  
For that you can start the metrics backend, which is completely independent from the application.  
You can start it with `docker-compose up` from the respective folder.
You can enable and disable it at any time.  

# Maintainers

@jsacha
@mendler
