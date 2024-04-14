.PHONY: generate generate-server generate-models deploy destroy
deploy:
	docker-compose -f ./build/docker-compose.yml rm && \
    docker-compose -f ./build/docker-compose.yml build --no-cache && \
    docker-compose -f ./build/docker-compose.yml up

destroy:
	 docker-compose -f ./build/docker-compose.yml down


generate: generate-server generate-models

generate-server:
	# Generate Contracts
	oapi-codegen -generate server -package api ./api.yaml > ./server_generated.go

generate-models:
	# Generate Models
	oapi-codegen -generate types -package api ./api.yaml > ./types_generated.go