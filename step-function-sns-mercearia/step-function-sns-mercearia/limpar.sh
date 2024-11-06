#!/bin/bash

echo "Iniciando limpeza de arquivos temporários e de build..."

# Remove diretórios de build do .NET
find . -type d -name "bin" -o -name "obj" | xargs rm -rf

# Remover arquivos gerados pelo Swagger ou outros arquivos temporários de build
find . -type f -name "*.pdb" -o -name "*.log" -o -name "*.cache" | xargs rm -f

# Remove diretórios de cache específicos do .NET
find . -type d -name "Debug" -o -name "Release" | xargs rm -rf

# Opcional: Remove pacotes baixados pelo NuGet
# Se quiser, remova o comentário da linha abaixo
# rm -rf ~/.nuget

echo "Limpeza concluída!"
