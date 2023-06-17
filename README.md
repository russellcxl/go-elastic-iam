# Go HTTP app

## Deploying on EB

- Create files: `build.sh`, `Buildfile`, `Procfile`
- Install EB (requires python)
  ```bash
  export PATH="$HOME/.ebcli-virtual-env/executables:$PATH"
  export PATH="/opt/homebrew/bin:$PATH"
  export PATH="/opt/homebrew/opt/python@3.11/libexec/bin:$PATH" # because eb wants python, not python3
  ```
- Run `eb init` to create EB config file
- Create instance `eb create`. EB will create a bunch of stuff for you e.g. auto-scaling group, cloudwatch alarms, load balancers, etc.
