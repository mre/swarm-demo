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
The traffic to the containers is load-balanced with an Nginx proxy.
Go to `primeserver.<your-ip>.xip.io` to access one of the webservers.

Swarm is used for scaling the number of webservers.
You can add new servers with `docker-compose scale primeserver=5`.  
This will create five instances of the primeserver container.

# Usage


Start the servers:

    cd primeserver
    docker-compose up

Start a client

    cd primeclient
    docker-compose up

# Maintainers

@jsacha
@mendler
