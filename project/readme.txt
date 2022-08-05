docker swarm init

docker stack deploy -c swarm.yml myqpp

docker swarm leave --force