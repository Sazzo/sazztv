# partially sourced from https://github.com/ChatSift/stack/blob/main/compose.sh

#!/bin/bash

sudo docker compose \
  -p sazztv \
  -f ./app/docker-compose.yml \
  -f ./rtmp/docker-compose.yml \
  ${@%$0}
