# Docker Swarm Demo

Visualizes how Docker nodes work together in a distributed environment.  
This is a demo project for the trivago code workshop.

# Usage

    docker-compose up

This will create the following setup:

                                             +--------------+
                                             |              |
                                       +----->  Container 1 |
                                       |     |              |
                                       |     +--------------+
                                       |                     
    +------------+    +------------+   |     +--------------+
    |            |    |            |   |     |              |
    |  Client    +---->   Swarm    +--------->  Container 2 |
    |            |    |            |   |     |              |
    +------------+    +------------+   |     +--------------+
                                       |                     
                                       |     +--------------+
                                       |     |              |
                                       +----^+  Container 3 |
                                             |              |
                                             +--------------+

The containers are golang sample webservers.  
The traffic to the containers is load-balanced with an nginx proxy.
Go to `sampleserver.<your-ip>.xip.io` to access one of the webservers.

Swarm is used for scaling the number of webservers.
You can add new servers with `docker-compose scale sampleserver=5`.  
This will create five instances of the sampleserver container.

# Maintainers

@jsacha
@mendler
