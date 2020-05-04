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

import asyncio
import argparse
import logging
import datetime
import concurrent.futures

from streamsim import sender, pacer, fastavro_pacer, replay

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
    if args.subcommand == "preload":
        raise NotImplementedError("preload isn't written yet")
    if args.subcommand == "replay":
        replay_cmd(args.broker, args.src_topic, args.dst_topic, args.timeout_sec, args.block_size)


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

    preload_cmd = subparsers.add_parser(
        "preload", help="prepare and load data into Kafka for replace",
        formatter_class=argparse.ArgumentDefaultsHelpFormatter,
    )
    preload_cmd.add_argument(
        "-B", "--broker", type=str, default="localhost:9092",
        help="address of the Kafka broker to connect to",
    )
    preload_cmd.add_argument(
        "--topic", type=str, default="alerts",
        help="Kafka topic name to use",
    )

    replay_cmd = subparsers.add_parser(
        "replay", help="play back a stream that's been loaded into kafka",
        formatter_class=argparse.ArgumentDefaultsHelpFormatter,
    )
    replay_cmd.add_argument(
        "-B", "--broker", type=str, default="localhost:9092",
        help="address of the Kafka broker to connect to",
    )
    replay_cmd.add_argument(
        "--src-topic", type=str, default="alerts",
        help="Kafka topic name to source data from",
    )
    replay_cmd.add_argument(
        "--dst-topic", type=str, default="alerts-replay",
        help="Kafka topic name to source data from",
    )
    replay_cmd.add_argument(
        "--timeout-sec", type=float, default=10.0,
        help="how long, in seconds, to wait before giving up when flushing a write to kafka",
    )
    replay_cmd.add_argument(
        "--block-size", type=int, default=100,
        help="number of messages to retrieve per consume call when replaying",
    )


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


def replay_cmd(broker, src_topic, dst_topic, timeout, block_size):
    replayer = replay.Replayer(broker, src_topic, dst_topic, timeout, block_size)
    replayer.replay_loop(block_size)
    replayer.close()
