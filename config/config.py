import os
import json
from variables import CONFIG_PATH
DEFAULT_CONFIG = {
    "default_tables": [
        {
            "name": "categories",
            "fields": [
                "id INTEGER PRIMARY KEY AUTOINCREMENT",
                "name VARCHAR(50) NOT NULL",
                "shorthand VARCHAR(20) UNIQUE",
            ],
        },
        {
            "name": "projects",
            "fields": [
                "id INTEGER PRIMARY KEY AUTOINCREMENT",
                "name VARCHAR(255) NOT NULL",
                "description VARCHAR(255)",
                "completed BOOLEAN  CHECK (completed IN (0, 1))",
                "favorite BOOLEAN  CHECK (favorite IN (0, 1))",
                "path VARCHAR(1000)",
                "category INTEGER",
                "FOREIGN KEY (id) REFERENCES categories(id))",
            ],
        },
    ]
}

def create_config():
  if(not os.path.isdir(CONFIG_PATH)):
    os.mkdir(CONFIG_PATH)

  if(not os.path.exists(os.path.join(CONFIG_PATH, "config.json"))):
    with open(os.path.join(CONFIG_PATH, "config.json"), "w") as f:
      json.dump(DEFAULT_CONFIG, f)

def read_config():
  try:
    with open(os.path.join(CONFIG_PATH, "config.json"), "r") as f:
      config = json.load(f)
      return config
  except FileNotFoundError:
    create_config()
    read_config()
  except json.JSONDecodeError:
    print("JSON File cannot be decoded")
