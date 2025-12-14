variable "aws_region" {
  description = "AWS region"
  default     = "ap-south-1"
}

variable "project_name" {
  description = "Project name prefix"
  default     = "ecs-demo"
}

variable "container_image" {
  description = "Docker image"
  default     = "prasad0508/simpletimeservice:latest"
}

variable "container_port" {
  description = "Container port"
  default     = 8080
}
