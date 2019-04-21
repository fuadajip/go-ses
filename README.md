# GO SES
Mywishes AWS SES microservice with go lang

## How To
1. Install go 1.11 and setup the go mod
2. Create app.config.dev.yml or app.config.prod.yml and fill the value
3. Run this command to get vendor
    ```bash
    make vendor
    ```
4. Run time app
    ```bash
    make start
    ```
5. Build image
    ```bash
    make docker
    ```
6. Start from docker image
    ```bash
    make run
    ```