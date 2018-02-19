CURRENT_DIR=$(shell pwd)

swag:
	ovhapi2openapi -i $(res)/$(res).yaml -o $(res)/swagger.json

model: swag
	swagger generate model -m $(res)swag -t ./$(res)/ -f $(res)/swagger.json
	rm $(res)/$(res)swag/*default_body.go

doc: swag	
	swagger serve $(res)/swagger.json

buildall:
	for res in auth me; do\
		$(MAKE) res=$$res model; \
	done