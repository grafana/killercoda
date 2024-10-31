# Installation

1. Clone the repository

```bash
git clone https://github.com/grafana/adventure.git
```{{exec}}

1. Spin up the Observability Stack using Docker Compose

```bash
docker-compose up -d
```{{exec}}

Quest World runs as a python application our recommended way to install it is to use a virtual environment.

1. Create a virtual environment

```bash
python3 -m venv .venv
```{{exec}}

1. Activate the virtual environment

```bash
source .venv/bin/activate
```{{exec}}

1. Install the required dependencies

```bash
pip install -r requirements.txt
```{{exec}}

1. Run the application

```bash
python main.py
```{{exec}}
