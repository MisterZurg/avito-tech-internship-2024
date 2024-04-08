.PHONY: generate generate-server generate-models

generate: generate-server generate-models

generate-server:
	# Generate Contracts
	oapi-codegen -generate server -package api ./api.yaml > ./server_generated.go

generate-models:
	# Generate Models
	oapi-codegen -generate types -package api ./api.yaml > ./types_generated.go