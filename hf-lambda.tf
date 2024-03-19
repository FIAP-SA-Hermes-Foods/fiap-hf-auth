provider "aws" {
  region = "us-east-1"
}

resource "aws_lambda_function" "hf_lambda" {
  function_name = "hf-api-gateway-func"
  handler       = "hf-auth"
  runtime       = "provided.al2"
  timeout       = 10 // tempo limite em segundos
  memory_size   = 128 // tamanho da memória em MB
  filename      = "hf_api_gateway_lambda.zip" // arquivo zipado com o código da função
  source_code_hash = filebase64sha256("hf_api_gateway_lambda.zip") // hash do arquivo zipado
  // Permissões para invocar a função
  role = "{{LAMBDA_EXEC_PERM}}"
}

// Definição da role (permissões) para a função Lambda
resource "aws_iam_role" "lambda_exec_role" {
  name = "lambda-exec-role"
  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Effect    = "Allow",
        Principal = {
          Service = "lambda.amazonaws.com"
        },
        Resource = "*",
        Action = [
            "lambda:InvokeAsync",
            "lambda:InvokeFunction"
        ]
      }
    ]
  })
}

 resource "aws_iam_policy" "lambda_logs_policy" {
   name = "lambda-logs-gateway-policy"
   description = "Policy for writing Lambda logs to CloudWatch"
   policy = jsonencode({
     Version = "2012-10-17",
     Statement = [
       {
         Effect   = "Allow",
         Action   = [
           "logs:CreateLogGroup",
           "logs:CreateLogStream",
           "logs:PutLogEvents"
         ],
         Resource = "*"
       }
     ]
   })
 } 
