DOCKER = docker --context $(CTX)

.DEFAULT_GOAL := help

help:
	@echo "Usage:"
	@echo "  make up CTX=<DEPLOY_SERVER>    - Bring up the services"
	@echo "  make down CTX=<DEPLOY_SERVER>  - Bring down the services"
	@echo "  make deploy CTX=<DEPLOY_SERVER> - Deploy the services"
	@echo "  make clean CTX=<DEPLOY_SERVER>  - Remove the containers and volumes"

down:
	@if [ -z "$(CTX)" ]; then \
		echo "Error: CTX is not set"; \
		exit 1; \
	fi
	$(DOCKER) compose down

up:
	@if [ -z "$(CTX)" ]; then \
		echo "Error: CTX is not set"; \
		exit 1; \
	fi
	$(DOCKER) compose up -d

deploy:
	@if [ -z "$(CTX)" ]; then \
		echo "Error: CTX is not set"; \
		exit 1; \
	fi
	$(DOCKER) compose down --remove-orphans
	$(DOCKER) compose rm --stop --force
	$(DOCKER) compose pull
	$(DOCKER) compose run --rm -e MODE=data_gen fastlane-online-solver
	$(DOCKER) compose up -d fastlane-online-solver

clean:
	@if [ -z "$(CTX)" ]; then \
		echo "Error: CTX is not set"; \
		exit 1; \
	fi
	$(DOCKER) compose down --volumes --remove-orphans
	$(DOCKER) compose rm --stop --force
