import sys

from cli.actions import cli
from db import database


def main():
    database.create_initial_tables()
    if len(sys.argv) == 1:
        pass
    else:
        cli()


if __name__ == "__main__":
    main()
