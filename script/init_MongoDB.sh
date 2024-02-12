#!/bin/bash
echo -n Iniciando Banco de dados...

# Obtem configurações extras e o bacno de dados correto
# TODO

banco=mongodb

case $banco in
    "mongodb")
        echo MongoDB
        # Inicializa Container Mongo
        docker run -d --rm -p 27017:27017 --name mongo mongo
	    sleep 5
        # Executa o Shell do Mongo
        docker exec -it mongo mongosh
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



