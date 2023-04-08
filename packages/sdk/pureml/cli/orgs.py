import requests
import typer
from rich import print
from rich.console import Console
from rich.table import Table
from urllib.parse import urljoin
from pureml.schema import BackendSchema, PathSchema
from pureml.components import get_token



path_schema = PathSchema().get_instance()
backend_schema = BackendSchema().get_instance()
app = typer.Typer()


@app.callback()
def callback():
    """
    Organization command

    Use with show or select option

    show - lists all the organizations\n
    select - select an organization
    """

@app.command()
def show():
    access_token = get_token()
    url_path = "org"
    url = urljoin(backend_schema.BASE_URL, url_path)

    headers = {
        "accept": "application/json",
        "Authorization": "Bearer {}".format(access_token),
    }
    response = requests.get(url, headers=headers)
    if response.ok:
        print()
        print("[bold green] You are part of following Organizations!")
        org_all = response.json()["data"]
        console = Console()
        count = 0
        table = Table("Sr.No.", "Name", "Description")
        for org in org_all:
            count += 1
            table.add_row(str(count), org["org"]["name"], org["org"]["description"])

        console.print(table)
        print()
        return count
    else:
        print("[bold red]Unable to get the list of organizations!")
        return None

@app.command()
def select():
    print("[bold green] Select the Sr.No. of Organization")
    count = show()
    if count:
        sr_no = -1
        while int(sr_no) not in range(1, count + 1):
            sr_no: str = typer.prompt("Enter your Sr.No. of Organization (1 .... " + str(count) + ")")
            if int(sr_no) not in range(1, count + 1):
                print("[bold red]Invalid Sr.No. of Organization!")
        return sr_no
    else:
        print("[bold red]Did not Select any organization!")
        return None

# Possibly useful for future commands
# Moved from auth.py
def check_org_status(access_token: str, base_url: str):

    org_id: str = typer.prompt("Enter your Org Id")

    url_path = "org/id/{}".format(org_id)
    url = urljoin(base_url, url_path)

    headers = {
        "accept": "application/json",
        "Authorization": "Bearer {}".format(access_token),
    }

    response = requests.get(url, headers=headers)

    if response.ok:
        # print("[green]Organization Exists!")
        return org_id
    else:
        print("[red]Organization does not Exists!")
        return None