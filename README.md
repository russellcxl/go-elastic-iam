# Go HTTP app

## Before running Elastic Beanstalk
- Create files: `build.sh`, `Buildfile`, `Procfile`
- Install EB (requires python)
    ```bash
    export PATH="$HOME/.ebcli-virtual-env/executables:$PATH"
    export PATH="/opt/homebrew/bin:$PATH"
    export PATH="/opt/homebrew/opt/python@3.11/libexec/bin:$PATH" # because eb wants python, not python3
    ```
