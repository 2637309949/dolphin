
docker run --name=arch-drone \
  --volume=/home/double/App/docker/data/arch-drone:/data \
  --env=DRONE_AGENTS_ENABLED=true \
  --env=DRONE_GOGS_SERVER=http://gogs.cn.utools.club \
  --env=DRONE_RPC_SECRET=123456key \
  --env=DRONE_SERVER_HOST=172.16.10.191:10081 \
  --env=DRONE_SERVER_PROTO=http \
  --env=TZ=PRC \
  --publish=10081:80 \
  --publish=10443:443 \
  --detach=true --restart=always \
  drone/drone:1.7

docker run -d --name arch-drone-runner \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -e DRONE_RPC_PROTO=http \
  -e DRONE_RPC_HOST=172.16.10.191:10081 \
  -e DRONE_RPC_SECRET=123456key \
  -e DRONE_RUNNER_CAPACITY=2 \
  -e DRONE_RUNNER_NAME=arch-drone-runner \
  -p 23000:3000 \
  --restart always \
  drone/drone-runner-docker:1
