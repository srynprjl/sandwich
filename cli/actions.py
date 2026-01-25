import argparse
import os
import shutil
import sqlite3
import subprocess

import tabulate
from flask.app import Flask

from api.app import run_server
from api.routes.projects import project
from config.variables import PORT
from controllers import category, projects

from .parser import parse_arguments

parser, args = parse_arguments()


def category_logics():
    match args.action:
        case "add":
            name = args.name
            shorthand = args.shorthand
            category.add_category(name=name, shorthand=shorthand)

        case "delete":
            id = args.id
            category.remove_category(id=id)

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
            category.update_category(id=id, new_data=new_data)

        case "list":
            argument = args.id
            if argument is None:
                categories = category.get_all_categories()
                data = categories["data"]
                col_headers = ("ID", "Name", "Shorthand")
                # print("ID || Name")
                # for id, name, shorthand in data:
                #     print(f" {id} ||  {name}")
                print(tabulate.tabulate(data, headers=col_headers))
            else:
                project_list = projects.list_all_projects(category=argument)
                data = project_list["data"]
                col_headers = ("ID", "Name", "Completed", "Favourite", "Description")
                table_data = []
                if not data:
                    print("No category found with that ID")
                    return
                for id, name, description, completed, favorite, *rest in data:
                    fav = True if favorite == 1 else False
                    com = True if completed == 1 else False
                    table_data.append((id, name, com, fav, description))
                print(tabulate.tabulate(table_data, headers=col_headers))
        case _:
            sub_action = next(
                a for a in parser._actions if isinstance(a, argparse._SubParsersAction)
            )
            sub_action.choices["category"].print_help()


def project_logics():
    match args.action:
        case "add":
            name = args.name
            description = args.description
            cat = args.category
            path = args.path
            fav = args.favourite
            completed = args.complete
            data = {
                "name": name,
                "description": description,
                "category": cat,
                "path": path,
                "favorite": fav,
                "completed": completed,
            }
            projects.add_project(project_data=data)

        case "delete":
            id = args.id
            catVal = projects.get_project_field_value(args.id, "category")
            projects.delete_project(id=id, catId=catVal["data"])

        case "update":
            data = {
                "name": args.name,
                "description": args.description,
                "category": args.category,
                "path": args.path,
                "favorite": args.favourite,
                "completed": args.complete,
            }
            catVal = projects.get_project_field_value(args.id, "category")
            data = {k: v for k, v in data.items() if v is not None}
            projects.update_project(id=args.id, new_data=data, catId=catVal["data"])

        case "view":
            id = args.id
            catVal = projects.get_project_field_value(args.id, "category")
            project = projects.get_project(id=id, catId=catVal["data"])
            if not project["data"]:
                print(f"{project['message']}")
                return
            project["data"] = list(project["data"])
            project["data"][4] = True if project["data"][4] == 1 else False
            project["data"][3] = True if project["data"][3] == 1 else False
            columns = (
                "ID",
                "Name",
                "Description",
                "Completed",
                "Favorite",
                "Path",
                "Category",
            )
            project_datas = [i for i in zip(columns, project["data"])]
            print(tabulate.tabulate(project_datas))

        case "favourite":
            id = args.id
            catVal = projects.get_project_field_value(args.id, "category")
            fav_value = projects.get_project_field_value(id=id, field="favorite")
            fav = 0 if fav_value == 1 else 1
            projects.update_project(
                id=id, new_data={"favorite": fav}, catId=catVal["data"]
            )

        case "complete":
            id = args.id
            catVal = projects.get_project_field_value(args.id, "category")
            completed_val = projects.get_project_field_value(id=id, field="completed")
            com = 0 if completed_val == 1 else 1
            projects.update_project(
                id=id, new_data={"completed": com}, catId=catVal["data"]
            )
        case _:
            sub_action = next(
                a for a in parser._actions if isinstance(a, argparse._SubParsersAction)
            )
            sub_action.choices["project"].print_help()

    try:
        if args.editor:
            id = args.id
            path = projects.get_project_field_value(id=id, field="path")
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


def cli():
    if args.web == "on":
        run_server(PORT)
    if args.command == "category":
        category_logics()
    elif args.command == "project":
        project_logics()

    if args.completed or args.favourite:
        dict = {"completed": args.completed, "favorite": args.favourite}
        query = {k: v for k, v in dict.items() if v == "True"}
        data = projects.get_projects_cf(query)
        print("ID || Name || Completed || Favorite || Description")
        for id, name, description, completed, favorite, *rest in data["data"]:
            fav = True if favorite == 1 else False
            com = True if completed == 1 else False
            print(f" {id} ||  {name} || {com} || {fav} || {description}")
