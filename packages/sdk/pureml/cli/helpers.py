import json
import os

import typer
from pureml.components import get_token, get_api_token
from pureml.schema import PathSchema
from rich import print

path_schema = PathSchema().get_instance()

def save_auth(org_id: str = None, access_token: str = None, email: str = None, api_id: str = None, api_key: str = None):
    token_path = path_schema.PATH_USER_TOKEN

    token_dir = os.path.dirname(token_path)
    os.makedirs(token_dir, exist_ok=True)

    # Read existing token
    if os.path.exists(token_path):
        with open(token_path, "r") as token_file:
            token = json.load(token_file)

        if org_id is not None:
            token["org_id"] = org_id
        if access_token is not None:
            token["accessToken"] = access_token
        if api_id is not None:
            token["api_id"] = api_id
        if api_key is not None:
            token["api_key"] = api_key
        if email is not None:
            if "email" in token and token["email"] != email:
                token["org_id"] = ""
            token["email"] = email
    else:
        token = {"org_id": org_id, "accessToken": access_token, "email": email, "api_id": api_id, "api_key": api_key}
        if org_id is None:
            token["org_id"] = ""

    token = json.dumps(token)

    with open(token_path, "w") as token_file:
        token_file.write(token)

def get_auth_headers(content_type: str = "application/x-www-form-urlencoded"):
    token = get_token()
    api_token = get_api_token()
    if token is None and api_token is None:
        print(f"[bold red]Authentication token or API token does not exist! Please login")
        typer.Exit()
        return None
    elif token is not None:
        return {
            "Content-Type": content_type,
            "Accept": "application/json",
            "Authorization": "Bearer {}".format(token),
        }
    else:
        return {
            "Content-Type": content_type,
            "Accept": "application/json",
            "X-Api-Id": api_token["api_id"],
            "X-Api-Key": api_token["api_key"],
        }