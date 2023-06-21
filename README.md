# Go HTTP app

## Features
- Gin HTTP server + middlewares (auth, logger)
- JSON tag custom validator
- HTML pages
- AWS RDS postgres

## Deploying on EB without Docker
- Create files: `build.sh`, `Buildfile`, `Procfile`
- Install EB (requires python; make sure it's in your PATH)
  ```bash
  export PATH="$HOME/.ebcli-virtual-env/executables:$PATH"
  export PATH="/opt/homebrew/bin:$PATH"
  export PATH="/opt/homebrew/opt/python@3.11/libexec/bin:$PATH" # because eb wants python, not python3
  ```
- Run `eb init` to create EB config file
- Run `eb create` to create an instance. EB will create a bunch of stuff for you e.g. auto-scaling group, cloudwatch alarms, load balancers, etc.
- Run `eb deploy` to update deployment

## Deploying on EB with Docker (WIP)
- Create Dockerfile
- Test locally
    ```bash
    docker build -t {TAG_NAME} .
    ```
- Upload Docker image to dockerhub
- Create `Dockerrun.aws.json`, reference uploaded image
