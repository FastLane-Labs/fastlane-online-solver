TAG = latest
# CTX = <DEPLOY_SERVER>
DOCKER = TAG=$(TAG) docker --context $(CTX)

.DEFAULT_GOAL := help

help:
	@echo "Usage:"
	@echo "  make up CTX=<DEPLOY_SERVER>    - Bring up the services"
	@echo "  make down CTX=<DEPLOY_SERVER>  - Bring down the services"
	@echo "  make deploy CTX=<DEPLOY_SERVER> - Deploy the services"
	@echo "  make clean                     - Remove all Docker images (use with caution)"

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
	$(DOCKER) compose down
	$(DOCKER) compose rm --stop --force
	$(DOCKER) compose pull
	$(DOCKER) compose up -d

clean:
	@if [ -z "$(CTX)" ]; then \
		echo "Error: CTX is not set"; \
		exit 1; \
	fi
	$(DOCKER) system prune -a --volumes -f
