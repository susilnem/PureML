from pathlib import Path
from typing import Optional
import jwt
import requests

# import typer
from rich import print
from rich.syntax import Syntax

import os
import json
import typing

from urllib.parse import urljoin

from pureml.cli.helpers import get_auth_headers

from . import get_org_id

from pureml.schema import PathSchema, BackendSchema, ContentTypeHeader
from joblib import Parallel, delayed

path_schema = PathSchema().get_instance()
backend_schema = BackendSchema().get_instance()


def details(
    model_name: str, model_branch: str, model_version: str = "latest", name: str = ""
):
    """This function returns the details of the artifact for a given model

    Parameters
    ----------
    model_name : str
        The name of the model you want to get the artifact details for
    model_version: str
        The version of the model
    name : str
        The name of the artifact.

    """

    org_id = get_org_id()

    url = "org/{}/model/{}/branch/{}/version/{}/log".format(
        org_id, model_name, model_branch, model_version
    )
    url = urljoin(backend_schema.BASE_URL, url)

    headers = get_auth_headers(content_type=ContentTypeHeader.APP_FORM_URL_ENCODED)

    response = requests.get(url, headers=headers)

    if response.status_code == 200:
        res_text = json.loads(response.text)

        if len(res_text) == 0:
            print("[bold yellow] No Artifacts have been found for the model")
            print(res_text)
            return
        else:
            print("[bold green]Artifacts have been found for the model")
            print(res_text)
            return res_text

    else:
        print("[bold red]Unable to obtain the artifact details")
        print(response.text)
        return


# @app.command()
def add(
    artifact: str,
    name: str,
    model_name: str,
    model_branch: str,
    model_version: str = "latest",
) -> str:
    """`add` function takes in the path of the artifact, name of the artifact and the model name and
    registers the artifact

    Parameters
    ----------
    artifact : str
        The path to the artifact file.
    name : str
        The name of the artifact.
    model_name : str
        The name of the model you want to add artifacts to.
    model_version: str
        The version of the model

    Returns
    -------
        The response is a JSON object

    """

    org_id = get_org_id()

    url = "org/{}/model/{}/branch/{}/version/{}/log".format(
        org_id, model_name, model_branch, model_version
    )
    url = urljoin(backend_schema.BASE_URL, url)

    headers = get_auth_headers(content_type=ContentTypeHeader.APP_FORM_URL_ENCODED)

    if os.path.isfile(artifact):
        file = {"file": (name, open(artifact, "rb"))}
    else:
        print("[bold red] Artifact does not exist at the given path")

    data = {"name": name, "path": artifact}

    response = requests.post(url, data=data, files=file, headers=headers)

    if response.status_code == 200:
        print(f"[bold green]Artifacts have been registered!")

    else:
        print(f"[bold red]Artifacts have not been registered!")
        print(response.text)

    return response.text


# @app.command()
def fetch(model_name: str, model_version: str = "latest", name: str = ""):
    """It fetches the artifact from the server and stores it in the local directory

    Parameters
    ----------
    model_name : str
        The name of the model you want to fetch the artifact from.
    model_version: str
        The version of the model
    name : str
        The name of the artifact to be fetched. If not specified, all artifacts will be fetched.

    Returns
    -------
        The response text is being returned.

    """

    org_id = get_org_id()

    def fetch_artifact(artifact_details: dict):

        url = artifact_details["location"]
        file_path_temp = artifact_details["path"]
        file_name = file_path_temp.split(os.path.sep)[-1]
        save_path = os.path.join(path_schema.PATH_ARTIFACT_DIR, file_name)
        print("save path", save_path)

        name_fetched = artifact_details["artifact"]

        headers = get_auth_headers(content_type=ContentTypeHeader.APP_FORM_URL_ENCODED)

        print("Artifact url", url)

        # response = requests.get(url, headers=headers)
        response = requests.get(url)

        print(response.status_code)

        if response.status_code == 200:
            print("[bold green] Artifact {} has been fetched".format(name_fetched))

            save_dir = os.path.dirname(save_path)

            os.makedirs(save_dir, exist_ok=True)

            artifact_bytes = response.content

            open(save_path, "wb").write(artifact_bytes)

            print(
                "[bold green] Artifact {} has been stored at {}".format(
                    name_fetched, save_path
                )
            )

            return response.text
        else:
            print("[bold red] Unable to fetch the artifact")

            return response.text

    artifact_details = details(
        model_name=model_name, name=name, model_version=model_version
    )

    if artifact_details is None:
        return

    if type(artifact_details) is dict:

        res_text = fetch_artifact(artifact_details)

    elif type(artifact_details) is list:
        res_text = Parallel(n_jobs=-1)(
            delayed(fetch_artifact)(art_det) for art_det in artifact_details
        )

    return res_text


# @app.command()
def delete(
    name: str, model_name: str, model_branch: str, model_version: str = "latest"
) -> str:
    """`delete()` deletes an artifact from a model

    Parameters
    ----------
    name : str
        The name of the artifact you want to delete.
    model_name : str
        The name of the model you want to delete the artifact from
    model_version: str
        The version of the model

    """

    org_id = get_org_id()

    url = "org/{}/model/{}/branch/{}/version/{}/log".format(
        org_id, model_name, model_branch, model_version
    )
    url = urljoin(backend_schema.BASE_URL, url)

    headers = get_auth_headers(content_type=ContentTypeHeader.APP_FORM_URL_ENCODED)

    # artifact_details = details(model_name=model_name, artifact=artifact)

    # if artifact_details is None:
    #     print('[bold red] Unable to find artifact details')
    #     return

    response = requests.delete(url, headers=headers)

    if response.status_code == 200:
        print(f"[bold green]Artifact has been deleted")

    else:
        print(f"[bold red]Unable to delete artifact")

    return response.text


# if __name__ == "__main__":
#     app()
