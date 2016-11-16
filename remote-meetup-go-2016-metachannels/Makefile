build:
	go build -o concfunc .

run:
	./concfunc

test-basic:
	@echo "/json"
	curl localhost:8080/json?val=abcd
	@echo ""
	@echo "/base64"
	curl localhost:8080/base64?val=abcd
	@echo ""

test-chained:
	@echo "/chained"
	curl localhost:8080/chained?encoder=json\&encoder=base64\&encoder=base64\&encoder=json\&val=abcd
	@echo ""

test-switch:
	curl localhost:8080/switch

test-scale:
	curl localhost:8080/scale?num=500
