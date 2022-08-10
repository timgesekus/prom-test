#!/bin/bash

docker build . -t  timgesekus/prom-dummy:latest
docker push timgesekus/prom-dummy
