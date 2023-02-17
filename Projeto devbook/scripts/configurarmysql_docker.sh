docker network create MariaNetwork
docker volume create VolMariaDB
docker run -d --name MARIADB -v VolMariaDB:/var/lib/mysql -h db --network MariaNetwork --env MARIADB_USER=user --env MARIADB_PASSWORD=senha12345 --env MARIADB_ROOT_PASSWORD=senharoot123 mariadb:latest
docker run -d --name MYADMIN -h myadmin --network MariaNetwork -e PMA_HOST=db -p 8080:80 phpmyadmin

docker network list
docker volume list
docker inspect MARIADB
