import argparse
import logging

from streamsim import sender


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
        logging.debug(f"dispatching publish-file command with args: {args}")
        publish_file(args.broker, args.topic, args.file)


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
        "file", type=argparse.FileType('rb'),
        help="alert file to send",
    )
    return parser


def publish_file(broker, topic, alert_file):
    """Send all the alerts in a single file to Kafka."""
    producer = sender.AlertProducer(broker, topic)
    producer.send_alert(alert_file.read())
    producer.flush(timeout=10)
