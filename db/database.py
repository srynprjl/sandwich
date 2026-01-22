import sqlite3
import os

def connect_db(db_name: str):
  print()
  con = sqlite3.connect(f"{os.getcwd()}/db/{db_name}.db")
  return (con, con.cursor())

def close_db(con: sqlite3.Connection):
  con.close()

def create_initial_tables(cur: sqlite3.Cursor):
  sql_statements = [
    '''CREATE TABLE categories(id INTEGER PRIMARY KEY AUTOINCREMENT, name VARCHAR(50) NOT NULL, shorthand VARCHAR(20) UNIQUE)''',
    '''CREATE TABLE projects(id INTEGER PRIMARY KEY AUTOINCREMENT, name VARCHAR(255), description VARCHAR(255), completed BOOLEAN  CHECK (completed IN (0, 1)) , favorite BOOLEAN  CHECK (favorite IN (0, 1)) , path VARCHAR(1000), category INTEGER, FOREIGN KEY (id) REFERENCES categories(id))'''
  ]
  for sql in sql_statements:
    try:
      cur.execute(sql)
    except sqlite3.Error as e:
      # print("Error", e)
      continue
