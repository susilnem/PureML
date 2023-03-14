from pydantic import BaseModel, Field
import typing
from typing import List, Any
from abc import ABC, abstractmethod
from pureml.schema import InputSchema, OutputSchema
from enum import Enum


class BasePredictor(BaseModel, ABC):
    model_label: list
    model: Any = None
    input: InputSchema
    output: OutputSchema = OutputSchema()
    requirements_py: list = None
    requirements_sys: list = None

    class Config:
        arbitrary_types_allowed = True

    @abstractmethod
    def load_models(self):
        pass

    @abstractmethod
    def predict(self, **kwargs: typing.Any):
        pass

    # @abstractmethod
    def load_requirements_py(self):
        pass

    # @abstractmethod
    def load_requirements_sys(self):
        pass

    # @abstractmethod
    def load_resources(self):
        pass
