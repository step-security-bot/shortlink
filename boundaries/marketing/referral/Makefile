SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))

# INCLUDE ==============================================================================================================
# Include Makefile
include $(SELF_DIR)/../../../ops/Makefile/common.mk

# REFERRAL TASKS =======================================================================================================
dep: ## Install dependencies
	# Create a virtual environment at .venv
	uv venv

	# Install dependencies
	uv pip install -r pyproject.toml --no-deps

lock: ## Lock dependencies
	-rm requirements.txt
	@uv pip compile pyproject.toml --generate-hashes -o requirements.txt --no-deps

test: ## Run tests
	@pytest --fixtures tests

lint: ## Run linter
	@uvx ruff format
	@uvx ruff check --fix .
