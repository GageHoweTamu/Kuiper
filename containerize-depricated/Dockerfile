# Dockerfile
FROM python:3.9-slim

# Install necessary packages
RUN apt-get update && apt-get install -y \
    bash \
    && rm -rf /var/lib/apt/lists/*

# Copy the Python script and the bash script
COPY run_script.py /app/
COPY my_script.sh /app/

# Set the working directory
WORKDIR /app

# Install Python dependencies
RUN pip install docker

# Set the entrypoint
ENTRYPOINT ["python", "run_script.py"]