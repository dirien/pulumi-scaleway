# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['CockpitAlertManagerArgs', 'CockpitAlertManager']

@pulumi.input_type
class CockpitAlertManagerArgs:
    def __init__(__self__, *,
                 contact_points: Optional[pulumi.Input[Sequence[pulumi.Input['CockpitAlertManagerContactPointArgs']]]] = None,
                 enable_managed_alerts: Optional[pulumi.Input[bool]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a CockpitAlertManager resource.
        :param pulumi.Input[Sequence[pulumi.Input['CockpitAlertManagerContactPointArgs']]] contact_points: A list of contact points with email addresses for the alert receivers. Each map should contain a single key email.
        :param pulumi.Input[bool] enable_managed_alerts: Indicates whether the alert manager should be enabled. Defaults to true.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the cockpit is associated with.
        :param pulumi.Input[str] region: `region`) The region in which alert_manager should be created.
        """
        if contact_points is not None:
            pulumi.set(__self__, "contact_points", contact_points)
        if enable_managed_alerts is not None:
            pulumi.set(__self__, "enable_managed_alerts", enable_managed_alerts)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="contactPoints")
    def contact_points(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CockpitAlertManagerContactPointArgs']]]]:
        """
        A list of contact points with email addresses for the alert receivers. Each map should contain a single key email.
        """
        return pulumi.get(self, "contact_points")

    @contact_points.setter
    def contact_points(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CockpitAlertManagerContactPointArgs']]]]):
        pulumi.set(self, "contact_points", value)

    @property
    @pulumi.getter(name="enableManagedAlerts")
    def enable_managed_alerts(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether the alert manager should be enabled. Defaults to true.
        """
        return pulumi.get(self, "enable_managed_alerts")

    @enable_managed_alerts.setter
    def enable_managed_alerts(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_managed_alerts", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        `project_id`) The ID of the project the cockpit is associated with.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`) The region in which alert_manager should be created.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _CockpitAlertManagerState:
    def __init__(__self__, *,
                 alert_manager_url: Optional[pulumi.Input[str]] = None,
                 contact_points: Optional[pulumi.Input[Sequence[pulumi.Input['CockpitAlertManagerContactPointArgs']]]] = None,
                 enable_managed_alerts: Optional[pulumi.Input[bool]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering CockpitAlertManager resources.
        :param pulumi.Input[str] alert_manager_url: Alert manager URL.
        :param pulumi.Input[Sequence[pulumi.Input['CockpitAlertManagerContactPointArgs']]] contact_points: A list of contact points with email addresses for the alert receivers. Each map should contain a single key email.
        :param pulumi.Input[bool] enable_managed_alerts: Indicates whether the alert manager should be enabled. Defaults to true.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the cockpit is associated with.
        :param pulumi.Input[str] region: `region`) The region in which alert_manager should be created.
        """
        if alert_manager_url is not None:
            pulumi.set(__self__, "alert_manager_url", alert_manager_url)
        if contact_points is not None:
            pulumi.set(__self__, "contact_points", contact_points)
        if enable_managed_alerts is not None:
            pulumi.set(__self__, "enable_managed_alerts", enable_managed_alerts)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="alertManagerUrl")
    def alert_manager_url(self) -> Optional[pulumi.Input[str]]:
        """
        Alert manager URL.
        """
        return pulumi.get(self, "alert_manager_url")

    @alert_manager_url.setter
    def alert_manager_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alert_manager_url", value)

    @property
    @pulumi.getter(name="contactPoints")
    def contact_points(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CockpitAlertManagerContactPointArgs']]]]:
        """
        A list of contact points with email addresses for the alert receivers. Each map should contain a single key email.
        """
        return pulumi.get(self, "contact_points")

    @contact_points.setter
    def contact_points(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CockpitAlertManagerContactPointArgs']]]]):
        pulumi.set(self, "contact_points", value)

    @property
    @pulumi.getter(name="enableManagedAlerts")
    def enable_managed_alerts(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether the alert manager should be enabled. Defaults to true.
        """
        return pulumi.get(self, "enable_managed_alerts")

    @enable_managed_alerts.setter
    def enable_managed_alerts(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_managed_alerts", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        `project_id`) The ID of the project the cockpit is associated with.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`) The region in which alert_manager should be created.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


class CockpitAlertManager(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 contact_points: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CockpitAlertManagerContactPointArgs']]]]] = None,
                 enable_managed_alerts: Optional[pulumi.Input[bool]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates and manages Scaleway Cockpit Alert Managers.

        For more information consult the [documentation](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#grafana-users).

        ## Import

        Alert managers can be imported using the project ID, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/cockpitAlertManager:CockpitAlertManager main fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CockpitAlertManagerContactPointArgs']]]] contact_points: A list of contact points with email addresses for the alert receivers. Each map should contain a single key email.
        :param pulumi.Input[bool] enable_managed_alerts: Indicates whether the alert manager should be enabled. Defaults to true.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the cockpit is associated with.
        :param pulumi.Input[str] region: `region`) The region in which alert_manager should be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[CockpitAlertManagerArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages Scaleway Cockpit Alert Managers.

        For more information consult the [documentation](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#grafana-users).

        ## Import

        Alert managers can be imported using the project ID, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/cockpitAlertManager:CockpitAlertManager main fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param CockpitAlertManagerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CockpitAlertManagerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 contact_points: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CockpitAlertManagerContactPointArgs']]]]] = None,
                 enable_managed_alerts: Optional[pulumi.Input[bool]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CockpitAlertManagerArgs.__new__(CockpitAlertManagerArgs)

            __props__.__dict__["contact_points"] = contact_points
            __props__.__dict__["enable_managed_alerts"] = enable_managed_alerts
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["region"] = region
            __props__.__dict__["alert_manager_url"] = None
        super(CockpitAlertManager, __self__).__init__(
            'scaleway:index/cockpitAlertManager:CockpitAlertManager',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            alert_manager_url: Optional[pulumi.Input[str]] = None,
            contact_points: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CockpitAlertManagerContactPointArgs']]]]] = None,
            enable_managed_alerts: Optional[pulumi.Input[bool]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None) -> 'CockpitAlertManager':
        """
        Get an existing CockpitAlertManager resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alert_manager_url: Alert manager URL.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CockpitAlertManagerContactPointArgs']]]] contact_points: A list of contact points with email addresses for the alert receivers. Each map should contain a single key email.
        :param pulumi.Input[bool] enable_managed_alerts: Indicates whether the alert manager should be enabled. Defaults to true.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the cockpit is associated with.
        :param pulumi.Input[str] region: `region`) The region in which alert_manager should be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CockpitAlertManagerState.__new__(_CockpitAlertManagerState)

        __props__.__dict__["alert_manager_url"] = alert_manager_url
        __props__.__dict__["contact_points"] = contact_points
        __props__.__dict__["enable_managed_alerts"] = enable_managed_alerts
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["region"] = region
        return CockpitAlertManager(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="alertManagerUrl")
    def alert_manager_url(self) -> pulumi.Output[str]:
        """
        Alert manager URL.
        """
        return pulumi.get(self, "alert_manager_url")

    @property
    @pulumi.getter(name="contactPoints")
    def contact_points(self) -> pulumi.Output[Optional[Sequence['outputs.CockpitAlertManagerContactPoint']]]:
        """
        A list of contact points with email addresses for the alert receivers. Each map should contain a single key email.
        """
        return pulumi.get(self, "contact_points")

    @property
    @pulumi.getter(name="enableManagedAlerts")
    def enable_managed_alerts(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates whether the alert manager should be enabled. Defaults to true.
        """
        return pulumi.get(self, "enable_managed_alerts")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        `project_id`) The ID of the project the cockpit is associated with.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        `region`) The region in which alert_manager should be created.
        """
        return pulumi.get(self, "region")

