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
import io

import astropy.time

from lsst.alert.packet import SchemaRegistry


alert_schema = SchemaRegistry.from_filesystem().get_by_version("2.1")


def mock_alert(alert_id, timestamp):
    """Generate a minimal mock alert. Timestamp should be an ISO time
    string.

    Parameters
    ----------
    alert_id : `int`
        An integer to pass in as the `alertId` field of the alert.
    timestamp : `str`
        An ISO8601-encoded timestamp string.

    Returns
    -------
    alert : `dict`
        A mock alert with most fields set to zero.
    """
    return {
        "alertId": alert_id,
        "diaSource": {
            "midPointTai": astropy.time.Time(timestamp).mjd,
            # Below are all the required fields. Set them to zero.
            "diaSourceId": 0,
            "ccdVisitId": 0,
            "filterName": "",
            "programId": 0,
            "ra": 0,
            "decl": 0,
            "x": 0,
            "y": 0,
            "apFlux": 0,
            "apFluxErr": 0,
            "snr": 0,
            "psFlux": 0,
            "psFluxErr": 0,
            "flags": 0,
        }
    }


def mock_alert_file(alerts):
    """Generate a mock serialized file of alerts from a list.

    Parameters
    ----------
    alerts : `list` of `dict`
        A list of alert dicts to serialize into a mock file.

    Returns
    -------
    mock_file : `io.BytesIO`
        A mock file in memory, seeked to the start of the file, which
        contains serialized bytes describing alerts.
    """
    mock_file = io.BytesIO()
    alert_schema.store_alerts(mock_file, alerts)
    mock_file.seek(0)
    return mock_file
