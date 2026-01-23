import sqlite3
from .category import category_exists

def schema_check(data: dict):
    required_fields = {"name", "category"}
    keys = set(data.keys())
    if not required_fields.issubset(keys):
        return None
    default_values = {
        "name": "",
        "description": "",
        "completed": False,
        "favorite": False,
        "path": "/",
        "category": 0,
    }

    key_fields = ["name", "description", "completed", "favorite", "path", "category"]
    new_data = {}
    for key in key_fields:
        new_data[key] = data.get(key, default_values[key])
    return new_data


def project_exists(con: sqlite3.Connection, id: int):
    sql = f"""SELECT * from projects WHERE id={id}"""
    cur = con.cursor()
    cur.execute(sql)
    data = cur.fetchone()
    returnValue = data if data else None
    status = 201 if data else 404
    message = "exists" if data else "doesn't exist"
    return {"message": f"The project {message}", "status": status, "data": returnValue}


def add_project(con: sqlite3.Connection, project_data: dict):
    cur = con.cursor()
    data = schema_check(project_data)
    if data is None:
        return {"message": "Please provide the proper data", "status": 404}
    sql = """INSERT INTO projects (name, description, completed, favorite, path, category) VALUES"""
    values = ""
    for _, value in data.items():
        if type(value) is str:
            value = f"'{value}'"
        values += f"{value}, "
    values = values.strip().strip(",")
    query = f"{sql}({values})"
    cur.execute(query)
    con.commit()
    return {"message": "The project successfully added", "status": 201}


def delete_project(con: sqlite3.Connection, id: int):
    exists = project_exists(con, id)
    data = exists["data"]
    if data is None:
        return {"message": "The deletion failed", "status": 400}
    sql = f"""DELETE FROM projects WHERE id={data[0]}"""
    cur = con.cursor()
    cur.execute(sql)
    con.commit()
    return {"message": "The project successfully deleted", "status": 201}


def update_project(con: sqlite3.Connection, id: int, new_data: dict):
    exists = project_exists(con, id)
    data = exists["data"]
    if data is None:
        return {"message": "Please provide the proper data", "status": 404}
    columns = ""
    for k, v in new_data.items():
        if type(v) is str:
            v = f"'{v}'"
        columns += f"{k} = {v},"
    columns = columns.strip(",")
    sql = f"""UPDATE projects SET {columns} WHERE id={data[0]}"""
    cur = con.cursor()
    cur.execute(sql)
    con.commit()
    return {"message": "The project successfully updated", "status": 201}


def get_project(con: sqlite3.Connection, id: int):
    exists = project_exists(con, id)
    data = exists["data"]
    if data:
        return {
            "message": "The project successfully fetched",
            "data": data,
            "status": 201,
        }
    else:
        return {"message": "The project doesn't exist", "data": data, "status": 201}


def list_all_projects(con: sqlite3.Connection, category: int):
    exists = category_exists(con=con, id=category)
    data = exists["data"]
    print
    if data is None:
        return {"message": "No Data found", "data": data, "status": 404}
    sql = f"""SELECT * from projects WHERE category = {data[0]}"""
    cur = con.cursor()
    cur.execute(sql)
    data = cur.fetchall()
    return {"message": "The project successfully fetched", "data": data, "status": 201}


def get_completed_projects(con: sqlite3.Connection):
    sql = f"""SELECT * from projects WHERE completed=True"""
    cur = con.cursor()
    cur.execute(sql)
    data = cur.fetchall()
    return {"message": "The project successfully fetched", "data": data, "status": 201}


def get_fav_projects(con: sqlite3.Connection):
    sql = f"""SELECT * from projects WHERE favorite=True"""
    cur = con.cursor()
    cur.execute(sql)
    data = cur.fetchall()
    return {"message": "The project successfully fetched", "data": data, "status": 201}


def get_project_field_value(con: sqlite3.Connection, id: int, field: str):
    if field not in (
        "name",
        "description",
        "path",
        "favorite",
        "completed",
        "category",
    ):
        return {
            "message": f"The {field} not found",
            "data": None,
            "status": 400,
        }
    map_v = {
        "name": 1,
        "description": 2,
        "path": 5,
        "favorite": 4,
        "completed": 3,
        "category": 6,
    }
    exists = project_exists(con, id)
    data = exists["data"][map_v.get(field)]
    if data:
        return {
            "message": f"The {field} successfully fetched",
            "data": data,
            "status": 201,
        }
    else:
        return {
            "message": f"The {field} not found",
            "data": None,
            "status": 400,
        }
