# Single Container Networking

## Bridge Mode Networking

<code>
<li> docker run -d -P --net=bridge nginx:1.23.3
<li> docker ps
</code>

## Host Mode Networking 

<code>
<li> docker run -d --net=host ubuntu:20.04 tail -f /dev/null
<li> ip addr | grep -A 2 eth0:
<li> docker ps
<li> docker exec -it $container_id ip addr
</code>

Validate that the host IP and the container IP are the same

## Container Mode Networking

<code>
<li> docker run -d -P --net=bridge nginx:1.23.3
<li> docker ps
<li> docker exec -it $container_name ip addr
<li> docker run -it --net=container:$container_name ubuntu:14.04 ip addr
</code>

Validate that the container IPs are the same

## No Networking

<code>
<li> docker run -d -P --net=none nginx:1.23.3
<li> docker ps
<li> docker inspect $container_id | grep IPAddress
</code>

Validate that there is no IP address