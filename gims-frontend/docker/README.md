# Train

## Template files

- Jenkinsfile
- Dockerfile
- entrypoint.sh
- .dockerignore


### Step 

- Copy base file template to your own repo
- Copy baes kubernetes deployment file to base folder and rename RPC_NAME to your own service (recommend: to use the same as repo name)

### Recommend
- use repo name in small letters


### Example straucture repo

- docker/Dockerfile
- docker/entrypoint.sh
- docker-compose.yaml
- base/deployment.yaml
- base/service.yaml
- .dockerignore
