import asyncio
import datetime
import logging

import avro.datafile
import avro.io
import avro.schema
import astropy.time

logger = logging.getLogger("rubin-alert-sim.pacer")

class SimplePacer(object):
    def __init__(self, source_file):
        datum_reader = avro.io.DatumReader()
        self.data_file_reader = avro.datafile.DataFileReader(source_file, datum_reader)
        logger.info(f"loaded file {source_file}")
        logger.debug(f"loaded schema {self.data_file_reader.meta['avro.schema']}")

        self.schema = avro.schema.Parse(self.data_file_reader.meta['avro.schema'])
        first_alert = next(self.data_file_reader)
        self.time_offset = self._now() - alert_time(first_alert)

        # Initialization of the iteration loop
        self.next_alert = first_alert

    async def iterate(self):
        while True:
            next_alert = await self.next()
            if next_alert is None:
                return
            else:
                yield next_alert

    async def next(self):
        if self.next_alert is None:
            # End of iteration.
            return None

        logger.debug(f"alert time: {alert_time(self.next_alert)}")
        until = self.adjusted_time_until(self.next_alert).total_seconds()
        logger.debug(f"calculated wait period of {until}")
        if until > 0:
            logger.debug(f"waiting")
            await asyncio.sleep(until)
        return_val = self.next_alert
        try:
            self.next_alert = next(self.data_file_reader)
        except StopIteration:
            logger.debug("iteration done - next value will be None")
            self.next_alert = None
        return return_val

    def adjusted_time_until(self, alert):
        """Calculate the delta between 'now' and the alert data's adjusted timestamp.
        "Adjusted" refers to the shift applied to all timestamps in the alert
        file so that they appear to be in the future

        """
        return (alert_time(alert) + self.time_offset) - self._now()

    def _now(self):
        """ Alias for datetime.datetime.now() to suppport monkeypatching in tests. """
        return datetime.datetime.now()

    def close(self):
        self.fp.close()


def alert_time(alert):
    """ Convert a timestamp's midPointTai field into an datetime.datetime."""
    return astropy.time.Time(alert["diaSource"]["midPointTai"], format="mjd").to_datetime()
