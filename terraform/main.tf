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
    API_TOKEN = module.api_token.var_value
    API_URL = module.api_url.var_value
    PORT = "8000"
    DB_USER = module.db_user.var_value
    DB_HOST = module.db_host.var_value
    DB_PORT = module.db_port.var_value
    DB_PASSWORD = module.db_password.var_value
    DB_DBNAME = module.db_dbname.var_value
    DB_SSL_MODE = module.db_ssl_mode.var_value
    EXTERNAL_API_URL = module.external_api_url.var_value
    EXTERNAL_API_KEY = module.external_api_key.var_value    
}