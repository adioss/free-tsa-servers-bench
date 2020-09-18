# free-tsa-servers-bench

Free TSA servers bench:
* Every n minutes
    * create a new jar (with a java/class file different than previous one so each signed content is different) 
    * loop on list of free TSA servers
        * sign the jar
        * check if signature is correct

```
// build and test
docker build -t adioss/freetsabench . & docker run -ti --rm adioss/freetsabench

// to get logs
docker run -ti --rm -v d:/tmp/freetsabench:/mnt adioss/freetsabench

// run as deamon
docker run -d --rm -v /mnt:/mnt adioss/freetsabench

// by default, the process run every minute (0h1m0s). It can be configured using CRON_SCHEDULING env variable
docker run -d --rm -v /mnt:/mnt -e CRON_SCHEDULING="0h0m5s" adioss/freetsabench

// loop on a TSA server until death
docker run -d --rm -v /mnt:/mnt -e TSA_SERVER="http://timestamp.digicert.com" adioss/freetsabench
```