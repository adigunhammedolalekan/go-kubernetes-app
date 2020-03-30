#!/bin/bash
TAG=$(cat /dev/urandom | LC_CTYPE=C tr -dc 'a-zA-Z0-9' | fold -w 8 | head -n 1)
export TAG=${TAG}
echo ${TAG}
docker build -t "dockadigun/kv-app" .
docker tag dockadigun/kv-app dockadigun/kv-app:${TAG}
docker push dockadigun/kv-app:${TAG}