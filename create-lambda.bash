aws lambda create-function \
  --function-name my-go-function \
  --runtime provided.al2 \
  --handler bootstrap \
  --zip-file fileb://function.zip \
  --role xxx \
  --region ap-northeast-1
