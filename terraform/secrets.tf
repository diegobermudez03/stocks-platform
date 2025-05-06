module "api_token"{
    source = "./modules/secret"
    var_name = "API_TOKEN"
    var_value = var.api_token
}

module "api_url"{
    source = "./modules/secret"
    var_name = "API_URL"
    var_value = var.api_url
}

module "db_host"{
    source = "./modules/secret"
    var_name = "DB_HOST"
    var_value = var.db_host
}

module "db_port"{
    source = "./modules/secret"
    var_name = "DB_PORT"
    var_value = var.db_port
}

module "db_user"{
    source = "./modules/secret"
    var_name = "DB_USER"
    var_value = var.db_user
}

module "db_password"{
    source = "./modules/secret"
    var_name = "DB_PASSWORD"
    var_value = var.db_password
}

module "db_dbname"{
    source = "./modules/secret"
    var_name = "DB_DBNAME"
    var_value = var.db_dbname
}

module "db_ssl_mode"{
    source = "./modules/secret"
    var_name = "DB_SSL_MODE"
    var_value = var.db_ssl_mode
}

module "external_api_url"{
    source = "./modules/secret"
    var_name = "EXTERNAL_API_URL"
    var_value = var.external_api_url
}

module "external_api_key"{
    source = "./modules/secret"
    var_name = "EXTERNAL_API_KEY"
    var_value = var.external_api_key
}