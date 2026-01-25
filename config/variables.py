import os

PROJECT_NAME = "sandwich"
AUTHOR = "sysnefo"
USER = os.environ.get("USER")

CONFIG_PATH = f"/home/{USER}/.config/{AUTHOR}/{PROJECT_NAME}"
DATABASE_PATH = f"/home/{USER}/.local/share/{AUTHOR}/{PROJECT_NAME}"
DATABASE_NAME = "database"
PORT = 5000
