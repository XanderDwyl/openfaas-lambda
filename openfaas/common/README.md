# Pit Project Examples

- env-secret | Small project to add confidential configuration in secrets instead of adding in environment variables.

  ```
    kubectl apply -f secret.yml
  ```
  NOTE: You need to fill in the value in secret.yml file and make sure to that each value is base64 format. 
