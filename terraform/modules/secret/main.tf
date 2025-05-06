resource "aws_secretsmanager_secret" "secret_env" {
    name = var.var_name
}

resource "aws_secretsmanager_secret_version" "secret_env"{
    secret_id = aws_secretsmanager_secret.secret_env.id
    secret_string = var.var_value
}