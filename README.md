# Auth service

## Run

make run.local

## Request

#### Local

##### grpcurl

grpcurl --plaintext -d '{"userId": "12345"}' localhost:8080 boutit.auth.api.AuthService/CreateToken

grpcurl --plaintext -d '{"token": ADD_TOKEN_HERE}' localhost:8080 boutit.auth.api.AuthService/ValidateToken

##### curl

NOTE: curl must be done through the api-gateway

curl -X POST -k http://localhost:8090/v1/auth/create_token -d '{"userId": "ey295-asdgjsg-asdgljkas-33dll", "roles": []}'

## Resources

### Generate Public/Private Key pair

1. Navigate to this website, select 512 bit and click "Generate New Keys": https://travistidwell.com/jsencrypt/demo/
2. Copy private key, navigate to this website, paste into textarea and click encode: https://www.base64encode.org/
3. Copy base64 encoded result to config file
4. Navigate back to step 2 and repeat steps 2 & 3 for hte public key
