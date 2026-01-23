import sqlite3


def category_exists(con: sqlite3.Connection, id: int = None, shorthand: str = None):
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
    returnValue = data if data else None
    messageValue = "exists" if data else "doesn't exist"
    status = 200 if data else 404
    return {
        "message": f"The category {messageValue}",
        "data": returnValue,
        "status": status,
    }


def add_category(
    con: sqlite3.Connection, name: str, shorthand: str, description: str = None
):
    sql = f"""INSERT INTO categories(name, shorthand) VALUES("{name}", "{shorthand}")"""
    cur = con.cursor()
    try:
        cur.execute(sql)
        con.commit()
    except:
        pass
    return {"message": "Category successfully created.", "status": 201}


def remove_category(con: sqlite3.Connection, id: int = None, shorthand: str = None):
    data = category_exists(con, id, shorthand)
    cur = con.cursor()
    if data["status"] not in (404, 400):
        cur.execute(f"""DELETE FROM categories WHERE id={data["data"][0]};""")
        con.commit()
        return {"message": f"Category {id} successfully deleted", "status": 201}
    else:
        return {"message": f"Category {id} not found", "status": 404}


def update_category(
    con: sqlite3.Connection,
    new_data: dict[str, str],
    id: int = None,
    shorthand: str = None,
):
    exists = category_exists(con, id, shorthand)
    data = exists["data"]
    cur = con.cursor()
    sql_piece = ""
    for k, v in new_data.items():
        if type(v) == str:
            v = f"'{v}'"
        sql_piece += f"'{k}' = {v},"
    sql_piece = sql_piece.strip(",")
    if data:
        sql = f"""UPDATE categories SET {sql_piece} WHERE id={data[0]}"""
        cur.execute(sql)
        con.commit()
        return {"message": "Data successfully updated.", "status": 201}
    else:
        return {"message": "Data not found", "status": 404}


def get_all_categories(con: sqlite3.Connection):
    sql = """SELECT * from categories"""
    cur = con.cursor()
    cur.execute(sql)
    con.commit()
    data = cur.fetchall()
    return {"message": "Data fetched successfully", "data": data, "status": 201}


def get_category(con: sqlite3.Connection, id: int = None, shorthand: str = None):
    exists = category_exists(con, id, shorthand)
    data = exists["data"]
    if exists:
        cur = con.cursor()
        cur.execute(f"""SELECT * from categories WHERE id={data[0]}""")
        return {"message": "Data fetched successfully", "data": data, "status": 201}
    else:
        return {"message": "Data fetched successfully", "data": data, "status": 201}
