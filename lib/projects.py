import sqlite3

from utils import database
from utils.variables import DATABASE_NAME

from .category import category_exists, get_category


def get_project_field_value(id: int, field: str):
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
    exists = project_exists(id)
    if not exists:
        return {"message": "Something went wrong!", "data": None, "status": 500}

    data = exists["data"][map_v.get(field)] if exists["data"] else None
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
        new_data[key] = (
            data.get(key) if data.get(key) is not None else default_values[key]
        )
    return new_data


def project_exists(id: int):
    con = database.connect_db(DATABASE_NAME)
    sql = f"""SELECT * from projects WHERE id={id}"""
    cur = con.cursor()
    try:
        cur.execute(sql)
        data = cur.fetchone()
        database.close_db(con)
        returnValue = data if data else None
        status = 201 if data else 404
        message = "exists" if data else "doesn't exist"
        return {
            "message": f"The project {message}",
            "status": status,
            "data": returnValue,
        }
    except sqlite3.Error:
        return None


def add_project(project_data: dict):
    con = database.connect_db(DATABASE_NAME)
    cur = con.cursor()
    data = schema_check(project_data)
    if data is None:
        return {"message": "Please provide the proper data", "status": 404}
    exists = category_exists(id=data["category"])
    if exists is None:
        return {"message": "Something happened"}
    if exists["status"] == 404:
        return {"message": "Category doesn't exist", "status": 404}
    sql = """INSERT INTO projects (name, description, completed, favorite, path, category) VALUES"""
    values = ""
    for _, value in data.items():
        if type(value) is str:
            value = f"'{value}'"
        values += f"{value}, "
    values = values.strip().strip(",")
    query = f"{sql}({values})"
    # print(query)
    try:
        cur.execute(query)
        con.commit()
        database.close_db(con)
        return {"message": "The project successfully added", "status": 201}
    except sqlite3.Error as e:
        return {"message": e, "data": None, "status": 201}


def delete_project(id: int, catId: int):
    con = database.connect_db(DATABASE_NAME)
    catExists = category_exists(catId)
    exists = project_exists(id)
    if catExists is None or exists is None:
        return {"message": "Something went wrong!", "status": 500}
    if catExists["status"] == 404:
        return {"message": "Category doesn't exist", "status": 404}
    if exists["status"] == 404:
        return {"message": "Project doesn't exist", "status": 404}
    catValue = get_project_field_value(id, "category")
    if catValue["data"] != catId:
        return {"message": "Project doesn't belong into the category", "status": 400}

    data = exists["data"]
    if data is None:
        return {"message": "No data exist to be deleted", "status": 500}
    sql = f"""DELETE FROM projects WHERE id={data[0]}"""
    cur = con.cursor()
    try:
        cur.execute(sql)
        con.commit()
        database.close_db(con)
        return {"message": "The project successfully deleted", "status": 201}
    except sqlite3.Error:
        return {
            "message": "Something went wrong! The project could not be deleted",
            "status": 500,
        }


def update_schema_check(data: dict):
    allowed_keys = (
        "name",
        "description",
        "path",
        "favorite",
        "completed",
        "category",
    )
    new_dict = {k: v for k, v in data.items() if k in allowed_keys}
    if new_dict == {}:
        return {"message": "No field found to update", "data": None, "status": 404}

    return {"message": "Data successfully parsed", "data": new_dict, "status": 20}


def update_project(id: int, new_data: dict, catId: int):
    con = database.connect_db(DATABASE_NAME)
    catExists = category_exists(catId)
    exists = project_exists(id)
    if catExists is None or exists is None:
        return {"message": "Something went wrong!", "status": 500}
    if catExists["status"] == 404:
        return {"message": "Category doesn't exist", "status": 404}
    if exists["status"] == 404:
        return {"message": "Project doesn't exist", "status": 404}
    catValue = get_project_field_value(id, "category")
    if catValue["data"] != catId:
        return {"message": "Project doesn't belong into the category", "status": 400}

    data = exists["data"]
    if data is None:
        return {"message": "Please provide the proper data", "status": 404}

    columns = ""
    parsed_data = update_schema_check(new_data)
    if not parsed_data["data"]:
        return parsed_data
    if parsed_data["data"].get("category", None) is None:
        parsed_data["data"]["category"] = catId
    if not category_exists(parsed_data["data"]["category"]):
        parsed_data["data"]["category"] = catId

    for k, v in parsed_data["data"].items():
        if type(v) is str:
            v = f"'{v}'"
        columns += f"{k} = {v},"
    columns = columns.strip(",")
    sql = f"""UPDATE projects SET {columns} WHERE id={data[0]}"""
    cur = con.cursor()
    try:
        cur.execute(sql)
        con.commit()
        return {"message": "The project successfully updated", "status": 201}
        database.close_db(con)
    except sqlite3.Error:
        return {"message": "Something went wrong!", "status": 500}


def get_project(id: int, catId: int):
    catExists = category_exists(catId)
    exists = project_exists(id)
    if catExists is None or exists is None:
        return {"message": "Something went wrong!", "data": None, "status": 500}
    if catExists["status"] == 404:
        return {"message": "The category doesn't exist", "data": None, "status": 404}
    if exists["status"] == 404:
        return {"message": "The project doesn't exist", "data": None, "status": 404}
    catValue = get_project_field_value(id=id, field="category")
    if catValue["data"] != catId:
        return {
            "message": "The project doesn't exist in the specified category",
            "data": None,
            "status": 404,
        }
    data = list(exists["data"])
    data[6] = get_category(id=catValue["data"])["data"][1]
    if data:
        return {
            "message": "The project successfully fetched",
            "data": data,
            "status": 201,
        }
    else:
        return {"message": "Something went wrong!", "data": None, "status": 500}


def list_all_projects(category: int):
    con = database.connect_db(DATABASE_NAME)
    exists = category_exists(id=category)
    if exists is None:
        return {"message": "Something went wrong!", "status": 500}
    if exists["status"] == 404:
        return exists
    data = exists["data"]
    if data is None:
        return exists
    sql = f"""SELECT * from projects WHERE category = {data[0]}"""
    cur = con.cursor()
    try:
        cur.execute(sql)
        data = cur.fetchall()
        database.close_db(con)
    except sqlite3.DataError:
        return {"message": "Something went wrong!", "status": 500}
    return {"message": "The project successfully fetched", "data": data, "status": 201}


def get_projects_cf(query: dict):
    con = database.connect_db(DATABASE_NAME)
    query_items = [
        f"{k}={int(bool(v))}"
        for k, v in query.items()
        if (k in ("favorite", "completed") and v in ("True", "False"))
    ]

    where_clause = " AND ".join(query_items)
    sql = f"""SELECT * from projects WHERE {where_clause}"""
    cur = con.cursor()
    try:
        cur.execute(sql)
        data = cur.fetchall()
        database.close_db(con)
    except sqlite3.Error as e:
        print(e)
        return {"message": "Something went wrong!", "data": None, "status": 201}
    return {"message": "The project successfully fetched", "data": data, "status": 201}
