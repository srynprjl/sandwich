from db import database
from controllers import category
import json

def parse_json():
    with open("plans.json", "r") as f:
        data = json.load(f)
    lists = []
    for i in data:
        data_tuple = (i["title"], i["shorthand"])
        lists.append(data_tuple)
    return lists

def test_main():
    data = parse_json()
    con  = database.connect_db("test")
    database.create_initial_tables(con)
    for name, shorthand in data:
        category.add_category(con, name, shorthand)
    database.close_db(con)
