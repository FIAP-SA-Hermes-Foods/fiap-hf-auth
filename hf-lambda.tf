provider "aws" {
  region = "us-east-1"
}

resource "aws_lambda_function" "hf_lambda" {
  function_name = "hf-api-gateway-func"
  handler       = "bootstrap"
  runtime       = "provided.al2"
  timeout       = 10 // tempo limite em segundos
  memory_size   = 128 // tamanho da memória em MB
  filename      = "bootstrap.zip" // arquivo zipado com o código da função
  source_code_hash = data.archive_file.lambda.output_base64sha256
  role = "{{LAMBDA_EXEC_PERM}}"
}
