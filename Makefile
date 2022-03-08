# .PHONY: docker-run
# docker-run:
# 	docker run -d -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=root mysql:5.7

# docker-down:
# 	docker rm -f mysql

.PHONY: docker-up
docker-up:
	docker-compose -f docker-compose.yaml up --build

.PHONY: docker-down
docker-down: ## Stop docker containers and clear artefacts.
	docker-compose -f docker-compose.yaml down
	docker system prune


# .PHONY: open
# open:
# 	docker exec -it mysql mysql -uroot -proot

# .PHONY: create
# create:
# 	docker exec -it mysql mysql -uroot -proot -e 'CREATE DATABASE todolist'