import os
import tomllib


PROJECT_NAME = "sandwich"
AUTHOR = "sysnefo"
USER = os.environ.get("USER")
VERSION = "0.0.1"  # major.minor.fix
CONFIG_PATH = f"/home/{USER}/.config/{AUTHOR}/{PROJECT_NAME}"
DATABASE_PATH = f"/home/{USER}/.local/share/{AUTHOR}/{PROJECT_NAME}"
DATABASE_NAME = "sandwich_data"
PORT = 5000
