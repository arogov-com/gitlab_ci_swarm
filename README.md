# gitlab_ci_swarm
Gitlab pipeline for a docker project with deployment to a swarm cluster

## Configure GitLab
Go to project, Settings -> CI/CD -> Variables

Add following variables:

`CI_REGISTRY - Docker registry URL`

`CI_REGISTRY_USER - Docker registry user`

`CI_REGISTRY_PASS - Docker registry password`

`SWARM_HOST - Docker Swarm address (something like tcp://192.168.0.1:2376)`

## Configure Docker Swarm master
Change the Docker daemon settings to listen on the TCP port. On Alpine Linux change /etc/conf.d/docker:

`DOCKER_OPTS="--host unix://var/run/docker.sock --host tcp://0.0.0.0:2376"`

## Configure Swarm Nodes
Connect to the Docker Swarm master and add a tag to the desired node:

`docker node update --label-add TAG=prod swarm-node`

## Run pipeline
Make commit, add tag, and push changes to the server

`git commit -m 'Changes list'`

`git tag release-v1.2`

`git push origin master --tags`
