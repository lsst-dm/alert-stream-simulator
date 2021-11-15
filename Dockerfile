# This Dockerfile has four stages:
#
# base-image
#   Updates the base Python image with security patches and common system
#   packages. This image becomes the base of all other images.
# build-image
#   Installs third-party dependencies and the app into a virtual
#   environment.
# runtime-image
#   - Copies the virtual environment into place.
#   - Runs a non-root user.
FROM python:3.9.7-buster AS base-image

# Update system packages
COPY script/install-base-packages.sh .
RUN ./install-base-packages.sh

FROM base-image AS build-image

# Create a Python virtual environment
ENV VIRTUAL_ENV=/opt/venv
RUN python -m venv $VIRTUAL_ENV
# Make sure we use the virtualenv
ENV PATH="$VIRTUAL_ENV/bin:$PATH"
# Put the latest pip and setuptools in the virtualenv
RUN pip install --upgrade --no-cache-dir pip setuptools wheel

COPY . /app
WORKDIR /app
RUN pip install --no-cache-dir .

# Load precomputed alert data
RUN make datasets

FROM base-image AS runtime-image

# Create a non-root user
RUN useradd --create-home appuser
WORKDIR /home/appuser

# Make sure we use the virtualenv
ENV PATH="/opt/venv/bin:$PATH"

COPY --from=build-image /opt/venv /opt/venv
COPY --from=build-image /app/data /var/sample_alert_data

# Switch to non-root user
USER appuser

# Run rubin-alert-stream command
ENTRYPOINT ["sh", "-c"]
CMD rubin-alert-sim --help
