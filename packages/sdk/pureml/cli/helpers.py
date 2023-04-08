from pathlib import Path
from pureml.cli.puremlconfig import PureMLConfigYML
from pureml.schema.backend import BackendSchema
from pureml.schema.paths import PathSchema


path_schema = PathSchema().get_instance()
backend_schema = BackendSchema().get_instance()
project_path = Path.cwd()
if Path.exists(project_path / "puremlconfig.yml"):
    puremlconfig = PureMLConfigYML(project_path / "puremlconfig.yml")
else:
    puremlconfig = None


def get_backend_base_url(backend_url: str = None):
    if backend_url is not None and backend_url != "":
        # override backend url (command specific option)
        return backend_url
    if puremlconfig is not None:
        # user configured backend url (self-hosted or custom pureml backend instance)
        backend_url = puremlconfig.data["backend_url"] if "backend_url" in puremlconfig.data else backend_schema.BACKEND_URL
    else:
        # default backend url (production cloud backend)
        backend_url = backend_schema.BACKEND_URL
    return backend_url

def get_frontend_base_url(frontend_url: str = None):
    if frontend_url is not None and frontend_url != "":
        # override frontend url (command specific option)
        return frontend_url
    if puremlconfig is not None:
        # user configured frontend url (self-hosted or custom pureml frontend instance)
        frontend_url = puremlconfig.data["frontend_url"] if "frontend_url" in puremlconfig.data else backend_schema.FRONTEND_URL
    else:
        # default frontend url (production cloud frontend)
        frontend_url = backend_schema.FRONTEND_URL
    return frontend_url