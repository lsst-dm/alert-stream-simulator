.PHONY: test-quick
test-quick:
	python -m unittest -v --without-integration

.PHONY: test
test:
	python -m unittest -v

.PHONY: lint
lint:
	flake8
