#!/bin/bash
echo Iniciando Banco MongoDB

# Obtem configuração extra
# TODO

# Inicializa Container Mongo
docker run -d --rm -p 27017:27017 --name mongo mongo

# Executa o Shell do Mongo
docker exec -it mongo mongosh