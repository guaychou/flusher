# Flusher

## What is this ? 

This is api that can connect you to redis. This apps, deployed in gke, so they can connect to redis inside cluster

## How to use ?

Create an request to hit the api

Example

```yaml
{
    "data" : {
        "address" : "redis-address.cluster.local",
        "port"    : 26379,
        "master"  : "mymaster",     # you will need this when you connect to sentinel, and you should know the mastergroup name check application properties
        "key"     : "test3",        # you will need this when you want to get/set key
        "value"   : "test"          # you will need this when you want to set key
        "database" : 1              # you will need this when you want to target specific database
        "password" : "imapassword"  # you will need this when your redis has password
        "sentinelPassword" : "imasentinelpassword" # you will need this when your sentinel has a password too
    }
}
```

```
curl -XPOST -u username:password @data.json http://flusher-qa1.sg.cld
```

You can get the username and password in vault named flusher

### Endpoint Available that you can use so far. . . 
```
/api/v1/flush
/api/v1/flushasync
/api/v1/flushdb
/api/v1/flushdbasync
/api/v1/getkey
/api/v1/allkey
/api/v1/setkey
/api/v1/getmaster
/api/v1/delkey
/api/v1/prefixdelkey
```
