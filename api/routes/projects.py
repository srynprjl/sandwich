from flask import Blueprint, request

from lib import projects

project = Blueprint("project", __name__)


@project.get("/api/category/<int:categoryId>/project")
def get_all_project(categoryId: int):
    return projects.list_all_projects(category=categoryId)


@project.post("/api/category/<int:categoryId>/project")
def add_project(categoryId):
    body = request.get_json()

    body["category"] = categoryId
    response = projects.add_project(project_data=body)
    print(response)
    return response


@project.delete("/api/category/<int:categoryId>/project/<int:projectId>")
def delete_project(categoryId, projectId):
    return projects.delete_project(projectId, categoryId)


@project.patch("/api/category/<int:categoryId>/project/<int:projectId>")
def update_project(categoryId, projectId):
    body = request.get_json()
    return projects.update_project(projectId, body, categoryId)


@project.get("/api/category/<int:categoryId>/project/<int:projectId>")
def get_a_project(categoryId, projectId):
    return projects.get_project(projectId, categoryId)


@project.get("/api/project/")
def get_project_by_query():
    return projects.get_projects_cf(dict(request.args))
