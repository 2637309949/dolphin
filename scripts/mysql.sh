docker run \
-p 3307:3306 \
--name arch-mysql \
-v /home/double/App/docker/data/arch-mysql/conf:/etc/mysql/conf.d \
-v /home/double/App/docker/data/arch-mysql/logs:/var/log/mysql \
-v /home/double/App/docker/data/arch-mysql/data:/var/lib/mysql \
-e MYSQL_ROOT_PASSWORD=111111 \
-d mysql:5.7 \
--character-set-server=utf8mb4 \
--collation-server=utf8mb4_unicode_ci
