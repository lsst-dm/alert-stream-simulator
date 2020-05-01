import multiprocessing
import io
import functools
import time
import sys
from fastavro import reader, schemaless_writer


def do_it_the_dumb_way(filename):
    # Read the file, all the way
    avro_reader = reader(open(filename, "rb"))
    records = [r for r in avro_reader]
    start = time.perf_counter()
    result = serialize(records, avro_reader.schema)

    for serialized_alert in result:
        pass
    end = time.perf_counter()
    duration = (end - start)
    print(f"processed 1000 in {duration} ({1000.0 / duration}/s)")


def serialize(records, schema):
    with multiprocessing.Pool(int(sys.argv[1])) as pool:
        return pool.map(
            functools.partial(serialize_one, schema=schema),
            records)


def serialize_one(alert, schema):
    buf = io.BytesIO()
    schemaless_writer(buf, schema, alert)
    return buf.getvalue()


if __name__ == "__main__":
    do_it_the_dumb_way("data/hits_sample.avro")
