import os
import unittest


LICENSE_PREAMBLE = """# This file is part of alert-stream-simulator.
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
"""


class TestLicenseHeaders(unittest.TestCase):
    def testFilesHaveLicenseHeader(self):
        repo_root = os.path.dirname(os.path.dirname(__file__))
        py_files = self.gather_python_files(os.path.join(repo_root, "python"))
        py_files += self.gather_python_files(os.path.join(repo_root, "test"))

        for filename in py_files:
            if not self.file_contains(filename, LICENSE_PREAMBLE):
                self.fail(f"{filename} is missing licensing header")

    @staticmethod
    def gather_python_files(source_dir):
        """Traverse source_dir and all subdirectories looking for files that
        end with .py. Put their filenames all in a list and return it.
        """

        python_files = []
        for (dirpath, _, filenames) in os.walk(source_dir):
            for filename in filenames:
                if filename.endswith(".py"):
                    python_files.append(os.path.join(dirpath, filename))
        return python_files

    @staticmethod
    def file_contains(filename, substring):
        """ Returns true iff the file at filename contains substring."""
        with open(filename, "r") as fp:
            contents = fp.read()
            return substring in contents


if __name__ == "__main__":
    unittest.main()
