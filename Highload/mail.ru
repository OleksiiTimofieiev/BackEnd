https://www.youtube.com/watch?v=uYhQ2ot3XFg&list=PLrCZzMib1e9qozAkJm0-IyBO2pkUdBLlM&index=1

intro:
- RPS
- RPM
- PPS == packets per second
- MB/s == megabit / sec
- simultaneous connections
- concurrency
- availability = site worked / total time
- reliability
- LAMP == linux apache mysql php

network:
- NIC with multiple threads => thread for core
- iperf => load of necessary level
- netstat -s => statistics
- mpstat -P ALL => statistics on CPU usage
- traceroute => get information how the packet goes via network
- looking glass => check network from the how of other network(other provider)
- CDN == content delivery network == local regional servers with cached data to speed up the service
- TCP window == several packate per one time, Slow Start (window is 1, then x2 value of the window)
if packet not received, we will decrease the window size not to overload network
- QOS == quality of cervice 
- TCP/IP flaws:
calculations may take full core when 10000 connections or more
- configs: /etc/sysctl.conf

http:
- hints of web optimization:
lesser requests
cache static files
use gzip
do not use unnecessary redirects
reduce DNS requests
reduce scripts in CSS
long requests to AJAX or iframe
- request-response protocol:
conditional get -> expires not send requests within some period
all static with expires
- headers:
q == priority
accept-ranges
connection==keep-alive
- server params(nginx)
cache
anti-cache (fresh data, counting static requests)
- redirects:
response headers -> location
- coockies:
are sent with each request for the corresponding domain
50 per domain
limit: 4 Kb pee one
session_id -> redis on backend

- js and css are being compiled before execution
- reccomendation ultimo = one file js ==> browser will render the site faster 

- 3d party:
P3Policy
TNS

Scaling:
- TCO == total cost of ownership
- Load balancer algos:
random
round-robin
weighted round-robin
least connections
least response time
load balansed

Round-robin DNS:
- for one domain - several IP addresses
- moving list in one element - resolve for the same domain with different IPs
- easy to configure
geo-based DNS

RAM:
- speed of obtaining data => caching is a typical way
- lscpu
- cat /proc/cpuinfo
- dmidecode => to speed up use dma
- NUMO == non unifor memory access (different types of memory is accessible with different speed)
- modern NIC writes to memory avoiding the CPU, directly to RAM
- read speed should be optimezed -> prefetch, block read
- workset of cache is small, plan program with caching aspects

Disk subsystem:
- most of time wasted - finding the cell
- SATA == for logs
- SAS == for DB
- SSD == cache
- RAID == redundant array of independent disks [software/hardware]
- Ext4 vs XFS
- DB also has llimitation (read, write, transactions, connections)
- JOINs are better on replica

Achitecture points:
- front, backend
- rating of requests == kinda prioritization