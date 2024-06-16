# Colors
RESET	= \033[0m
RED 	= \033[1;31m
GREEN 	= \033[1;32m

# Variables
TERRAFORM_DIR = deploy
WEBSERVER_DIR = go_webserver

# Commands
all: vars init apply

vars:
	@echo "$(GREEN)Generating terraform.tfvars file...$(RESET)"
	@cd $(TERRAFORM_DIR) && ../scripts/create_tfvars.sh

init:
	@echo "$(GREEN)Running terraform init...$(RESET)"
	@cd $(TERRAFORM_DIR) && terraform init

plan:
	@echo "$(GREEN)Running terraform plan...$(RESET)"
	@cd $(TERRAFORM_DIR) && terraform plan

apply:
	@echo "$(GREEN)Deploying go-webserver application to Digital Ocean App Platform...$(RESET)"
	@cd $(TERRAFORM_DIR) && terraform apply -auto-approve

destroy:
	@echo "$(GREEN)Deploying go-webserver application to Digital Ocean App Platform...$(RESET)"
	@cd $(TERRAFORM_DIR) && terraform destroy -auto-approve

update_image:
	@echo "$(GREEN)Updating Docker Hub repository with the most recent changes...$(RESET)"
	@cd $(WEBSERVER_DIR) && ../scripts/update_image.sh

test:
	@echo "$(GREEN)Testing webserver...$(RESET)"
	@./scripts/test.sh

clean:
	@echo "$(RED)Removing terraform.tfvars file...$(RESET)"
	@rm -f $(TERRAFORM_DIR)/terraform.tfvars
	
fclean: clean
	@echo "$(RED)Cleaning terraform files...$(RESET)"
	@rm -rf $(TERRAFORM_DIR)/.terraform
	@rm -f $(TERRAFORM_DIR)/.terraform.lock.hcl
	@rm -f $(TERRAFORM_DIR)/terraform.tfstate
	@rm -f $(TERRAFORM_DIR)/terraform.tfstate.backup
	@rm -f $(TERRAFORM_DIR)/*.tfout
	@rm -f $(TERRAFORM_DIR)/*.plan

re: fclean all