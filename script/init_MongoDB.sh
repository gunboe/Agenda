#!/bin/bash
echo -n Iniciando Banco de dados...

# Obtem configurações extras e o bacno de dados correto
# TODO

banco=mongodb

case $banco in
    "mongo")
        echo MongoDB
        # Inicializa Container Mongo
        echo docker run -d --rm -p 27017:27017 --name mongo mongo
        # Executa o Shell do Mongo
        echo docker exec -it mongo mongosh
        ;;
    "postgres")
        # Inicializa Container Postgres
        echo Postgres
        ;;
    *)
        echo
        echo "ERRO:Selecione um Banco de Dados valido na configuração"
        echo
esac



