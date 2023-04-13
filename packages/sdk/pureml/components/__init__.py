import os, json
from pureml.schema import PathSchema

path_schema = PathSchema().get_instance()


def get_token():
    """It checks if the token exists in the user's home directory. If it does, it returns the token. If it
    doesn't, it returns None

    Returns
    -------
        The token is being returned.
    """
    path = path_schema.PATH_USER_TOKEN
    # path = os.path.expanduser(path)

    if os.path.exists(path):
        creds = open(path, "r").read()

        creds_json = json.loads(creds)
        token = creds_json["accessToken"]
        # print(f"[bold green]Authentication token exists!")

        # print(token)
        return token
    else:
        return


def get_api_token():
    """It checks if the api token exists in the user's home directory. If it does, it returns the api token id and key. If it
    doesn't, it returns None

    Returns
    -------
        The api token id and key is returned.
        Dictionary format: {"api_id": api_id, "api_key": api_key}
    """
    path = path_schema.PATH_USER_TOKEN
    # path = os.path.expanduser(path)

    if os.path.exists(path):
        creds = open(path, "r").read()

        creds_json = json.loads(creds)
        api_id = creds_json["api_id"]
        api_key = creds_json["api_key"]
        # print(f"[bold green]Authentication token exists!")

        # print(token)
        return {
            "api_id": api_id,
            "api_key": api_key
        }
    else:
        return


def delete_token():
    """It checks if the token exists in the user's home directory. If it does, it deletes the token. If it
    doesn't, it returns None

    Returns
    -------
        The token is being returned.
    """
    path = path_schema.PATH_USER_TOKEN
    # path = os.path.expanduser(path)

    if os.path.exists(path):
        os.remove(path)
        print(f"[bold green]Authentication token deleted!")

        # print(token)
        return
    else:
        print(f"[bold red]Authentication token does not exist! Please login")

        return


def get_org_id():
    """It checks if the org exists in the user's home directory. If it does, it returns the org. If it
    doesn't, it returns None

    Returns
    -------
        The org is being returned.

    """

    path = path_schema.PATH_USER_TOKEN

    path = os.path.expanduser(path)

    if os.path.exists(path):
        creds = open(path, "r").read()

        creds_json = json.loads(creds)

        org_id = creds_json["org_id"]
        # print(f"[bold green]Organization exists!")

        # print(org_id)
        return org_id
    else:
        print(f"[bold red]Organization id does not exist! Please login")

        return


def convert_values_to_string(logged_dict):

    for key in logged_dict:
        logged_dict[key] = str(logged_dict[key])

    return logged_dict
