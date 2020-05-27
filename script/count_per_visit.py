#!/usr/bin/env python
import fastavro
import sys
import collections


def count(input_filename):
    counts = collections.defaultdict(int)
    with open(input_filename, "rb") as infile:
        reader = fastavro.reader(infile)
        for m in reader:
            counts[m['diaSource']['midPointTai']] += 1
    for visit_id in sorted(counts, key=lambda id: counts[id]):
        print(visit_id, counts[visit_id])


if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("usage: count_per_mid_point_tai.py INPUT")
        print("    counts the number of alerts per midPointTai in INPUT")
        sys.exit(1)
    count(sys.argv[1])
