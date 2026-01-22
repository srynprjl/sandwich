import argparse

def parse_arguments():
    # Main
    parser = argparse.ArgumentParser(prog="sandwich", description="Manage your projects easily in the terminal")

    # Help & Version
    parser.add_argument("--version", action="version", version="%(prog)s 1.0")
    subparsers = parser.add_subparsers(dest="command", help="Available commands")


    # Category
    cat_parser = subparsers.add_parser("category", help="Manage categories")
    cat_sub = cat_parser.add_subparsers(dest="action")
    list_cat = cat_sub.add_parser("list", help="List categories")
    list_cat.add_argument("id", nargs="?", type=int, help="List all project in the category")

    # category add
    add_cat = cat_sub.add_parser("add", help="Add a category")
    add_cat.add_argument("--name", required=True)
    add_cat.add_argument("--shorthand", required=True)

    update_cat = cat_sub.add_parser("update", help="Update a category")
    update_cat.add_argument("id")
    update_cat.add_argument("--name",  nargs="?")
    update_cat.add_argument("--shorthand",  nargs="?")

    delete_cat = cat_sub.add_parser("delete", help="Delete a category")
    delete_cat.add_argument("id", type=int)


    ## Project
    proj_parser = subparsers.add_parser("project", help="Manage projects")
    proj_sub = proj_parser.add_subparsers(dest="action")

    view_proj = proj_sub.add_parser("view", help="View project details")
    view_proj.add_argument("id", type=int)

    # project add
    add_proj = proj_sub.add_parser("add", help="Add a project")
    add_proj.add_argument("--name", required=True)
    add_proj.add_argument("--description")
    add_proj.add_argument("--category", type=int, required=True)
    add_proj.add_argument("--path", required=True)
    add_proj.add_argument("--favourite", action="store_true")
    add_proj.add_argument("--complete", action="store_true")

    update_proj = proj_sub.add_parser("update", help="Update a project")
    update_proj.add_argument("id")
    update_proj.add_argument("--name")
    update_proj.add_argument("--description")
    update_proj.add_argument("--category", type=int)
    update_proj.add_argument("--path")
    update_proj.add_argument("--favourite", action="store_true")
    update_proj.add_argument("--complete", action="store_true")

    delete_proj = proj_sub.add_parser("delete", help="Delete a project")
    delete_proj.add_argument("id", type=int)

    fav_proj = proj_sub.add_parser("favourite", help="Mark project as favorite")
    fav_proj.add_argument("id", type=int)

    comp_proj = proj_sub.add_parser("complete", help="Mark project as complete")
    comp_proj.add_argument("id", type=int)

    edit_proj = proj_sub.add_parser("edit", help="Open in editor")
    edit_proj.add_argument("--editor", choices=["code", "nvim", "zed", "default"],default="default", help="Open in editor")
    edit_proj.add_argument("id", type=int)

    parser.add_argument("--completed", nargs="?", const="all", help="List completed projects")
    parser.add_argument("--favourite", nargs="?", const="all", help="List favourite projects")



    args = parser.parse_args()

    return args
