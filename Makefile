.PHONY: test
test:
	python -m unittest

.PHONY: lint
lint:
	flake8
