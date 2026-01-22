import sqlite3
from .category import category_exists
def schema_check(data: dict):
    required_fields ={"name", "category"}
    keys = set(data.keys())
    if (not required_fields.issubset(keys)):
        return None
    default_values = {
        "name": "",
        "description": "",
        "completed": False,
        "favorite": False,
        "path": "/",
        "category": 0
    }

    key_fields = ["name", "description", "completed", "favorite", "path", "category"]
    new_data = {}
    for key in key_fields:
      new_data[key] = data.get(key, default_values[key])
    return new_data

def project_exists(con: sqlite3.Connection, id: int):
    sql = f'''SELECT * from projects WHERE id={id}'''
    cur = con.cursor()
    cur.execute(sql)
    data = cur.fetchone()
    if(data):
        return data
    else:
        return None

def add_project(con: sqlite3.Connection, project_data: dict):
    cur = con.cursor()
    data = schema_check(project_data)
    if(data is None):
        return None
    sql = '''INSERT INTO projects (name, description, completed, favorite, path, category) VALUES'''
    values = ""
    for _, value in data.items():
        if(type(value) is str):
            value = f"'{value}'"
        values += f"{value}, "
    values = values.strip().strip(",")
    query = f"{sql}({values})"
    cur.execute(query)
    con.commit()

def delete_project(con: sqlite3.Connection, id: int):
    exists = project_exists(con, id)
    if(exists is None):
        return "Failed"
    sql = f'''DELETE FROM projects WHERE id={exists[0]}'''
    cur = con.cursor()
    cur.execute(sql)
    con.commit()
    return "Deleted"

def update_project(con: sqlite3.Connection, id: int, new_data: dict):
    exists = project_exists(con, id)
    if(exists is None):
        return None
    columns = ""
    for k, v in new_data.items():
        columns += f"{k} = {v},"
    columns = columns.strip(",")
    sql = f'''UPDATE projects SET {columns} WHERE id={exists[0]}'''
    cur = con.cursor()
    cur.execute(sql)
    con.commit()
    return "Updated"

def get_project(con: sqlite3.Connection, id: int):
    exists = project_exists(con, id)
    if(exists):
        return exists
    else:
        return None

def list_all_projects(con: sqlite3.Connection, category: int):
  exists = category_exists(con=con, id=category)
  if(exists is None):
      return []
  sql = f'''SELECT * from projects WHERE category = {exists[0]}'''
  cur = con.cursor()
  cur.execute(sql)
  data = cur.fetchall()
  return data

def get_completed_projects(con: sqlite3.Connection):
  sql = f'''SELECT * from projects WHERE completed=True'''
  cur = con.cursor()
  cur.execute(sql)
  data = cur.fetchall()
  return data


def get_fav_projects(con: sqlite3.Connection):
  sql = f'''SELECT * from projects WHERE favorite=True'''
  cur = con.cursor()
  cur.execute(sql)
  data = cur.fetchall()
  return data
