#!/usr/bin/env python
import fastavro
import sys
import glob
import itertools


def combine(input_glob, output_filepath):
    readers = []
    for in_file_path in glob.glob(input_glob):
        infile = open(in_file_path, "rb")
        readers.append(fastavro.reader(infile))
    combined_reader = itertools.chain(*readers)
    with open(output_filepath, "a+") as outfile:
        fastavro.writer(outfile, readers[0].writer_schema,
                        combined_reader, codec="snappy")


if __name__ == "__main__":
    if len(sys.argv) != 3:
        print("usage: combine_avro.py INPUT_GLOB OUTPUT")
        print("    combines all avro files that match INPUT_GLOB into one file at OUTPUT")
        sys.exit(1)
    combine(sys.argv[1], sys.argv[2])
