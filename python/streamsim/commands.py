import asyncio
import argparse
import logging
import datetime
import concurrent.futures

from streamsim import sender, pacer, fastavro_pacer

logger = logging.getLogger("rubin-alert-sim")

def run():
    """Parse command line arguments and execute a command."""
    parser = construct_argparser()
    args = parser.parse_args()

    if args.debug:
        logging.basicConfig(level=logging.DEBUG)
    elif args.verbose:
        logging.basicConfig(level=logging.INFO)
    else:
        logging.basicConfig(level=logging.WARNING)

    if args.subcommand is None:
        parser.print_usage()
    if args.subcommand == "publish-file":
        logger.debug(f"dispatching publish-file command with args: {args}")
        asyncio.run(publish_file(args.broker, args.topic, args.file, args.timeout_sec, fastavro=args.use_fastavro, compression=args.compression, parallelism=args.parallelism))


def construct_argparser():
    """Set up an argument parser for the rubin-alert-sim CLI.

    Returns
    -------
    parser : `argparse.ArgumentParser`
        A parser for command line arguments to dispatch commands.
    """
    parser = argparse.ArgumentParser(
        prog="rubin-alert-sim",
        formatter_class=argparse.ArgumentDefaultsHelpFormatter,
    )
    parser.add_argument("-v", "--verbose", action="store_true", help="enable info-level logging")
    parser.add_argument("-d", "--debug", action="store_true", help="enable debug-level logging")
    subparsers = parser.add_subparsers(title="subcommands", dest="subcommand")

    publish_file_cmd = subparsers.add_parser(
        "publish-file", help="publish alert data sourced from a single file",
        formatter_class=argparse.ArgumentDefaultsHelpFormatter,
    )
    publish_file_cmd.add_argument(
        "-B", "--broker", type=str, default="localhost:9092",
        help="address of the Kafka broker to connect to",
    )
    publish_file_cmd.add_argument(
        "--topic", type=str, default="alerts",
        help="Kafka topic name to use",
    )
    publish_file_cmd.add_argument(
        "--timeout-sec", type=float, default=10.0,
        help="how long, in seconds, to wait before giving up when flushing a write to kafka",
    )
    publish_file_cmd.add_argument(
        "--use-fastavro", action="store_true",
        help="use fastavro library for serde",
    )
    publish_file_cmd.add_argument(
        "--compression", type=str, default="null",
        help="compression codec to use on writes (only works with fastavro)",
    )
    publish_file_cmd.add_argument(
        "--parallelism", type=str, default="None",
        help="parallelism mode (can be 'none', 'thread', or 'process')",
    )
    publish_file_cmd.add_argument(
        "file", type=argparse.FileType('rb'),
        help="alert file to send",
    )
    return parser


async def publish_file(broker, topic, alert_file, timeout, fastavro, compression, parallelism):
    """Send all the alerts in a single file to Kafka."""
    if parallelism == "thread":
        logger.debug("thread parallelism")
        pool = concurrent.futures.ThreadPoolExecutor()
    elif parallelism == "process":
        logger.debug("process parallelism")
        pool = concurrent.futures.ProcessPoolExecutor()
    else:
        logger.debug("no parallelism")
        pool = None
    if fastavro:
        pace = fastavro_pacer.FastavroSimplePacer(alert_file)
        producer = sender.FastavroAlertProducer(broker, topic, pace.schema, compression, pool)
    else:
        pace = pacer.SimplePacer(alert_file)
        producer = sender.AlertProducer(broker, topic, pace.schema)
    start = datetime.datetime.now()
    n = 0
    async for alert in pace.iterate():
        logger.info(f"sending alert (alert={alert['alertId']} source={alert['diaSource']['diaSourceId']} ccdVisit={alert['diaSource']['ccdVisitId']})")
        await producer.send_alert(alert)
        n += 1
    end = datetime.datetime.now()
    duration = end - start
    rate = n / duration.total_seconds()
    print(f"sent {n} alerts in {duration.total_seconds():.2f} sec ({rate:.2f}/s)")
    producer.flush(timeout=timeout)
