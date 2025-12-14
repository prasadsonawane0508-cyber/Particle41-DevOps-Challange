#  ECS Fargate Infrastructure using Terraform (AWS)

## Overview

This project provisions a production-style container infrastructure on AWS using Terraform.  
A Dockerized application runs on Amazon ECS (Fargate) inside private subnets and is exposed to the internet via an Application Load Balancer (ALB) deployed in public subnets.

The setup follows AWS and DevOps best practices:
- Containers run without public IPs
- Network isolation using public and private subnets
- Infrastructure managed using Terraform (IaC)

---

## Architecture

Internet  
→ Application Load Balancer (Public Subnets)  
→ ECS Fargate Service  
→ Private Subnets (No Public IPs)  
→ NAT Gateway (Outbound Internet)

---

## Tech Stack

- AWS (VPC, ECS Fargate, ALB, NAT Gateway)
- Terraform
- Docker
- GitHub

---

## Repository Structure

terraform/
├── main.tf  
├── variables.tf  
├── outputs.tf  
├── versions.tf  
├── .gitignore  
└── README.md  

---

## Infrastructure Details

### VPC
- CIDR block: 10.0.0.0/16
- 2 Public subnets
- 2 Private subnets
- Internet Gateway for public subnets
- NAT Gateway for private subnet outbound traffic

### ECS (Fargate)
- ECS Cluster
- ECS Task Definition
- ECS Service running in private subnets
- No public IP assigned to containers

### Load Balancer
- Application Load Balancer in public subnets
- Listener on port 80
- Target Group forwarding traffic to port 8080

---

## Application

Docker Image:
prasad0508/simpletimeservice:latest

Container Port:
8080

---

## How to Deploy

### Prerequisites
- AWS CLI configured
- Terraform >= 1.3 installed

### Deployment Steps

terraform init  
terraform apply  

Type `yes` when prompted.

---

## Access the Application

After deployment, Terraform outputs the ALB DNS name.

Open in browser:

http://<alb_dns_name>

Example:
http://ecs-demo-alb-xxxxxxxx.us-east-1.elb.amazonaws.com

---

## Cleanup

To avoid AWS charges, destroy the infrastructure:

terraform destroy

---

## Security & Best Practices

- ECS tasks run in private subnets
- No public IPs assigned to containers
- ALB handles all inbound traffic
- Terraform state files are excluded using .gitignore
- Dynamic availability zones used to avoid region mismatch

---

## Author

Prasad sonawane  
DevOps / Cloud Engineer
