#!/usr/bin/env python
import fastavro
import sys
import glob
import itertools
import os


def combine(input_glob, output_filepath):
    readers = []

    if os.path.exists(output_filepath):
        os.remove(output_filepath)
    # names are ccdVisitId; sort to put them in temporal order
    input_files = sorted(glob.glob(input_glob))
    print(f'Combining {len(input_files)} files')
    for in_file_path in input_files:
        infile = open(in_file_path, "rb")
        readers.append(fastavro.reader(infile))
    combined_reader = itertools.chain(*readers)
    with open(output_filepath, "a+b") as outfile:
        fastavro.writer(outfile, readers[0].writer_schema,
                        combined_reader, codec="snappy")


if __name__ == "__main__":
    if len(sys.argv) != 3:
        print("usage: combine_avro.py INPUT_GLOB OUTPUT")
        print("    combines all avro files that match INPUT_GLOB into one file at OUTPUT")
        sys.exit(1)
    combine(sys.argv[1], sys.argv[2])
