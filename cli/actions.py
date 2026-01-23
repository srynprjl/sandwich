import argparse
import sqlite3
import os
import subprocess
import shutil
from .parser import parse_arguments
from controllers import category, projects
args: argparse.Namespace = parse_arguments()

def category_logics(con: sqlite3.Connection):
    match (args.action):
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
                print(f"ID || Name")
                for id, name, shorthand in data:
                    print(f" {id} ||  {name}")
            else:
                project_list = projects.list_all_projects(con=con, category=argument)
                data = project_list["data"]
                print(f"ID || Name || Completed || Favorite || Description")
                for id, name, description, completed, favorite, *rest in data:
                    fav = True if favorite == 1 else False
                    com = True if completed == 1 else False
                    print(f" {id} ||  {name} || {com} || {fav} || {description}")


def project_logics(con: sqlite3.Connection):
    match (args.action):
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
            projects.delete_project(con=con, id=id)

        case "update":
            data = {
                "name": args.name,
                "description": args.description,
                "category": args.category,
                "path": args.path,
                "favorite": args.favourite,
                "completed": args.complete,
            }
            data = {k: v for k, v in data.items() if v is not None}
            print(data)
            projects.update_project(con, id=args.id, new_data=data)

        case "view":
            id = args.id
            project = projects.get_project(con=con, id=id)
            fav = True if project["data"][4] == 1 else False
            com = True if project["data"][3] == 1 else False
            print(f"ID -> {project["data"][0]} ")
            print(f"Name -> {project["data"][1]}")
            print(f"Description -> {project["data"][2]}")
            print(f"Path -> {project["data"][5]}")
            print(f"Favorite -> {fav}")
            print(f"Complete -> {com}")

        case "favourite":
            id = args.id
            project = projects.get_project(con=con, id=id)
            fav = 0 if project[4] == 1 else 1
            projects.update_project(con=con, id=id, new_data={"favorite": fav})
            pass

        case "complete":
            id = args.id
            project = projects.get_project(con=con, id=id)
            com = 0 if project[3] == 1 else 1
            projects.update_project(con=con, id=id, new_data={"completed": com})
            pass

    try:

        if args.editor:
            id = args.id
            path = projects.get_project_field_value(con=con, id=id, field="path")
            match (args.editor):
                case "code":
                   executable = shutil.which("code")
                case "zed":
                    executable = shutil.which("zeditor")
                case "nvim":
                    executable = shutil.which("nvim")
                case "default":
                    executable = os.environ.get("EDITOR", "vi")
        if executable:
            subprocess.run([executable, path])
        else:
            print(f"Editor '{executable}' not found.")
    except AttributeError:
        pass


def cli(con: sqlite3.Connection):
    if args.command == "category":
        category_logics(con)
    elif args.command == "project":
        project_logics(con)

    if args.completed == "all":
        coms = projects.get_completed_projects(con=con)
        print(f"ID || Name")
        for id, name, *rest in coms:
            print(f" {id} ||  {name}")

    if args.favourite == "all":
        fav = projects.get_fav_projects(con=con)
        print(f"ID || Name")
        for id, name, *rest in fav:
            print(f" {id} ||  {name}")
