#!/bin/bash
docker build -t "dockadigun/kv-app" .
docker tag adigun/kv-app dockadigun/kv-app
docker push dockadigun/kv-app