.PHONY: install
install: datasets
	docker-compose pull
	docker-compose build
	python setup.py install

.PHONY: test-quick
test-quick:
	python -m pytest -v --without-integration

.PHONY: test
test:
	python -m unittest -v

.PHONY: lint
lint:
	flake8

.PHONY: image
image: datasets
	docker build .

.PHONY: datasets
datasets: data/rubin_single_ccd_sample.avro data/rubin_single_visit_sample.avro

data:
	mkdir -p data

data/rubin_single_ccd_sample.avro: data
	wget --output-document data/rubin_single_ccd_sample.avro https://lsst.ncsa.illinois.edu/~ebellm/sample_precursor_alerts/latest_single_ccd_sample.avro

data/rubin_single_visit_sample.avro: data
	wget --output-document data/rubin_single_visit_sample.avro https://lsst.ncsa.illinois.edu/~ebellm/sample_precursor_alerts/latest_single_visit_sample.avro
