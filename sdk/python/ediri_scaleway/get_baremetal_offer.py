# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities
from . import outputs

__all__ = [
    'GetBaremetalOfferResult',
    'AwaitableGetBaremetalOfferResult',
    'get_baremetal_offer',
    'get_baremetal_offer_output',
]

@pulumi.output_type
class GetBaremetalOfferResult:
    """
    A collection of values returned by getBaremetalOffer.
    """
    def __init__(__self__, bandwidth=None, commercial_range=None, cpus=None, disks=None, id=None, include_disabled=None, memories=None, name=None, offer_id=None, stock=None, subscription_period=None, zone=None):
        if bandwidth and not isinstance(bandwidth, int):
            raise TypeError("Expected argument 'bandwidth' to be a int")
        pulumi.set(__self__, "bandwidth", bandwidth)
        if commercial_range and not isinstance(commercial_range, str):
            raise TypeError("Expected argument 'commercial_range' to be a str")
        pulumi.set(__self__, "commercial_range", commercial_range)
        if cpus and not isinstance(cpus, list):
            raise TypeError("Expected argument 'cpus' to be a list")
        pulumi.set(__self__, "cpus", cpus)
        if disks and not isinstance(disks, list):
            raise TypeError("Expected argument 'disks' to be a list")
        pulumi.set(__self__, "disks", disks)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if include_disabled and not isinstance(include_disabled, bool):
            raise TypeError("Expected argument 'include_disabled' to be a bool")
        pulumi.set(__self__, "include_disabled", include_disabled)
        if memories and not isinstance(memories, list):
            raise TypeError("Expected argument 'memories' to be a list")
        pulumi.set(__self__, "memories", memories)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if offer_id and not isinstance(offer_id, str):
            raise TypeError("Expected argument 'offer_id' to be a str")
        pulumi.set(__self__, "offer_id", offer_id)
        if stock and not isinstance(stock, str):
            raise TypeError("Expected argument 'stock' to be a str")
        pulumi.set(__self__, "stock", stock)
        if subscription_period and not isinstance(subscription_period, str):
            raise TypeError("Expected argument 'subscription_period' to be a str")
        pulumi.set(__self__, "subscription_period", subscription_period)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def bandwidth(self) -> int:
        """
        Available Bandwidth with the offer.
        """
        return pulumi.get(self, "bandwidth")

    @property
    @pulumi.getter(name="commercialRange")
    def commercial_range(self) -> str:
        """
        Commercial range of the offer.
        """
        return pulumi.get(self, "commercial_range")

    @property
    @pulumi.getter
    def cpus(self) -> Sequence['outputs.GetBaremetalOfferCpusResult']:
        """
        A list of cpu specifications. (Structure is documented below.)
        """
        return pulumi.get(self, "cpus")

    @property
    @pulumi.getter
    def disks(self) -> Sequence['outputs.GetBaremetalOfferDiskResult']:
        """
        A list of disk specifications. (Structure is documented below.)
        """
        return pulumi.get(self, "disks")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="includeDisabled")
    def include_disabled(self) -> Optional[bool]:
        return pulumi.get(self, "include_disabled")

    @property
    @pulumi.getter
    def memories(self) -> Sequence['outputs.GetBaremetalOfferMemoryResult']:
        """
        A list of memory specifications. (Structure is documented below.)
        """
        return pulumi.get(self, "memories")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Name of the CPU.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="offerId")
    def offer_id(self) -> Optional[str]:
        return pulumi.get(self, "offer_id")

    @property
    @pulumi.getter
    def stock(self) -> str:
        """
        Stock status for this offer. Possible values are: `empty`, `low` or `available`.
        """
        return pulumi.get(self, "stock")

    @property
    @pulumi.getter(name="subscriptionPeriod")
    def subscription_period(self) -> Optional[str]:
        return pulumi.get(self, "subscription_period")

    @property
    @pulumi.getter
    def zone(self) -> str:
        return pulumi.get(self, "zone")


class AwaitableGetBaremetalOfferResult(GetBaremetalOfferResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBaremetalOfferResult(
            bandwidth=self.bandwidth,
            commercial_range=self.commercial_range,
            cpus=self.cpus,
            disks=self.disks,
            id=self.id,
            include_disabled=self.include_disabled,
            memories=self.memories,
            name=self.name,
            offer_id=self.offer_id,
            stock=self.stock,
            subscription_period=self.subscription_period,
            zone=self.zone)


def get_baremetal_offer(include_disabled: Optional[bool] = None,
                        name: Optional[str] = None,
                        offer_id: Optional[str] = None,
                        subscription_period: Optional[str] = None,
                        zone: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBaremetalOfferResult:
    """
    Gets information about a baremetal offer. For more information, see [the documentation](https://developers.scaleway.com/en/products/baremetal/api).

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    my_offer = scaleway.get_baremetal_offer(offer_id="25dcf38b-c90c-4b18-97a2-6956e9d1e113",
        zone="fr-par-2")
    ```


    :param str name: The offer name. Only one of `name` and `offer_id` should be specified.
    :param str offer_id: The offer id. Only one of `name` and `offer_id` should be specified.
    :param str subscription_period: Period of subscription the desired offer. Should be `hourly` or `monthly`.
    :param str zone: `zone`) The zone in which the offer should be created.
    """
    __args__ = dict()
    __args__['includeDisabled'] = include_disabled
    __args__['name'] = name
    __args__['offerId'] = offer_id
    __args__['subscriptionPeriod'] = subscription_period
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:index/getBaremetalOffer:getBaremetalOffer', __args__, opts=opts, typ=GetBaremetalOfferResult).value

    return AwaitableGetBaremetalOfferResult(
        bandwidth=pulumi.get(__ret__, 'bandwidth'),
        commercial_range=pulumi.get(__ret__, 'commercial_range'),
        cpus=pulumi.get(__ret__, 'cpus'),
        disks=pulumi.get(__ret__, 'disks'),
        id=pulumi.get(__ret__, 'id'),
        include_disabled=pulumi.get(__ret__, 'include_disabled'),
        memories=pulumi.get(__ret__, 'memories'),
        name=pulumi.get(__ret__, 'name'),
        offer_id=pulumi.get(__ret__, 'offer_id'),
        stock=pulumi.get(__ret__, 'stock'),
        subscription_period=pulumi.get(__ret__, 'subscription_period'),
        zone=pulumi.get(__ret__, 'zone'))
def get_baremetal_offer_output(include_disabled: Optional[pulumi.Input[Optional[bool]]] = None,
                               name: Optional[pulumi.Input[Optional[str]]] = None,
                               offer_id: Optional[pulumi.Input[Optional[str]]] = None,
                               subscription_period: Optional[pulumi.Input[Optional[str]]] = None,
                               zone: Optional[pulumi.Input[Optional[str]]] = None,
                               opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetBaremetalOfferResult]:
    """
    Gets information about a baremetal offer. For more information, see [the documentation](https://developers.scaleway.com/en/products/baremetal/api).

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    my_offer = scaleway.get_baremetal_offer(offer_id="25dcf38b-c90c-4b18-97a2-6956e9d1e113",
        zone="fr-par-2")
    ```


    :param str name: The offer name. Only one of `name` and `offer_id` should be specified.
    :param str offer_id: The offer id. Only one of `name` and `offer_id` should be specified.
    :param str subscription_period: Period of subscription the desired offer. Should be `hourly` or `monthly`.
    :param str zone: `zone`) The zone in which the offer should be created.
    """
    __args__ = dict()
    __args__['includeDisabled'] = include_disabled
    __args__['name'] = name
    __args__['offerId'] = offer_id
    __args__['subscriptionPeriod'] = subscription_period
    __args__['zone'] = zone
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('scaleway:index/getBaremetalOffer:getBaremetalOffer', __args__, opts=opts, typ=GetBaremetalOfferResult)
    return __ret__.apply(lambda __response__: GetBaremetalOfferResult(
        bandwidth=pulumi.get(__response__, 'bandwidth'),
        commercial_range=pulumi.get(__response__, 'commercial_range'),
        cpus=pulumi.get(__response__, 'cpus'),
        disks=pulumi.get(__response__, 'disks'),
        id=pulumi.get(__response__, 'id'),
        include_disabled=pulumi.get(__response__, 'include_disabled'),
        memories=pulumi.get(__response__, 'memories'),
        name=pulumi.get(__response__, 'name'),
        offer_id=pulumi.get(__response__, 'offer_id'),
        stock=pulumi.get(__response__, 'stock'),
        subscription_period=pulumi.get(__response__, 'subscription_period'),
        zone=pulumi.get(__response__, 'zone')))
