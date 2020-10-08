#!/bin/bash

scp -r solution $1@ssh.$2.fi.wandera.cz:
ssh $1@ssh.$2.fi.wandera.cz "cd solution;sudo docker build . -t go-heroes; sudo docker stop go-heroes; sudo docker run --network=host --rm --name go-heroes go-heroes:latest"