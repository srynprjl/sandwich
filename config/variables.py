import os
PROJECT_NAME="sandwich"
AUTHOR="sysnefo"
USER = os.environ.get("USER")
# XDG_CONFIG_HOME = os.environ.get("XDG_CONFIG_HOME")
# XDG_DATA_HOME = os.environ.get("XDG_DATA_HOME")

CONFIG_PATH = f"/home/{USER}/.config/{AUTHOR}/{PROJECT_NAME}"
DATABASE_PATH = f"/home/{USER}/.local/share/{AUTHOR}/{PROJECT_NAME}"
