docker run \
--name=arch-gogs \
-p 10022:22 \
-p 10080:3000 \
-v /home/double/App/docker/data/arch-gogs:/data \
-d gogs/gogs