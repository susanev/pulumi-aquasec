# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetIntegrationRegistryResult',
    'AwaitableGetIntegrationRegistryResult',
    'get_integration_registry',
    'get_integration_registry_output',
]

@pulumi.output_type
class GetIntegrationRegistryResult:
    """
    A collection of values returned by getIntegrationRegistry.
    """
    def __init__(__self__, auto_cleanup=None, auto_pull=None, auto_pull_interval=None, auto_pull_max=None, auto_pull_rescan=None, auto_pull_time=None, id=None, image_creation_date_condition=None, name=None, password=None, prefixes=None, pull_image_age=None, pull_image_count=None, scanner_names=None, scanner_type=None, type=None, url=None, username=None):
        if auto_cleanup and not isinstance(auto_cleanup, bool):
            raise TypeError("Expected argument 'auto_cleanup' to be a bool")
        pulumi.set(__self__, "auto_cleanup", auto_cleanup)
        if auto_pull and not isinstance(auto_pull, bool):
            raise TypeError("Expected argument 'auto_pull' to be a bool")
        pulumi.set(__self__, "auto_pull", auto_pull)
        if auto_pull_interval and not isinstance(auto_pull_interval, int):
            raise TypeError("Expected argument 'auto_pull_interval' to be a int")
        pulumi.set(__self__, "auto_pull_interval", auto_pull_interval)
        if auto_pull_max and not isinstance(auto_pull_max, int):
            raise TypeError("Expected argument 'auto_pull_max' to be a int")
        pulumi.set(__self__, "auto_pull_max", auto_pull_max)
        if auto_pull_rescan and not isinstance(auto_pull_rescan, bool):
            raise TypeError("Expected argument 'auto_pull_rescan' to be a bool")
        pulumi.set(__self__, "auto_pull_rescan", auto_pull_rescan)
        if auto_pull_time and not isinstance(auto_pull_time, str):
            raise TypeError("Expected argument 'auto_pull_time' to be a str")
        pulumi.set(__self__, "auto_pull_time", auto_pull_time)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if image_creation_date_condition and not isinstance(image_creation_date_condition, str):
            raise TypeError("Expected argument 'image_creation_date_condition' to be a str")
        pulumi.set(__self__, "image_creation_date_condition", image_creation_date_condition)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if password and not isinstance(password, str):
            raise TypeError("Expected argument 'password' to be a str")
        pulumi.set(__self__, "password", password)
        if prefixes and not isinstance(prefixes, list):
            raise TypeError("Expected argument 'prefixes' to be a list")
        pulumi.set(__self__, "prefixes", prefixes)
        if pull_image_age and not isinstance(pull_image_age, str):
            raise TypeError("Expected argument 'pull_image_age' to be a str")
        pulumi.set(__self__, "pull_image_age", pull_image_age)
        if pull_image_count and not isinstance(pull_image_count, int):
            raise TypeError("Expected argument 'pull_image_count' to be a int")
        pulumi.set(__self__, "pull_image_count", pull_image_count)
        if scanner_names and not isinstance(scanner_names, list):
            raise TypeError("Expected argument 'scanner_names' to be a list")
        pulumi.set(__self__, "scanner_names", scanner_names)
        if scanner_type and not isinstance(scanner_type, str):
            raise TypeError("Expected argument 'scanner_type' to be a str")
        pulumi.set(__self__, "scanner_type", scanner_type)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if url and not isinstance(url, str):
            raise TypeError("Expected argument 'url' to be a str")
        pulumi.set(__self__, "url", url)
        if username and not isinstance(username, str):
            raise TypeError("Expected argument 'username' to be a str")
        pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter(name="autoCleanup")
    def auto_cleanup(self) -> bool:
        """
        Automatically clean up images and repositories which are no longer present in the registry from Aqua console
        """
        return pulumi.get(self, "auto_cleanup")

    @property
    @pulumi.getter(name="autoPull")
    def auto_pull(self) -> bool:
        """
        Whether to automatically pull images from the registry on creation and daily
        """
        return pulumi.get(self, "auto_pull")

    @property
    @pulumi.getter(name="autoPullInterval")
    def auto_pull_interval(self) -> int:
        """
        The interval in days to start pulling new images from the registry, Defaults to 1
        """
        return pulumi.get(self, "auto_pull_interval")

    @property
    @pulumi.getter(name="autoPullMax")
    def auto_pull_max(self) -> int:
        """
        Maximum number of repositories to pull every day, defaults to 100
        """
        return pulumi.get(self, "auto_pull_max")

    @property
    @pulumi.getter(name="autoPullRescan")
    def auto_pull_rescan(self) -> bool:
        """
        Whether to automatically pull and rescan images from the registry on creation and daily
        """
        return pulumi.get(self, "auto_pull_rescan")

    @property
    @pulumi.getter(name="autoPullTime")
    def auto_pull_time(self) -> str:
        """
        The time of day to start pulling new images from the registry, in the format HH:MM (24-hour clock), defaults to 03:00
        """
        return pulumi.get(self, "auto_pull_time")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="imageCreationDateCondition")
    def image_creation_date_condition(self) -> str:
        """
        Additional condition for pulling and rescanning images, Defaults to 'none'
        """
        return pulumi.get(self, "image_creation_date_condition")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the registry; string, required - this will be treated as the registry's ID, so choose a simple alphanumerical name without special signs and spaces
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def password(self) -> str:
        """
        The password for registry authentication
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def prefixes(self) -> Sequence[str]:
        """
        List of possible prefixes to image names pulled from the registry
        """
        return pulumi.get(self, "prefixes")

    @property
    @pulumi.getter(name="pullImageAge")
    def pull_image_age(self) -> str:
        """
        When auto pull image enabled, sets maximum age of auto pulled images
        """
        return pulumi.get(self, "pull_image_age")

    @property
    @pulumi.getter(name="pullImageCount")
    def pull_image_count(self) -> int:
        """
        When auto pull image enabled, sets maximum age of auto pulled images tags from each repository.
        """
        return pulumi.get(self, "pull_image_count")

    @property
    @pulumi.getter(name="scannerNames")
    def scanner_names(self) -> Sequence[str]:
        """
        List of scanner names
        """
        return pulumi.get(self, "scanner_names")

    @property
    @pulumi.getter(name="scannerType")
    def scanner_type(self) -> str:
        """
        Scanner type
        """
        return pulumi.get(self, "scanner_type")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Registry type (HUB / V1 / V2 / ENGINE / AWS / GCR).
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def url(self) -> str:
        """
        The URL, address or region of the registry
        """
        return pulumi.get(self, "url")

    @property
    @pulumi.getter
    def username(self) -> str:
        """
        The username for registry authentication.
        """
        return pulumi.get(self, "username")


class AwaitableGetIntegrationRegistryResult(GetIntegrationRegistryResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIntegrationRegistryResult(
            auto_cleanup=self.auto_cleanup,
            auto_pull=self.auto_pull,
            auto_pull_interval=self.auto_pull_interval,
            auto_pull_max=self.auto_pull_max,
            auto_pull_rescan=self.auto_pull_rescan,
            auto_pull_time=self.auto_pull_time,
            id=self.id,
            image_creation_date_condition=self.image_creation_date_condition,
            name=self.name,
            password=self.password,
            prefixes=self.prefixes,
            pull_image_age=self.pull_image_age,
            pull_image_count=self.pull_image_count,
            scanner_names=self.scanner_names,
            scanner_type=self.scanner_type,
            type=self.type,
            url=self.url,
            username=self.username)


def get_integration_registry(image_creation_date_condition: Optional[str] = None,
                             name: Optional[str] = None,
                             pull_image_age: Optional[str] = None,
                             pull_image_count: Optional[int] = None,
                             scanner_names: Optional[Sequence[str]] = None,
                             scanner_type: Optional[str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIntegrationRegistryResult:
    """
    Use this data source to access information about an existing resource.

    :param str image_creation_date_condition: Additional condition for pulling and rescanning images, Defaults to 'none'
    :param str name: The name of the registry; string, required - this will be treated as the registry's ID, so choose a simple alphanumerical name without special signs and spaces
    :param str pull_image_age: When auto pull image enabled, sets maximum age of auto pulled images
    :param int pull_image_count: When auto pull image enabled, sets maximum age of auto pulled images tags from each repository.
    :param Sequence[str] scanner_names: List of scanner names
    :param str scanner_type: Scanner type
    """
    __args__ = dict()
    __args__['imageCreationDateCondition'] = image_creation_date_condition
    __args__['name'] = name
    __args__['pullImageAge'] = pull_image_age
    __args__['pullImageCount'] = pull_image_count
    __args__['scannerNames'] = scanner_names
    __args__['scannerType'] = scanner_type
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aquasec:index/getIntegrationRegistry:getIntegrationRegistry', __args__, opts=opts, typ=GetIntegrationRegistryResult).value

    return AwaitableGetIntegrationRegistryResult(
        auto_cleanup=__ret__.auto_cleanup,
        auto_pull=__ret__.auto_pull,
        auto_pull_interval=__ret__.auto_pull_interval,
        auto_pull_max=__ret__.auto_pull_max,
        auto_pull_rescan=__ret__.auto_pull_rescan,
        auto_pull_time=__ret__.auto_pull_time,
        id=__ret__.id,
        image_creation_date_condition=__ret__.image_creation_date_condition,
        name=__ret__.name,
        password=__ret__.password,
        prefixes=__ret__.prefixes,
        pull_image_age=__ret__.pull_image_age,
        pull_image_count=__ret__.pull_image_count,
        scanner_names=__ret__.scanner_names,
        scanner_type=__ret__.scanner_type,
        type=__ret__.type,
        url=__ret__.url,
        username=__ret__.username)


@_utilities.lift_output_func(get_integration_registry)
def get_integration_registry_output(image_creation_date_condition: Optional[pulumi.Input[Optional[str]]] = None,
                                    name: Optional[pulumi.Input[str]] = None,
                                    pull_image_age: Optional[pulumi.Input[Optional[str]]] = None,
                                    pull_image_count: Optional[pulumi.Input[Optional[int]]] = None,
                                    scanner_names: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                    scanner_type: Optional[pulumi.Input[Optional[str]]] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetIntegrationRegistryResult]:
    """
    Use this data source to access information about an existing resource.

    :param str image_creation_date_condition: Additional condition for pulling and rescanning images, Defaults to 'none'
    :param str name: The name of the registry; string, required - this will be treated as the registry's ID, so choose a simple alphanumerical name without special signs and spaces
    :param str pull_image_age: When auto pull image enabled, sets maximum age of auto pulled images
    :param int pull_image_count: When auto pull image enabled, sets maximum age of auto pulled images tags from each repository.
    :param Sequence[str] scanner_names: List of scanner names
    :param str scanner_type: Scanner type
    """
    ...
