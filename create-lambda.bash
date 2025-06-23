aws lambda create-function \
  --function-name my-go-function \
  --runtime go1.x \
  --handler main \
  --zip-file fileb://function.zip \
  --role xxx \
  --region ap-northeast-1
