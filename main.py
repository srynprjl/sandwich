# from test import crud_test
# from controllers import projects
from db import database
from cli.actions import cli
def main():
    con = database.connect_db("database")
    database.create_initial_tables(con=con)
    cli(con)

if __name__ == "__main__":
    main()
