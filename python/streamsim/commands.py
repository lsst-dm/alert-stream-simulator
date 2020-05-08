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

from streamsim import creator, player


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

    if args.subcommand == "create-stream":
        logging.debug(f"dispatching create-stream command with args: {args}")
        n = creator.create(args.broker, args.dst_topic, args.file, args.timeout_sec, args.force)
        print(f"successfully preloaded stream with {n} alerts")
    elif args.subcommand == "play-stream":
        logging.debug(f"dispatching play-stream command with args: {args}")
        n = player.play(args.broker, args.src_topic, args.dst_topic, args.dst_topic_partitions,
                        args.force, args.repeat_interval)
        print(f"played {n} alerts from the stream")
    else:
        parser.print_usage()


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

    create_cmd = subparsers.add_parser(
        "create-stream", help="create a stream dataset to be run through the simulation.",
        formatter_class=argparse.ArgumentDefaultsHelpFormatter,
    )
    create_cmd.add_argument(
        "-B", "--broker", type=str, default="localhost:9092",
        help="address of the Kafka broker to connect to",
    )
    create_cmd.add_argument(
        "--dst-topic", type=str, default="alerts-reservoir",
        help="Kafka topic name to use for the dataset reservoir",
    )
    create_cmd.add_argument(
        "--force", action="store_true", help="overwrite dst-topic if it already exists",
    )
    create_cmd.add_argument(
        "--timeout-sec", type=float, default=10.0,
        help="how long, in seconds, to wait before giving up when flushing a write to kafka",
    )
    create_cmd.add_argument(
        "file", type=argparse.FileType('rb'),
        help="alert file to send",
    )

    play_cmd = subparsers.add_parser(
        "play-stream", help="play back a stream that has already been created",
        formatter_class=argparse.ArgumentDefaultsHelpFormatter,
    )
    play_cmd.add_argument(
        "-B", "--broker", type=str, default="localhost:9092",
        help="address of the Kafka broker to connect to",
    )
    play_cmd.add_argument(
        "--dst-topic", type=str, default="alerts-stream",
        help="name of the Kafka topic that will provide a live feed of the stream contents"
    )
    play_cmd.add_argument(
        "--dst-topic-partitions", type=int, default="1",
        help="number of partitions to create for the destination topic"
    )
    play_cmd.add_argument(
        "--src-topic", type=str, default="alerts-reservoir",
        help="name of the Kafka topic that is the source of the stream",
    )
    play_cmd.add_argument(
        "--force", action="store_true", help="overwrite dst-topic if it already exists",
    )
    play_cmd.add_argument(
        "--repeat-interval", type=int, default=-1,
        help="interval to repeat the stream, in seconds. <0 means don't repeat",
    )

    return parser
