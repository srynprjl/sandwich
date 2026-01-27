import sqlite3

from utils import database
from utils.variables import DATABASE_NAME


def category_exists(id: (int | None) = None, shorthand: (str | None) = None):
    con = database.connect_db(DATABASE_NAME)
    if id is None and shorthand is None:
        return {
            "message": "Please insert at least the id or the shorthand",
            "data": None,
            "status": 400,
        }

    filters = []
    params = []
    if id is not None:
        filters.append("id = ?")
        params.append(id)

    if shorthand is not None:
        filters.append("shorthand = '?'")
        params.append(shorthand)

    where_clause = " AND ".join(filters)
    sql = f"SELECT id, name, shorthand FROM categories WHERE {where_clause}"
    cur = con.cursor()
    try:
        cur.execute(sql, params)
    except sqlite3.Error:
        return None
    data = cur.fetchone()
    database.close_db(con)
    returnValue = data if data else None
    messageValue = "exists" if data else "doesn't exist"
    status = 200 if data else 404
    return {
        "message": f"The category {messageValue}",
        "data": returnValue,
        "status": status,
    }


def add_category(name: str, shorthand: str, description: (str | None) = None):
    con = database.connect_db(DATABASE_NAME)
    sql = f"""INSERT INTO categories(name, shorthand) VALUES("{name}", "{shorthand}")"""
    cur = con.cursor()
    try:
        cur.execute(sql)
        con.commit()

    except sqlite3.Error:
        database.close_db(con)
        return {"message": "Something went wrong", "status": 500}

    database.close_db(con)
    return {"message": "Category successfully created.", "status": 201}


def remove_category(id: (int | None) = None, shorthand: (str | None) = None):
    con = database.connect_db(DATABASE_NAME)
    data = category_exists(id, shorthand)
    if not data:
        return {"message": "Something went wrong", "status": 500}
    cur = con.cursor()
    if data["status"] not in (404, 400):
        cur.execute(f"""DELETE FROM categories WHERE id={data["data"][0]};""")
        con.commit()
        database.close_db(con)
        return {"message": f"Category {id} successfully deleted", "status": 201}
    else:
        database.close_db(con)
        return {"message": f"Category {id} not found", "status": 404}


def update_category(
    new_data: dict[str, str],
    id: (int | None) = None,
    shorthand: (str | None) = None,
):
    con = database.connect_db(DATABASE_NAME)
    exists = category_exists(id, shorthand)
    if not exists:
        return {"message": "Something went wrong", "status": 500}
    data = exists["data"]
    cur = con.cursor()
    sql_piece = ""
    for k, v in new_data.items():
        if type(v) is str:
            v = f"'{v}'"
        sql_piece += f"'{k}' = {v},"
    sql_piece = sql_piece.strip(",")
    if data:
        sql = f"""UPDATE categories SET {sql_piece} WHERE id={data[0]}"""
        cur.execute(sql)
        con.commit()
        database.close_db(con)
        return {"message": "Data successfully updated.", "status": 201}
    else:
        database.close_db(con)
        return {"message": "Data not found", "status": 404}


def get_all_categories():
    con = database.connect_db(DATABASE_NAME)
    sql = """SELECT * from categories"""
    cur = con.cursor()
    try:
        cur.execute(sql)
        con.commit()
        data = cur.fetchall()
        database.close_db(con)
        return {"message": "Data fetched successfully", "data": data, "status": 201}
    except sqlite3.Error:
        database.close_db(con)
        return {"message": "Data not found", "status": 404}


def get_category(id: (int | None) = None, shorthand: (str | None) = None):
    con = database.connect_db(DATABASE_NAME)
    exists = category_exists(id, shorthand)
    if not exists:
        return {"message": "Something went wrong", "status": 500}
    data = exists["data"]
    if exists:
        cur = con.cursor()
        cur.execute(f"""SELECT * from categories WHERE id={data[0]}""")
        database.close_db(con)
        return {"message": "Data fetched successfully", "data": data, "status": 201}
    else:
        database.close_db(con)
        return {"message": "Data not found", "data": [], "status": 201}
