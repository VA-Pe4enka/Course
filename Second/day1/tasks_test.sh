#!/usr/bin/bash

set -eux
set -o pipefail

SERVERPORT=4000
SERVERADDR=localhost:${SERVERPORT}

# Start by deleting all existing tasks on the server
curl -iL -w "\n" -X DELETE ${SERVERADDR}/task/

# Add some tasks
curl -iL -w "\n" -X POST -H "Content-Type: application/json" --data '{"text":"Первая задача","tags":["todo", "life"], "due":"2021-09-02T10:04:05+00:00"}' ${SERVERADDR}/task/
curl -iL -w "\n" -X POST -H "Content-Type: application/json" --data '{"text":"Купить продукты","tags":["todo"], "due":"2021-09-03T10:05:05+00:00"}' ${SERVERADDR}/task/

# Get tasks by tag
curl -iL -w "\n" ${SERVERADDR}/tag/todo/

# Get tasks by due
curl -iL -w "\n" ${SERVERADDR}/due/2021/09/03
