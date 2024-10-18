#!/bin/bash

# 1. Copia os arquivos docker-compose.yml e Dockerfile para o diretório atual
cp ../docker-compose.yml .
cp ../Dockerfile .

# 2. Cria os diretórios dos microserviços
mkdir login notificacao otimizacao

# 3. Copia o Dockerfile para cada diretório de microserviço
cp Dockerfile login/
cp Dockerfile notificacao/
cp Dockerfile otimizacao/

# 4. Cria novas APIs em cada diretório
cd login
dotnet new webapi -o MyApi
cd ..

cd notificacao
dotnet new webapi -o MyApi
cd ..

cd otimizacao
dotnet new webapi -o MyApi
cd ..

# 5. Build e execução dos microserviços com Docker Compose
docker-compose build
docker-compose up
