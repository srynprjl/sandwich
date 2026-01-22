import json

from textual.app import App, ComposeResult
from textual.widgets import (
    Footer,
    Header,
    Label,
    ListItem,
    ListView,
    OptionList,
)
from textual.widgets.option_list import Option


class ProjectApp(App):
    BINDINGS = [
        ("v", "view_project", "View Project"),
        ("d", "delete_project", "Delete Project"),
        ("a", "add_project", "Add Project"),
        ("q", "quit", "Quit"),
    ]

    def compose(self) -> ComposeResult:
        yield Header()
        yield OptionList(id="my_list")
        yield Footer()

    def on_list_view_selected(self, event):
        self.notify(f"You clicked on: {event}")
        pass

    def on_mount(self) -> None:
        self.title = "Categories"

    def on_ready(self) -> None:
        self.load_data()

    def load_data(self) -> None:
        list_view = self.query_one("#my_list", OptionList)
        with open("plans.json", "r") as f:
            data = json.load(f)
        items = [(i["title"], i["shorthand"]) for i in data]
        for title, id in items:
            list_view.add_option(Option(title, id=id))

    def action_quit(self) -> None:
        exit()


def main():
    pass


if __name__ == "__main__":
    app = ProjectApp()
    app.run()
