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
- 