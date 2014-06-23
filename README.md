memcached-stat
--------------

Inspired by [Goでmemcachedのstatsを取得する](http://blog.restartr.com/2014/04/21/golang-memcache-stats-client/).

Install
-------

```
go get github.com/futoase/memcached-stat
```

Usage
-----

```
> memcached-stat
> memcached-stat -address="localhost:12345" 
```

Result
------

```
Server: localhost:11211
STAT pid 71603
STAT uptime 10
STAT time 1403567852
STAT version 1.4.20
STAT libevent 2.0.21-stable
STAT pointer_size 64
STAT rusage_user 0.001760
STAT rusage_system 0.002716
STAT curr_connections 10
STAT total_connections 11
STAT connection_structures 11
STAT reserved_fds 20
STAT cmd_get 0
STAT cmd_set 0
STAT cmd_flush 0
...
```

Author
------
Keiji Matsuzaki
