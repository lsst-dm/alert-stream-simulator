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

from streamsim import creator, player, printer, _kafka


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

    try:
        tls_config = _unpack_tls_arguments(args)
    except AttributeError:
        tls_config = None
    if args.subcommand == "create-stream":
        logging.debug(f"dispatching create-stream command with args: {args}")
        n = creator.create(args.broker, args.dst_topic, args.file, args.timeout_sec,
                           args.force, tls_config, schema_id=args.schema_id)
        print(f"successfully preloaded stream with {n} alerts")
    elif args.subcommand == "play-stream":
        logging.debug(f"dispatching play-stream command with args: {args}")
        n = player.play(args.broker, args.src_topic, args.dst_topic, args.dst_topic_partitions,
                        args.create_dst_topic, args.repeat_interval, tls_config)
        print(f"played {n} alerts from the stream")
    elif args.subcommand == "print-stream":
        logging.debug(f"dispatching print-stream command with args: {args}")
        printer.print_stream(args.broker, args.src_topic, tls_config)
    else:
        parser.print_usage()


def _add_tls_arguments(subcommand):
    subcommand.add_argument(
        "--tls-client-key-location", type=str, default="",
        help="path to a client PEM key used for mTLS authentication"
    )
    subcommand.add_argument(
        "--tls-client-crt-location", type=str, default="",
        help="path to a client public cert used for mTLS authentication"
    )
    subcommand.add_argument(
        "--tls-server-ca-crt-location", type=str, default="",
        help="path to a CA public cert used to verify the server's TLS cert"
    )


def _unpack_tls_arguments(args):
    # Params should either be all empty, or all set
    parameters = [
        args.tls_server_ca_crt_location,
        args.tls_client_key_location,
        args.tls_client_crt_location,
    ]
    if all(x == "" for x in parameters):
        return None

    if any(x == "" for x in parameters):
        raise ValueError("either all --tls options must be set, or none of them")

    return _kafka.TLSConfig(
        args.tls_client_key_location,
        args.tls_client_crt_location,
        args.tls_server_ca_crt_location,
    )


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
        "--schema-id", type=int, default=1,
        help="Confluent Schema Registry schema ID to use in alert packet header",
    )
    _add_tls_arguments(create_cmd)
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
        "--create-dst-topic", action="store_true",
        help="create dst-topic, overwriting if it already exists",
    )
    play_cmd.add_argument(
        "--src-topic", type=str, default="alerts-reservoir",
        help="name of the Kafka topic that is the source of the stream",
    )
    play_cmd.add_argument(
        "--repeat-interval", type=int, default=-1,
        help="interval to repeat the stream, in seconds. <0 means don't repeat",
    )
    _add_tls_arguments(play_cmd)

    print_cmd = subparsers.add_parser(
        "print-stream", help="print the size of messages in the stream in real time"
    )
    print_cmd.add_argument(
        "-B", "--broker", type=str, default="localhost:9092",
        help="address of the Kafka broker to connect to",
    )
    print_cmd.add_argument(
        "--src-topic", type=str, default="alerts-reservoir",
        help="name of the Kafka topic that is the source of the stream",
    )
    _add_tls_arguments(print_cmd)

    return parser
