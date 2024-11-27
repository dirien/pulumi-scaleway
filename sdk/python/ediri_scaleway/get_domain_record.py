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
    'GetDomainRecordResult',
    'AwaitableGetDomainRecordResult',
    'get_domain_record',
    'get_domain_record_output',
]

@pulumi.output_type
class GetDomainRecordResult:
    """
    A collection of values returned by getDomainRecord.
    """
    def __init__(__self__, data=None, dns_zone=None, fqdn=None, geo_ips=None, http_services=None, id=None, keep_empty_zone=None, name=None, priority=None, project_id=None, record_id=None, root_zone=None, ttl=None, type=None, views=None, weighteds=None):
        if data and not isinstance(data, str):
            raise TypeError("Expected argument 'data' to be a str")
        pulumi.set(__self__, "data", data)
        if dns_zone and not isinstance(dns_zone, str):
            raise TypeError("Expected argument 'dns_zone' to be a str")
        pulumi.set(__self__, "dns_zone", dns_zone)
        if fqdn and not isinstance(fqdn, str):
            raise TypeError("Expected argument 'fqdn' to be a str")
        pulumi.set(__self__, "fqdn", fqdn)
        if geo_ips and not isinstance(geo_ips, list):
            raise TypeError("Expected argument 'geo_ips' to be a list")
        pulumi.set(__self__, "geo_ips", geo_ips)
        if http_services and not isinstance(http_services, list):
            raise TypeError("Expected argument 'http_services' to be a list")
        pulumi.set(__self__, "http_services", http_services)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if keep_empty_zone and not isinstance(keep_empty_zone, bool):
            raise TypeError("Expected argument 'keep_empty_zone' to be a bool")
        pulumi.set(__self__, "keep_empty_zone", keep_empty_zone)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if priority and not isinstance(priority, int):
            raise TypeError("Expected argument 'priority' to be a int")
        pulumi.set(__self__, "priority", priority)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if record_id and not isinstance(record_id, str):
            raise TypeError("Expected argument 'record_id' to be a str")
        pulumi.set(__self__, "record_id", record_id)
        if root_zone and not isinstance(root_zone, bool):
            raise TypeError("Expected argument 'root_zone' to be a bool")
        pulumi.set(__self__, "root_zone", root_zone)
        if ttl and not isinstance(ttl, int):
            raise TypeError("Expected argument 'ttl' to be a int")
        pulumi.set(__self__, "ttl", ttl)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if views and not isinstance(views, list):
            raise TypeError("Expected argument 'views' to be a list")
        pulumi.set(__self__, "views", views)
        if weighteds and not isinstance(weighteds, list):
            raise TypeError("Expected argument 'weighteds' to be a list")
        pulumi.set(__self__, "weighteds", weighteds)

    @property
    @pulumi.getter
    def data(self) -> Optional[str]:
        return pulumi.get(self, "data")

    @property
    @pulumi.getter(name="dnsZone")
    def dns_zone(self) -> Optional[str]:
        return pulumi.get(self, "dns_zone")

    @property
    @pulumi.getter
    def fqdn(self) -> str:
        return pulumi.get(self, "fqdn")

    @property
    @pulumi.getter(name="geoIps")
    def geo_ips(self) -> Sequence['outputs.GetDomainRecordGeoIpResult']:
        """
        Information about dynamic records based on user geolocation. Find out more about dynamic records.
        """
        return pulumi.get(self, "geo_ips")

    @property
    @pulumi.getter(name="httpServices")
    def http_services(self) -> Sequence['outputs.GetDomainRecordHttpServiceResult']:
        """
        Information about dynamic records based on URL resolution. Find out more about dynamic records.
        """
        return pulumi.get(self, "http_services")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="keepEmptyZone")
    def keep_empty_zone(self) -> bool:
        return pulumi.get(self, "keep_empty_zone")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def priority(self) -> int:
        """
        The priority of the record, mainly used with `MX` records.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[str]:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="recordId")
    def record_id(self) -> Optional[str]:
        return pulumi.get(self, "record_id")

    @property
    @pulumi.getter(name="rootZone")
    def root_zone(self) -> bool:
        return pulumi.get(self, "root_zone")

    @property
    @pulumi.getter
    def ttl(self) -> int:
        """
        The Time To Live (TTL) of the record in seconds.
        """
        return pulumi.get(self, "ttl")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def views(self) -> Sequence['outputs.GetDomainRecordViewResult']:
        """
        Information about dynamic records based on the client’s (resolver) subnet. Find out more about dynamic records.
        """
        return pulumi.get(self, "views")

    @property
    @pulumi.getter
    def weighteds(self) -> Sequence['outputs.GetDomainRecordWeightedResult']:
        """
        Information about dynamic records based on IP weights. Find out more about dynamic records.
        """
        return pulumi.get(self, "weighteds")


class AwaitableGetDomainRecordResult(GetDomainRecordResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDomainRecordResult(
            data=self.data,
            dns_zone=self.dns_zone,
            fqdn=self.fqdn,
            geo_ips=self.geo_ips,
            http_services=self.http_services,
            id=self.id,
            keep_empty_zone=self.keep_empty_zone,
            name=self.name,
            priority=self.priority,
            project_id=self.project_id,
            record_id=self.record_id,
            root_zone=self.root_zone,
            ttl=self.ttl,
            type=self.type,
            views=self.views,
            weighteds=self.weighteds)


def get_domain_record(data: Optional[str] = None,
                      dns_zone: Optional[str] = None,
                      name: Optional[str] = None,
                      project_id: Optional[str] = None,
                      record_id: Optional[str] = None,
                      type: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDomainRecordResult:
    """
    The `DomainRecord` data source is used to get information about an existing domain record.

    Refer to the Domains and DNS [product documentation](https://www.scaleway.com/en/docs/network/domains-and-dns/) and [API documentation](https://www.scaleway.com/en/developers/api/domains-and-dns/) for more information.

    ## Query domain records

    The following commands allow you to:

    - query a domain record specified by the DNS zone (`domain.tld`), the record name (`www`), the record type (`A`), and the record content (`1.2.3.4`).
    - query a domain record specified by the DNS zone (`domain.tld`) and the unique record ID (`11111111-1111-1111-1111-111111111111`).

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    by_content = scaleway.get_domain_record(data="1.2.3.4",
        dns_zone="domain.tld",
        name="www",
        type="A")
    by_id = scaleway.get_domain_record(dns_zone="domain.tld",
        record_id="11111111-1111-1111-1111-111111111111")
    ```


    :param str data: The content of the record (e.g., an IPv4 address for an `A` record or a string for a `TXT` record). Cannot be used with `record_id`.
    :param str dns_zone: The DNS zone (domain) to which the record belongs. This is a required field in both examples above but is optional in the context of defining the data source.
    :param str name: The name of the record, which can be an empty string for a root record. Cannot be used with `record_id`.
    :param str project_id: ). The ID of the Project associated with the domain.
    :param str record_id: The unique identifier of the record. Cannot be used with `name`, `type`, and `data`.
    :param str type: The type of the record (`A`, `AAAA`, `MX`, `CNAME`, etc.). Cannot be used with `record_id`.
    """
    __args__ = dict()
    __args__['data'] = data
    __args__['dnsZone'] = dns_zone
    __args__['name'] = name
    __args__['projectId'] = project_id
    __args__['recordId'] = record_id
    __args__['type'] = type
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:index/getDomainRecord:getDomainRecord', __args__, opts=opts, typ=GetDomainRecordResult).value

    return AwaitableGetDomainRecordResult(
        data=pulumi.get(__ret__, 'data'),
        dns_zone=pulumi.get(__ret__, 'dns_zone'),
        fqdn=pulumi.get(__ret__, 'fqdn'),
        geo_ips=pulumi.get(__ret__, 'geo_ips'),
        http_services=pulumi.get(__ret__, 'http_services'),
        id=pulumi.get(__ret__, 'id'),
        keep_empty_zone=pulumi.get(__ret__, 'keep_empty_zone'),
        name=pulumi.get(__ret__, 'name'),
        priority=pulumi.get(__ret__, 'priority'),
        project_id=pulumi.get(__ret__, 'project_id'),
        record_id=pulumi.get(__ret__, 'record_id'),
        root_zone=pulumi.get(__ret__, 'root_zone'),
        ttl=pulumi.get(__ret__, 'ttl'),
        type=pulumi.get(__ret__, 'type'),
        views=pulumi.get(__ret__, 'views'),
        weighteds=pulumi.get(__ret__, 'weighteds'))
def get_domain_record_output(data: Optional[pulumi.Input[Optional[str]]] = None,
                             dns_zone: Optional[pulumi.Input[Optional[str]]] = None,
                             name: Optional[pulumi.Input[Optional[str]]] = None,
                             project_id: Optional[pulumi.Input[Optional[str]]] = None,
                             record_id: Optional[pulumi.Input[Optional[str]]] = None,
                             type: Optional[pulumi.Input[Optional[str]]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDomainRecordResult]:
    """
    The `DomainRecord` data source is used to get information about an existing domain record.

    Refer to the Domains and DNS [product documentation](https://www.scaleway.com/en/docs/network/domains-and-dns/) and [API documentation](https://www.scaleway.com/en/developers/api/domains-and-dns/) for more information.

    ## Query domain records

    The following commands allow you to:

    - query a domain record specified by the DNS zone (`domain.tld`), the record name (`www`), the record type (`A`), and the record content (`1.2.3.4`).
    - query a domain record specified by the DNS zone (`domain.tld`) and the unique record ID (`11111111-1111-1111-1111-111111111111`).

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    by_content = scaleway.get_domain_record(data="1.2.3.4",
        dns_zone="domain.tld",
        name="www",
        type="A")
    by_id = scaleway.get_domain_record(dns_zone="domain.tld",
        record_id="11111111-1111-1111-1111-111111111111")
    ```


    :param str data: The content of the record (e.g., an IPv4 address for an `A` record or a string for a `TXT` record). Cannot be used with `record_id`.
    :param str dns_zone: The DNS zone (domain) to which the record belongs. This is a required field in both examples above but is optional in the context of defining the data source.
    :param str name: The name of the record, which can be an empty string for a root record. Cannot be used with `record_id`.
    :param str project_id: ). The ID of the Project associated with the domain.
    :param str record_id: The unique identifier of the record. Cannot be used with `name`, `type`, and `data`.
    :param str type: The type of the record (`A`, `AAAA`, `MX`, `CNAME`, etc.). Cannot be used with `record_id`.
    """
    __args__ = dict()
    __args__['data'] = data
    __args__['dnsZone'] = dns_zone
    __args__['name'] = name
    __args__['projectId'] = project_id
    __args__['recordId'] = record_id
    __args__['type'] = type
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('scaleway:index/getDomainRecord:getDomainRecord', __args__, opts=opts, typ=GetDomainRecordResult)
    return __ret__.apply(lambda __response__: GetDomainRecordResult(
        data=pulumi.get(__response__, 'data'),
        dns_zone=pulumi.get(__response__, 'dns_zone'),
        fqdn=pulumi.get(__response__, 'fqdn'),
        geo_ips=pulumi.get(__response__, 'geo_ips'),
        http_services=pulumi.get(__response__, 'http_services'),
        id=pulumi.get(__response__, 'id'),
        keep_empty_zone=pulumi.get(__response__, 'keep_empty_zone'),
        name=pulumi.get(__response__, 'name'),
        priority=pulumi.get(__response__, 'priority'),
        project_id=pulumi.get(__response__, 'project_id'),
        record_id=pulumi.get(__response__, 'record_id'),
        root_zone=pulumi.get(__response__, 'root_zone'),
        ttl=pulumi.get(__response__, 'ttl'),
        type=pulumi.get(__response__, 'type'),
        views=pulumi.get(__response__, 'views'),
        weighteds=pulumi.get(__response__, 'weighteds')))
