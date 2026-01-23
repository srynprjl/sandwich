import argparse
import os
import shutil
import sqlite3
import subprocess

from flask.app import Flask

from api.app import run_server
from api.routes.projects import project
from config.variables import PORT
from controllers import category, projects

from .parser import parse_arguments

parser, args = parse_arguments()


def category_logics(con: sqlite3.Connection):
    match args.action:
        case "add":
            name = args.name
            shorthand = args.shorthand
            category.add_category(con, name=name, shorthand=shorthand)

        case "delete":
            id = args.id
            category.remove_category(con=con, id=id)

        case "update":
            if args.name is None and args.shorthand is None:
                print("At least one argument [name or shorthand] is required")
                exit(1)
            new_data = {}
            if args.name is not None:
                new_data["name"] = args.name
            if args.shorthand is not None:
                new_data["shorthand"] = args.shorthand
            id = args.id
            category.update_category(con=con, id=id, new_data=new_data)

        case "list":
            argument = args.id
            if argument is None:
                categories = category.get_all_categories(con)
                data = categories["data"]
                print("ID || Name")
                for id, name, shorthand in data:
                    print(f" {id} ||  {name}")
            else:
                project_list = projects.list_all_projects(con=con, category=argument)
                data = project_list["data"]
                print("ID || Name || Completed || Favorite || Description")
                for id, name, description, completed, favorite, *rest in data:
                    fav = True if favorite == 1 else False
                    com = True if completed == 1 else False
                    print(f" {id} ||  {name} || {com} || {fav} || {description}")
        case _:
            sub_action = next(
                a for a in parser._actions if isinstance(a, argparse._SubParsersAction)
            )
            sub_action.choices["category"].print_help()


def project_logics(con: sqlite3.Connection):
    match args.action:
        case "add":
            name = args.name
            description = args.description
            category = args.category
            path = args.path
            fav = args.favourite
            completed = args.complete
            data = {
                "name": name,
                "description": description,
                "category": category,
                "path": path,
                "favorite": fav,
                "completed": completed,
            }
            projects.add_project(con, project_data=data)

        case "delete":
            id = args.id
            catVal = projects.get_project_field_value(con, args.id, "category")
            projects.delete_project(con=con, id=id, catId=catVal["data"])

        case "update":
            data = {
                "name": args.name,
                "description": args.description,
                "category": args.category,
                "path": args.path,
                "favorite": args.favourite,
                "completed": args.complete,
            }
            catVal = projects.get_project_field_value(con, args.id, "category")
            data = {k: v for k, v in data.items() if v is not None}
            projects.update_project(
                con, id=args.id, new_data=data, catId=catVal["data"]
            )

        case "view":
            id = args.id
            catVal = projects.get_project_field_value(con, args.id, "category")
            project = projects.get_project(con=con, id=id, catId=catVal["data"])
            if not project["data"]:
                print(f"{project['message']}")
                return
            fav = True if project["data"][4] == 1 else False
            com = True if project["data"][3] == 1 else False
            print(f"ID -> {project['data'][0]} ")
            print(f"Name -> {project['data'][1]}")
            print(f"Description -> {project['data'][2]}")
            print(f"Path -> {project['data'][5]}")
            print(f"Favorite -> {fav}")
            print(f"Complete -> {com}")

        case "favourite":
            id = args.id
            catVal = projects.get_project_field_value(con, args.id, "category")
            fav_value = projects.get_project_field_value(
                con=con, id=id, field="favorite"
            )
            fav = 0 if fav_value == 1 else 1
            projects.update_project(
                con=con, id=id, new_data={"favorite": fav}, catId=catVal["data"]
            )

        case "complete":
            id = args.id
            catVal = projects.get_project_field_value(con, args.id, "category")
            completed_val = projects.get_project_field_value(
                con=con, id=id, field="completed"
            )
            com = 0 if completed_val == 1 else 1
            projects.update_project(
                con=con, id=id, new_data={"completed": com}, catId=catVal["data"]
            )
        case _:
            sub_action = next(
                a for a in parser._actions if isinstance(a, argparse._SubParsersAction)
            )
            sub_action.choices["project"].print_help()

    try:
        if args.editor:
            id = args.id
            path = projects.get_project_field_value(con=con, id=id, field="path")
            if not path["data"]:
                print(path["message"])
                return
            match args.editor:
                case "code":
                    executable = shutil.which("code")
                case "zed":
                    executable = shutil.which("zeditor")
                case "nvim":
                    executable = shutil.which("nvim")
                case _:
                    executable = os.environ.get("EDITOR", "vi")
            if executable:
                subprocess.run([executable, path["data"]])
            else:
                print(f"Editor '{executable}' not found.")
    except AttributeError:
        pass


def cli(con: sqlite3.Connection):
    if args.web == "on":
        run_server(PORT)
    if args.command == "category":
        category_logics(con)
    elif args.command == "project":
        project_logics(con)

    if args.completed or args.favourite:
        dict = {"completed": args.completed, "favorite": args.favourite}
        query = {k: v for k, v in dict.items() if v == "True"}
        data = projects.get_projects_cf(con, query)
        print("ID || Name || Completed || Favorite || Description")
        for id, name, description, completed, favorite, *rest in data["data"]:
            fav = True if favorite == 1 else False
            com = True if completed == 1 else False
            print(f" {id} ||  {name} || {com} || {fav} || {description}")
