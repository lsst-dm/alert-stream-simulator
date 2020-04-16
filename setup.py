from setuptools import setup

install_requires = ['avro-python3']

dev_requires = []

setup(name='rubin-alert-stream-simulator',
      version="0.1.0",
      description="Auxiliary code for simulating rthe Vera Rubin Observatory Alert Stream",
      url='https://github.com/lsst-dm/alert-stream-simulator',
      classifiers=[
          "Programming Language :: Python :: 3",
          "License :: OSI Approved :: GNU General Public License v3 (GPLv3)",
          "Development Status :: 3 - Alpha",
          "Operating System :: POSIX :: Linux",
      ],
      author='Spencer Nelson',
      author_email='swnelson@uw.edu',
      license='GPLv3',
      install_requires=install_requires,
      extras_require={"dev": dev_requires},
      zip_safe=False)
