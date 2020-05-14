.PHONY: install
install: datasets
	docker-compose pull
	docker-compose build
	python setup.py install

.PHONY: test-quick
test-quick:
	python -m unittest -v --without-integration

.PHONY: test
test:
	python -m unittest -v

.PHONY: lint
lint:
	flake8

.PHONY: datasets
datasets: data/rubin_single_ccd_sample.avro data/rubin_single_exposure_sample.avro

data:
	mkdir -p data

data/rubin_single_ccd_sample.avro: data
	curl "https://lsst-web.ncsa.illinois.edu/~swnelson/alert-stream-simulator/ap_verify_hits2015/42160601.avro" > data/rubin_single_ccd_sample.avro

data/rubin_single_exposure_sample.avro: data
	curl "https://lsst-web.ncsa.illinois.edu/~swnelson/alert-stream-simulator/sample_10k_alerts.avro" > data/rubin_single_exposure_sample.avro
