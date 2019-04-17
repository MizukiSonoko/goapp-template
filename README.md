
# App template

this repo contains functions as follows:
- Http server
- Deploy ECR 
- Deploy ECS

# How to use 

### Set `AWS_ECR_ACCOUNT_URL` in circleci environments settings
This value will be like `1234567890123.dkr.ecr.ap-northeast-1.amazonaws.com`.  
also this value is same.  
- `AWS_REGION`  
- `AWS_ACCESS_KEY_ID`  
- `AWS_SECRET_ACCESS_KEY`  

### Replace this value in `.circleci/config.yml`  
#### `<repo name>`
This value is `goapp-template` in this picture.    
<img width="300" alt="Screen Shot" src="https://user-images.githubusercontent.com/6281583/56255279-4dd14c80-60ff-11e9-990d-7392c819afeb.png">


#### `<task name>` 
This value is `goapp-template` in this picture.    
<img width="300" alt="Screen Shot" src="https://user-images.githubusercontent.com/6281583/56255422-c59f7700-60ff-11e9-8007-fcb9219c8fff.png">  

  
 
#### `<service name>`
This value is `goapp-service` in this picture.   
<img width="300" alt="Screen Shot" src="https://user-images.githubusercontent.com/6281583/56255723-de5c5c80-6100-11e9-8ae5-3c47162d973f.png">  

#### `<cluster name>`
This value is `goapp-cluster` in this picture.  
<img width="300" alt="Screen Shot" src="https://user-images.githubusercontent.com/6281583/56255607-7e65b600-6100-11e9-8b18-4017c645906d.png">

#### `<container name>`  
This value is `goapp-template` in this picture.  
<img width="300" alt="Screen Shot" src="https://user-images.githubusercontent.com/6281583/56255783-18c5f980-6101-11e9-961f-b373707fe104.png">

        
