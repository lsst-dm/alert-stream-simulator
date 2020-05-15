#!/usr/bin/env python
import fastavro
import sys


def run_filter(input_filename, output_filename, filter_tai):
    with open(input_filename, "rb") as input_fp:
        reader = fastavro.reader(input_fp)
        filtered_reader = filter(
            lambda alert: alert['diaSource']['midPointTai'] == filter_tai,
            reader
        )
        with open(output_filename, "ab") as output_fp:
            fastavro.writer(
                fo=output_fp,
                schema=reader.writer_schema,
                records=filtered_reader,
                codec=reader.codec,
            )


if __name__ == "__main__":
    if len(sys.argv) != 4:
        print("usage: filter_by_midpoint_tai.py INPUT OUTPUT TAI")
        print("    copies all alerts with given TAI in INPUT to OUTPUT")
        sys.exit(1)
    run_filter(sys.argv[1], sys.argv[2], float(sys.argv[3]))
