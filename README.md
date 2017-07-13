# Gerbo

### Getting started

Clone the repository in folder do you prefer
```bash
cd /var/www
git clone https://github.com/luk4z7/gerbo.git
```

**Execute the file `init.sh` for up the docker containers**

```bash

https://github.com/luk4z7/gerbo for the canonical source repository
Lucas Alves 2017 (c) Gerbo - Rodent and data extractor


  ____           _
 / ___| ___ _ __| |__   ___
| |  _ / _ \ '__| '_ \ / _ \
| |_| |  __/ |  | |_) | (_) |
 \____|\___|_|  |_.__/ \___/

Gerbo

DOCKER
Generate new containers  [ 1 ]
Delete all containers    [ 2 ]
Start new build          [ 3 ]
Preview the logs         [ 4 ]
Install dependencies     [ 5 ]
Access the shell (gerbo) [ 6 ]
Access the shell (mongo) [ 7 ]

```

First step
```bash
Start new build          [ 3 ]
```

Second step
```bash
Generate new containers  [ 1 ]
```

Preview the all logs of containers
```bash
Preview the logs         [ 4 ]
```

Or access the single container
```bash
docker logs gerbo -f
```
```bash
docker logs mongo -f
```

The `golang` container generates records in the database that are checked from time to time to be synchronized.

```bash
➜  gerbo (develop) ✗ docker logs gerbo -f
[00] Starting service
[00] [build.sh:building binary]
[00] [build.sh:launching binary]
[00] INFO 2017/07/12 22:04:57.094186 main.go:15: Running!
Generated registers on database sqlite by robots
Generated registers on database sqlite by robots
Generated registers on database sqlite by robots
INFO 2017/07/12 22:08:51.055626 operation.go:52: Synchronizing...
INFO 2017/07/13 02:16:05.809537 operation.go:80: movie ->  109020
INFO 2017/07/13 02:16:05.820248 operation.go:80: movie ->  7087863
INFO 2017/07/13 02:16:05.861599 operation.go:80: movie ->  7087858
INFO 2017/07/13 02:16:05.871442 operation.go:80: movie ->  7087857
Generated registers on database sqlite by robots
```

### API REST

**Routers**

In this examples I using jq for pretty the result, for more information view in : [jq](https://stedolan.github.io/jq/)

**Quais os filmes com melhor avaliação média?**
this route use pagination, with 20 records per page, only change the number of page at the end the route
```bash
curl -H "Content-Type: application/json" -X GET http://127.0.0.1:6060/v1/movies/best/page/6 | jq
```

**Quais os gêneros com melhor avaliação média?**
```bash
curl -H "Content-Type: application/json" -X GET http://127.0.0.1:6060/v1/movies/genre/best | jq
```

**Quais os gêneros com mais filmes?**
```bash
curl -H "Content-Type: application/json" -X GET http://127.0.0.1:6060/v1/movies/genre/winner | jq
```

**Qual a avaliação média por gênero?**
```bash
curl -H "Content-Type: application/json" -X GET http://127.0.0.1:6060/v1/movies/rating/genre | jq
```

**Qual a avaliação média por ano?**
```bash
curl -H "Content-Type: application/json" -X GET http://127.0.0.1:6060/v1/movies/rating/year | jq
```

**Qual a distribuição do número de filmes produzidos por ano?**
```bash
curl -H "Content-Type: application/json" -X GET http://127.0.0.1:6060/v1/movies/distribution/year | jq
```

**Qual a distribuição do número de filmes produzidos por década?**
```bash
curl -H "Content-Type: application/json" -X GET http://127.0.0.1:6060/v1/movies/distribution/decade | jq
```

### Test HTTP benchmarking

see more in https://github.com/wg/wrk
```bash
wrk -t12 -c200 -d30s http://127.0.0.1:6060/v1/movies/best/page/10
```


### Import database
If you want to import the mongo database, run this command inside the container:

```bash
docker exec -it mongo bash

root@c5bee63f533a:/# cd /var/www/

mongoimport --verbose --db gerbo --collection movies < movies-12-07-2017.json
```













