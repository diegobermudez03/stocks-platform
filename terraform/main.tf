module "backend_repo"{
    source = "./modules/repository"
    repository_name = "stocks-backend"
}


module "iam_policy"{
    source = "./modules/iam_policy"
    role_name = "apprunner-access-role"
    policy_name = "apprunner-access-policy"
}


module "app_runner_backend"{
    source = "./modules/app_runner"
    app_runner_name = "stocks-backend-runner"
    iam_role_arn = module.iam_policy.apprunner_role_arn
    repository_url = module.backend_repo.ecr_repository_url
    API_TOKEN = module.api_token.arn
    API_URL = module.api_url.arn
    PORT = "8000"
    DB_USER = module.db_user.arn
    DB_HOST = module.db_host.arn
    DB_PORT = module.db_port.arn
    DB_PASSWORD = module.db_password.arn
    DB_DBNAME = module.db_dbname.arn
    DB_SSL_MODE = module.db_ssl_mode.arn
    EXTERNAL_API_URL = module.external_api_url.arn
    EXTERNAL_API_KEY = module.external_api_key.arn
    WEBSOCKET_API_URL = module.websocket_api_url.arn
    ALL_OR_NOTHING = module.all_or_nothing.arn
}

module "frontend_bucket"{
    source = "./modules/frontend_bucket"
    bucket_name = "stocks-frontend-bucket"
}