#!/bin/bash

docker build . -t  timgesekus/prom-dummy
docker push timgesekus/prom-dummy
