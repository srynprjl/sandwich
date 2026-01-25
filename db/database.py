import os
import sqlite3

from config.variables import DATABASE_NAME, DATABASE_PATH


def create_database_location():
    if not os.path.isdir(DATABASE_PATH):
        os.makedirs(DATABASE_PATH, exist_ok=True)


def connect_db(db_name: str):
    try:
        con = sqlite3.connect(
            os.path.join(DATABASE_PATH, f"{db_name}.db"), check_same_thread=False
        )
        return con
    except sqlite3.Error:
        create_database_location()
        return connect_db(db_name)


def close_db(con: sqlite3.Connection):
    con.close()


def create_initial_tables():
    con = connect_db(DATABASE_NAME)
    cur = con.cursor()
    sql_statements = [
        """CREATE TABLE categories(id INTEGER PRIMARY KEY AUTOINCREMENT, name VARCHAR(50) NOT NULL, shorthand VARCHAR(20) UNIQUE)""",
        """CREATE TABLE projects(id INTEGER PRIMARY KEY AUTOINCREMENT, name VARCHAR(255), description VARCHAR(255), completed BOOLEAN  CHECK (completed IN (0, 1)) , favorite BOOLEAN  CHECK (favorite IN (0, 1)) , path VARCHAR(1000), category INTEGER, FOREIGN KEY (id) REFERENCES categories(id))""",
    ]
    for sql in sql_statements:
        try:
            cur.execute(sql)
        except sqlite3.Error as e:
            continue
    close_db(con)
