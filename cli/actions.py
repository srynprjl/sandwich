import argparse
import sqlite3
import os
import subprocess
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
                print("Atleast one argument [name or shorthand] is required")
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
                print(f"ID || Name")
                for id, name, shorthand in categories:
                    print(f" {id} ||  {name}")
            else:
                project_list = projects.list_all_projects(con=con, category=argument)
                print(f"ID || Name")
                for id, name, *rest in project_list:
                    print(f" {id} ||  {name}")


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
            pass

        case "view":
            id = args.id
            project = projects.get_project(con=con, id=id)
            fav = True if project[4] == 1 else False
            com = True if project[3] == 1 else False
            print(project)
            print(f"ID -> {project[0]} ")
            print(f"Name -> {project[1]}")
            print(f"Description -> {project[2]}")
            print(f"Path -> {project[5]}")
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

    match (args.editor):

        case "code":
            try:
                id = args.id
                project = projects.get_project(con=con, id=id)
                path = project[5]
                code_loc = subprocess.run(
                    ["which", "code"], capture_output=True, text=True
                )
                subprocess.run([code_loc.stdout.strip(), path])
            except subprocess.CalledProcessError:
                print("Code not found.")
            pass
        case "zed":
            try:
                id = args.id
                project = projects.get_project(con=con, id=id)
                path = project[5]
                code_loc = subprocess.run(
                    ["which", "zeditor"], capture_output=True, text=True
                )
                subprocess.run([code_loc.stdout.strip(), path])
            except subprocess.CalledProcessError:
                print("Zed Editor not found.")
            pass
        case "nvim":
            try:
                id = args.id
                project = projects.get_project(con=con, id=id)
                path = project[5]
                code_loc = subprocess.run(
                    ["which", "nvim"], capture_output=True, text=True
                )
                subprocess.run([code_loc.stdout.strip(), path])
            except subprocess.CalledProcessError:
                print("Neovim not found.")
        case "default":
            try:
                id = args.id
                project = projects.get_project(con=con, id=id)
                path = project[5]
                code_loc = os.environ.get("EDITOR")
                subprocess.run([code_loc.strip(), path])
            except subprocess.CalledProcessError:
                print("Default not found.")
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
        pass

    if args.favourite == "all":
        fav = projects.get_fav_projects(con=con)
        print(f"ID || Name")
        for id, name, *rest in fav:
            print(f" {id} ||  {name}")
        pass
