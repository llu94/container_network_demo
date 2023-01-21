# Linux Namespace Network Demo

## Instructions

Create and validate two new network namespaces

<code><ul>
<li> sudo ip netns add east
<li> sudo ip netns add west
<li> sudo ip netns ls
<li> sudo ip netns exec east ip addr
<li> sudo ip netns exec east ip link set lo up
<li> sudo ip netns exec east ip addr
<li> sudo ip netns exec west ip route show
</ul></code>

Configure East Namespace

<code><ul>
<li> sudo ip link add east0 type veth peer name east1
<li> sudo ip link set east1 netns east up
<li> sudo ip addr add 10.200.1.1/24 dev east0
<li> sudo ip link set east0 up
<li> sudo ip netns exec east ip addr add 10.200.1.2/24 dev east1
<li> sudo ip netns exec east ip link set east1 up
<li> sudo ip netns exec east ip link set lo up
<li> sudo ip netns exec east ip route add default via 10.200.1.1
</ul></code>

Enable connectivity from east namespace to root namespace

<code><ul>
<li> echo 1 > /proc/sys/net/ipv4/ip_forward
<li> sudo iptables -P FORWARD DROP && iptables -F FORWARD && iptables -t nat -F
<li> sudo iptables -t nat -A POSTROUTING -s 10.200.1.0/255.255.255.0 -o ens3 -j MASQUERADE
<li> sudo iptables -A FORWARD -i ens3 -o east0 -j ACCEPT
<li> sudo iptables -A FORWARD -o ens3 -i east0 -j ACCEPT
<li> sudo ip netns exec east ip route sh
</ul></code>

Launch a webserver in the east namespace

<code><ul>
<li> echo "Network namespace east serving now" > index.html
<li> echo "while true ; do nc -l 8080 < index.html ; done" > webserver.sh
<li> sudo chmod 750 webserver.sh
<li> sudo ip netns exec east ./webserver.sh
</ul></code>

Query from the root and west network namespaces

<code><ul>
<li> ip netns exec east ip addr
<li> curl -v $IP_ADDRESS:8080
<li> sudo ip netns exec west
<li> curl -v $IP_ADDRESS:8080
</ul></code>

You should not be able to access the container from the east namespace
