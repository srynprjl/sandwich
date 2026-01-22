import json
from db import database
from logic import category

def parse_json():
  with open("plans.json", "r") as f:
    data = json.load(f)
  lists = []
  for i in data:
    data_tuple = (i["title"], i["shorthand"])
    lists.append(data_tuple)
  return lists


def main():
    data = parse_json()
    print(data)
    con, _ = database.connect_db("database")
    for name, shorthand in data:
      category.add_category(con, name, shorthand)

    database.close_db(con)


if __name__ == "__main__":
  main()
