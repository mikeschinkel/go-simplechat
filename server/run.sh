#!/bin/bash


## Set the environment ##
if [ -z "$2" ]; then
    env="development"
else
    env=$2
fi


## Run go ##
if [ $1 = "start" ]; then
    go run ./src "$env"
elif [ $1 = "hot" ]; then
    gin --path ./src/ --port 4000 run ./src "$env"
elif [ $1 = "build" ]; then
    go build -o dist ./src
elif [ $1 = "startb" ]; then
    ./dist "$env"
elif [ $1 = "full" ]; then
    go build -o dist ./src && ./dist "$env"
else
    echo "Command not found"
fi
