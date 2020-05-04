# This file is part of alert-stream-simulator.
#
# Developed for the LSST Data Management System.
# This product includes software developed by the LSST Project
# (https://www.lsst.org).
# See the COPYRIGHT file at the top-level directory of this distribution
# for details of code ownership.
#
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with this program.  If not, see <https://www.gnu.org/licenses/>.

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
        publish_file(args.broker, args.topic, args.file, args.timeout_sec)


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
        "file", type=argparse.FileType('rb'),
        help="alert file to send",
    )
    return parser


def publish_file(broker, topic, alert_file, timeout):
    """Send all the alerts in a single file to Kafka."""
    producer = sender.AlertProducer(broker, topic)
    producer.send_alert(alert_file.read())
    producer.flush(timeout=timeout)
