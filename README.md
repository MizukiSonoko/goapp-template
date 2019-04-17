
# App template

this repo contains functions as follows:
- Http server
- Deploy ECR 
- Deploy ECS

# How to use 

- Set `AWS_ECR_ACCOUNT_URL` in circleci environments settings
This value will be like `1234567890123.dkr.ecr.ap-northeast-1.amazonaws.com`
also this value is same.
    - `AWS_REGION`
    - `AWS_ACCESS_KEY_ID`
    - `AWS_SECRET_ACCESS_KEY`

- Replace this value in `.circleci/config.yml`
`<repo name>`
`<task name>` 
`<service name>`
`<cluster name>`
`<container name>`
`<repo name>`


        
