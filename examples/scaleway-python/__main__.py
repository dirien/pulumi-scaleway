"""A Python Pulumi program"""

import pulumi
from ediri_scaleway import RegistryNamespace, RegistryNamespaceArgs

registry = RegistryNamespace("registry", RegistryNamespaceArgs(
    name="pulumi-scaleway-dirien",
    is_public=True,
))