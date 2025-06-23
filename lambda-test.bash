aws lambda invoke \
  --function-name my-go-function \
  --payload '"World"' \
  --cli-binary-format raw-in-base64-out \
  response.json
cat response.json
