variable "region" {
  type        = string
  description = "AWS default region"
  default     = "eu-south-1"
}

variable "webapp" {
  default = {
    repository = {
      organization    = "pagopa"
      name            = "ecs-app-test"
      branch_name     = "refs/heads/main"
      pipelines_path  = "src/pipelines"
      yml_prefix_name = null
    }
    pipeline = {
      enable_code_review = true
      enable_deploy      = true
    }
  }
}