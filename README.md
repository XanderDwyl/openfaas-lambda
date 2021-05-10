# openfaas-lambda

These are sets of example that i've used when migrating aws lambda to openfaas platform. A part of documenting on how it works, I'm open for your contributions for the improvement and future reference in working projects related to openfaas.

## Pre-requisite

Make sure that you successfully installed OpenFaas. 

https://docs.openfaas.com/cli/install/

## How do I set this up
- Update the stack.yml and replace the my docker account `${docker_name}` to your docker account in each of the function.

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

  Note: The OpenFaaS philosophy is that environment variables should be used for non-confidential configuration values only, and not used to inject secrets. I just add it here for quick and easy configuration during my testing but will get rid of this in production for sure.


- Then deploy the functions with `faas-cli up -f stack.yml`

## TODO
- CloudStorage support

## References

[OpenFaas](https://www.openfaas.com/) => Simple, powerful functions from anywhere.

[AWS SDK](https://github.com/aws/aws-sdk-go) => The SDK makes it easy to integrate your Go application with the full suite of AWS services including Amazon S3, Amazon DynamoDB, Amazon SQS, and more.
