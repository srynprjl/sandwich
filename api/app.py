from flask import Flask

from .routes import category, projects
app = Flask(__name__)
app.register_blueprint(category.categories)
app.register_blueprint(projects.project)


def run_server(port: int = 5000):
    app.run("0.0.0.0", port=port)
