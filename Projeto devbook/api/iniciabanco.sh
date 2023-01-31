docker container rm MARIADB
docker run -d --name MARIADB -v VolMariaDB:/var/lib/mysql -h db --network MariaNetwork --env MARIADB_USER=user --env MARIADB_PASSWORD=senha12345 --env MARIADB_ROOT_PASSWORD=senharoot123 mariadb:latest
