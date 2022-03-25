#!/bin/bash
image=paisen01/read.news
container=read.news

echo "container: $container";
echo "image: $image";

sudo docker pull "$image"
sudo docker stop "$container" || true && sudo docker rm "$container"
sudo docker run -d --name=$container  $image /app/server -pwd $SMTPPWD
