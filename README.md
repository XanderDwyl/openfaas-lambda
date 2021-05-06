# openfaas-go-examples

This is a set of examples that i've used when migrating aws lambda to openfaas platform.

## Prerequisite

Make sure that you successfully installed OpenFaas. 

https://docs.openfaas.com/cli/install/

## How do I set this up
- Update the stack.yml and replace the my docker account `${docker_name}` to your user docker account in each of the function.

  stack.yml
  ```
  ...
    lang: go
    handler: ./aws/openfaas/redis
    image: ${docker_name}/redis:latest
  ...
  ```
- Fill up the environment variable in `stack.yml`

  stack.yml
  ```
  ...
    environment:
      write_debug: true
      DB_URL: ''
      AWS_ID: ''
      AWS_KEY: ''
      REGION: ''
      SENTINEL_1: ''
      SENTINEL_2: ''
      SENTINEL_3: ''
      SENTINEL_PASS: ''
  ...
  ```
- Then deploy the functions with `faas-cli up -f stack.yml`
