import sys

from cli.actions import cli
from db import database


def main():
    from api.app import run_server

    con = database.db()
    database.create_initial_tables(con=con)
    if len(sys.argv) == 1:
        # run_server(5000)
        pass
    else:
        cli(con)


if __name__ == "__main__":
    main()
