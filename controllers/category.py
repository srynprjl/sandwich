import sqlite3

def category_exists(con: sqlite3.Connection, id: int = None, shorthand: str = None):
    if id is None and shorthand is None:
        return "Please insert at least the id or the shorthand"
    elif id and shorthand:
        sql = f"""SELECT id, name, shorthand FROM categories WHERE id={id} AND shorthand={shorthand}"""
    elif id:
        sql = f"""SELECT id, name, shorthand FROM categories WHERE id={id}"""
    elif shorthand:
        sql = (
            f"""SELECT id, name, shorthand FROM categories WHERE shorthand={shorthand}"""
        )
    else:
        return "This should never happen ig"
    cur = con.cursor()
    cur.execute(sql)
    data = cur.fetchone()
    return data if data else None


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
    return f"Category {name} successfully added."


def remove_category(con: sqlite3.Connection, id: int = None, shorthand: str = None):
    data = category_exists(con, id, shorthand)
    cur = con.cursor()
    if data:
        print(data)
        cur.execute(f"""DELETE FROM categories WHERE id={data[0]};""")
        con.commit()
        print("Successfully deleted")
    else:
        return f"Category {id} not found"
    return f"Category {id} successfully deleted. "


def update_category(
    con: sqlite3.Connection,
    new_data: dict[str, str],
    id: int = None,
    shorthand: str = None,
):
    data = category_exists(con, id, shorthand)
    cur = con.cursor()
    sql_piece = ""
    for k, v in new_data.items():
        if(type(v) == str):
            v = f"'{v}'"
        sql_piece += f"'{k}' = {v},"
    sql_piece = sql_piece.strip(",")
    # print(sql_piece)
    if data:
        sql = f"""UPDATE categories SET {sql_piece} WHERE id={data[0]}"""
        cur.execute(sql)
        con.commit()
        return "Success."
    else:
        return "Failed"


def get_all_categories(con: sqlite3.Connection):
    sql = """SELECT * from categories"""
    cur = con.cursor()
    cur.execute(sql)
    con.commit()
    data = cur.fetchall()
    return data


def get_category(con: sqlite3.Connection, id: int = None, shorthand: str = None):
    exists = category_exists(con, id, shorthand)
    if exists:
        cur = con.cursor()
        cur.execute(f"""SELECT * from categories WHERE id={exists[0]}""")
    else:
        return "BHAAAK!! No categories found"
