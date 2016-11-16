build:
	go build -o concfunc .

run:
	./concfunc

test:
	@echo "json:"
	curl localhost:8080/json?val=abcd
	@echo ""
	@echo "base64:"
	curl localhost:8080/base64?val=abcd
	@echo ""
	@echo "stopping"
	curl localhost:8080/stop
	@echo ""
	@echo "json:"
	curl localhost:8080/json?val=abcd
	@echo ""
	@echo "base64:"
	curl localhost:8080/base64?val=abcd
	@echo ""
