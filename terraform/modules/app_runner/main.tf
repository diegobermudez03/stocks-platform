resource "aws_apprunner_service" "backend" {
  service_name = var.app_runner_name

  source_configuration {
    authentication_configuration {
      access_role_arn = var.iam_role_arn
    }

    image_repository {
      image_identifier      = "${var.repository_url}:latest"
      image_repository_type = "ECR"

      image_configuration {
        port = var.PORT

        runtime_environment_secrets = {
          DB_HOST     = var.DB_HOST
          DB_PORT     = var.DB_PORT
          API_TOKEN = var.API_TOKEN
          API_URL = var.API_URL
          PORT = var.PORT
          DB_USER = var.DB_USER
          DB_PASSWORD = var.DB_PASSWORD
          DB_DBNAME = var.DB_DBNAME
          DB_SSL_MODE = var.DB_SSL_MODE
          EXTERNAL_API_URL = var.EXTERNAL_API_URL
          EXTERNAL_API_KEY = var.EXTERNAL_API_KEY
        }
      }
    }

    auto_deployments_enabled = true
  }

  instance_configuration {
    cpu    = "1024"
    memory = "2048"
    instance_role_arn = aws_iam_role.app_runner_role.arn
  }
}


resource "aws_iam_role" "app_runner_role" {
  name = "app-runner-secrets-access-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [{
      Action = "sts:AssumeRole",
      Principal = {
        Service = "build.apprunner.amazonaws.com"
      },
      Effect = "Allow"
    }]
  })
}

resource "aws_iam_policy" "secrets_policy" {
  name = "AppRunnerSecretsPolicy"

  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [{
      Action = [
        "secretsmanager:GetSecretValue",
        "secretsmanager:DescribeSecret"
      ],
      Effect = "Allow",
      Resource = "*"
    }]
  })
}

resource "aws_iam_role_policy_attachment" "attach_secrets_policy" {
  role       = aws_iam_role.app_runner_role.name
  policy_arn = aws_iam_policy.secrets_policy.arn
}
