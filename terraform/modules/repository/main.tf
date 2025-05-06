resource "aws_ecr_repository" "backend"{
    name = var.repository_name
    image_tag_mutability = "MUTABLE"
    encryption_configuration {
      encryption_type = "AES256"
    }
    force_delete = true
}


resource "aws_ecr_lifecycle_policy" "backend" {
  repository = aws_ecr_repository.backend.name
  policy = jsonencode({
    rules = [
      {
        rulePriority = 1
        description  = "Expire untagged images older than 7 days"
        selection = {
          tagStatus     = "untagged"
          countType     = "sinceImagePushed"
          countUnit     = "days"
          countNumber   = 7
        }
        action = {
          type = "expire"
        }
      }
    ]
  })
}