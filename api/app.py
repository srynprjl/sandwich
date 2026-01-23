from flask import Flask
from flask_cors import CORS

from .routes import category, projects
app = Flask(__name__)
CORS(app, origins=["http://127.0.0.1:5500"])
app.register_blueprint(category.categories)
app.register_blueprint(projects.project)


def run_server(port: int = 5000):
    app.run("0.0.0.0", port=port)
