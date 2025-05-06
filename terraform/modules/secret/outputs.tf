output "arn"{
    value = aws_secretsmanager_secret.secret_env.arn
}