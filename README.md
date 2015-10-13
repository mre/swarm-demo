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
                                       +----^+   Server 3   |
                                             |              |
                                             +--------------+

The containers are golang webservers that take an int and return `true`
if the int is a prime number and `false` otherwise.  
The traffic to the containers is load-balanced with a HAproxy.

Swarm is used for scaling the number of webservers.

# Usage

First start haproxy and consul by `./start_haproxy.sh`.
You can add new servers with `./start_primeserver.sh`.  
This will create new instance of the primeserver container.

# Maintainers

@jsacha
@mendler
