#!/bin/bash
set +xeo pipefail

# Little script for constructing sample datasets on lsst-dev servers. Should be
# run with the alert-stream-simulator's 'scripts' directory in $PATH, and with
# fastavro and python-snappy Python packages installed.

INPUT_REPO=/project/morriscb/src/ap_verify_hits2015/DM-25930/
DATASET=DECam-HiTS
DATE=2020-07-15
PUBLIC_DIR=/project/ebellm/sample_precursor_alerts/$DATE/$DATASET

mkdir -p $PUBLIC_DIR


cp $INPUT_REPO/alerts/42160601.avro ccdvisit_42160601.avro
snappy_recompress.py ccdvisit_42160601.avro ccdvisit_42160601_snappy.avro
rm ccdvisit_42160601.avro

echo "Combining one visit..."
combine_avro.py "${INPUT_REPO}/alerts/421590*.avro" visit_421590.avro

echo "Combining all visits..."
combine_avro.py "${INPUT_REPO}/alerts/*.avro" all_visits.avro

echo "Moving files..."
cp -f $INPUT_REPO/association.db $PUBLIC_DIR/association_${DATASET}_${DATE}.db
mv -f ccdvisit_42160601_snappy.avro $PUBLIC_DIR/single_ccd_sample_${DATASET}_${DATE}.avro
mv -f visit_421590.avro $PUBLIC_DIR/single_visit_sample_${DATASET}_${DATE}.avro
mv -f all_visits.avro $PUBLIC_DIR/all_visits_${DATASET}_${DATE}.avro
