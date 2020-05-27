#!/usr/bin/env python
import fastavro
import sys


def recompress(input_filename, output_filename):
    with open(input_filename, "rb") as infile:
        reader = fastavro.reader(infile)

        with open(output_filename, "wb") as outfile:
            fastavro.writer(outfile, reader.writer_schema, reader,
                            codec="snappy")


if __name__ == "__main__":
    if len(sys.argv) != 3:
        print("usage: snappy_recompress.py INPUT OUTPUT")
        print("    reserializes INPUT to OUTPUT, compressing the avro stream with snappy")
        sys.exit(1)

    recompress(sys.argv[1], sys.argv[2])
