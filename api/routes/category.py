import sqlite3

from flask import Blueprint, request

from lib import category

categories = Blueprint("categories", __name__)


@categories.get("/api/category/")
def get_all_category_api():
    data = category.get_all_categories()
    return data


@categories.get("/api/category/<int:id>/")
def get_category(id):
    data = category.get_category(id=id)
    return data


@categories.post("/api/category/")
def add_category():
    data = request.get_json()
    if "name" not in data.keys() and "shorthand" not in data.keys():
        return "Failed"
    try:
        resp = category.add_category(name=data["name"], shorthand=data["shorthand"])
        return resp
    except KeyError:
        return {"message": "Error", "status": 400}


@categories.patch("/api/category/<int:id>/")
def update_category(id):
    if type(id) is not int:
        return {"message": "ID Required", "status": 400}
    data = request.get_json()
    try:
        resp = category.update_category(new_data=data, id=id)
        return resp
    except sqlite3.OperationalError:
        return {"message": "Bad Request", "status": 404}


@categories.delete("/api/category/<int:id>")
def delete_category(id):
    req = category.remove_category(id)
    return req
