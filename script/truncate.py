#!/usr/bin/env python
import fastavro
import sys
import itertools


def truncate(input_filename, output_filename, n):
    with open(input_filename, "rb") as infile:
        reader = fastavro.reader(infile)

        partial_reader = itertools.islice(reader, n)
        with open(output_filename, "wb") as outfile:
            fastavro.writer(outfile, reader.writer_schema,
                            partial_reader, codec=reader.codec)


if __name__ == "__main__":
    if len(sys.argv) != 4:
        print("usage: truncate.py INPUT OUTPUT N")
        print("    copy the first N alerts from INPUT to OUTPUT")
        sys.exit(1)

    truncate(sys.argv[1], sys.argv[2], int(sys.argv[3]))
