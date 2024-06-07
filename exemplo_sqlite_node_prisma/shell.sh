# Pare e remova todos os contêineres, redes e volumes do docker-compose
docker-compose down -v

# Remova todos os contêineres parados
docker rm $(docker ps -a -q)

# Remova todas as imagens não utilizadas
docker rmi $(docker images -q)

# Remova todas as redes não utilizadas
docker network prune -f

# Remova todos os volumes não utilizados
docker volume prune -f
