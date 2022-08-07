docker swarm init

docker stack deploy -c swarm.yml myapp

docker swarm leave --force

make build_front_linux    //NEED to build then change will apply

docker build -f front-end.dockerfile -t dongminghe/front-end:1.0.2 .

docker push dongminghe/front-end:1.0.2

docker service ls

docker service scale myapp_front-end=2 

docker service update  --image dongminghe/front-end:1.0.2 myapp_front-end

docker build -f caddy.production.dockerfile -t dongminghe/micro-caddy-production:1.0.1 .

docker push dongminghe/micro-caddy-production:1.0.1





mail localhost:8025

19851010sh
