aws lambda create-function --function-name maio-21-2024-web --zip-file fileb://golambda-handler.zip --handler main --runtime provided.al2023 --role arn:aws:iam::937852338641:role/service-role/go-hello-maio-21-role-km44lsfv


1) 
    GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o bootstrap main.go

2) zip

zip golambda-handler.zip bootstrap


4) 

aws lambda create-function \
    --function-name 22-maio-2024-web-com-fiber \
    --zip-file fileb://golambda-handler.zip \
    --handler main \
    --runtime provided.al2023 \
    --role arn:aws:iam::937852338641:role/service-role/go-hello-maio-21-role-km44lsfv \
    --architectures arm64


Update:

aws lambda update-function-code \
    --function-name maio-21-2024-web \
    --zip-file fileb://golambda-handler.zip
