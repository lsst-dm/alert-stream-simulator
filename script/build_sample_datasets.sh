#!/bin/bash
set +xeo pipefail

# Little script for constructing sample datasets on lsst-dev servers. Should be
# run with the alert-stream-simulator's 'scripts' directory in $PATH, and with
# fastavro and python-snappy Python packages installed.

combine_avro.py "/project/morriscb/src/ap_verify_hits2015/alerts/421590*.avro" visit_421590.avro

cp /project/morriscb/src/ap_verify_hits2015/alerts/42160601.avro ccdvisit_42160601.avro
snappy_recompress.py ccdvisit_42160601.avro ccdvisit_42160601_snappy.avro
rm ccdvisit_42160601.avro

PUBLIC_DIR=/home/swnelson/public_html/alert-stream-simulator

mv visit_421590.avro $PUBLIC_DIR/rubin_single_visit_sample.avro
mv ccdvisit_42160601_snappy.avro $PUBLIC_DIR/rubin_single_ccd_sample.avro
