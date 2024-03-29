// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ediri.Scaleway
{
    public static class GetIpamIp
    {
        /// <summary>
        /// Gets information about IP managed by IPAM service. IPAM service is used for dhcp bundled in VPCs' private networks.
        /// 
        /// ## Examples
        /// 
        /// ### Instance Private Network IP
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// using Scaleway = ediri.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     // Find the private IPv4 using resource name
        ///     var pn = new Scaleway.VpcPrivateNetwork("pn");
        /// 
        ///     // Get Instance IP in a private network
        ///     var nic = new Scaleway.InstancePrivateNic("nic", new()
        ///     {
        ///         ServerId = scaleway_instance_server.Server.Id,
        ///         PrivateNetworkId = pn.Id,
        ///     });
        /// 
        ///     var byMac = Scaleway.GetIpamIp.Invoke(new()
        ///     {
        ///         MacAddress = nic.MacAddress,
        ///         Type = "ipv4",
        ///     });
        /// 
        ///     var byId = Scaleway.GetIpamIp.Invoke(new()
        ///     {
        ///         Resource = new Scaleway.Inputs.GetIpamIpResourceInputArgs
        ///         {
        ///             Id = nic.Id,
        ///             Type = "instance_private_nic",
        ///         },
        ///         Type = "ipv4",
        ///     });
        /// 
        ///     var main = new Scaleway.RdbInstance("main", new()
        ///     {
        ///         NodeType = "DB-DEV-S",
        ///         Engine = "PostgreSQL-15",
        ///         IsHaCluster = true,
        ///         DisableBackup = true,
        ///         UserName = "my_initial_user",
        ///         Password = "thiZ_is_v&amp;ry_s3cret",
        ///         PrivateNetwork = new Scaleway.Inputs.RdbInstancePrivateNetworkArgs
        ///         {
        ///             PnId = pn.Id,
        ///         },
        ///     });
        /// 
        ///     var byName = Scaleway.GetIpamIp.Invoke(new()
        ///     {
        ///         Resource = new Scaleway.Inputs.GetIpamIpResourceInputArgs
        ///         {
        ///             Name = main.Name,
        ///             Type = "rdb_instance",
        ///         },
        ///         Type = "ipv4",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetIpamIpResult> InvokeAsync(GetIpamIpArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIpamIpResult>("scaleway:index/getIpamIp:getIpamIp", args ?? new GetIpamIpArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about IP managed by IPAM service. IPAM service is used for dhcp bundled in VPCs' private networks.
        /// 
        /// ## Examples
        /// 
        /// ### Instance Private Network IP
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// using Scaleway = ediri.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     // Find the private IPv4 using resource name
        ///     var pn = new Scaleway.VpcPrivateNetwork("pn");
        /// 
        ///     // Get Instance IP in a private network
        ///     var nic = new Scaleway.InstancePrivateNic("nic", new()
        ///     {
        ///         ServerId = scaleway_instance_server.Server.Id,
        ///         PrivateNetworkId = pn.Id,
        ///     });
        /// 
        ///     var byMac = Scaleway.GetIpamIp.Invoke(new()
        ///     {
        ///         MacAddress = nic.MacAddress,
        ///         Type = "ipv4",
        ///     });
        /// 
        ///     var byId = Scaleway.GetIpamIp.Invoke(new()
        ///     {
        ///         Resource = new Scaleway.Inputs.GetIpamIpResourceInputArgs
        ///         {
        ///             Id = nic.Id,
        ///             Type = "instance_private_nic",
        ///         },
        ///         Type = "ipv4",
        ///     });
        /// 
        ///     var main = new Scaleway.RdbInstance("main", new()
        ///     {
        ///         NodeType = "DB-DEV-S",
        ///         Engine = "PostgreSQL-15",
        ///         IsHaCluster = true,
        ///         DisableBackup = true,
        ///         UserName = "my_initial_user",
        ///         Password = "thiZ_is_v&amp;ry_s3cret",
        ///         PrivateNetwork = new Scaleway.Inputs.RdbInstancePrivateNetworkArgs
        ///         {
        ///             PnId = pn.Id,
        ///         },
        ///     });
        /// 
        ///     var byName = Scaleway.GetIpamIp.Invoke(new()
        ///     {
        ///         Resource = new Scaleway.Inputs.GetIpamIpResourceInputArgs
        ///         {
        ///             Name = main.Name,
        ///             Type = "rdb_instance",
        ///         },
        ///         Type = "ipv4",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetIpamIpResult> Invoke(GetIpamIpInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIpamIpResult>("scaleway:index/getIpamIp:getIpamIp", args ?? new GetIpamIpInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIpamIpArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Mac Address linked to the IP.
        /// </summary>
        [Input("macAddress")]
        public string? MacAddress { get; set; }

        /// <summary>
        /// The ID of the private network the IP belong to.
        /// </summary>
        [Input("privateNetworkId")]
        public string? PrivateNetworkId { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the IP is associated with.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        /// <summary>
        /// `region`) The region in which the IP exists.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        /// <summary>
        /// Filter by resource ID, type or name. If specified, `type` is required, and at least one of `id` or `name` must be set.
        /// </summary>
        [Input("resource")]
        public Inputs.GetIpamIpResourceArgs? Resource { get; set; }

        [Input("tags")]
        private List<string>? _tags;

        /// <summary>
        /// The tags associated with the IP.
        /// As datasource only returns one IP, the search with given tags must return only one result.
        /// </summary>
        public List<string> Tags
        {
            get => _tags ?? (_tags = new List<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The type of the resource to get the IP from. [Documentation](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/ipam/v1#pkg-constants) with type list.
        /// </summary>
        [Input("type", required: true)]
        public string Type { get; set; } = null!;

        /// <summary>
        /// Only IPs that are zonal, and in this zone, will be returned.
        /// </summary>
        [Input("zonal")]
        public string? Zonal { get; set; }

        public GetIpamIpArgs()
        {
        }
        public static new GetIpamIpArgs Empty => new GetIpamIpArgs();
    }

    public sealed class GetIpamIpInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Mac Address linked to the IP.
        /// </summary>
        [Input("macAddress")]
        public Input<string>? MacAddress { get; set; }

        /// <summary>
        /// The ID of the private network the IP belong to.
        /// </summary>
        [Input("privateNetworkId")]
        public Input<string>? PrivateNetworkId { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the IP is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`) The region in which the IP exists.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Filter by resource ID, type or name. If specified, `type` is required, and at least one of `id` or `name` must be set.
        /// </summary>
        [Input("resource")]
        public Input<Inputs.GetIpamIpResourceInputArgs>? Resource { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The tags associated with the IP.
        /// As datasource only returns one IP, the search with given tags must return only one result.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The type of the resource to get the IP from. [Documentation](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/ipam/v1#pkg-constants) with type list.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// Only IPs that are zonal, and in this zone, will be returned.
        /// </summary>
        [Input("zonal")]
        public Input<string>? Zonal { get; set; }

        public GetIpamIpInvokeArgs()
        {
        }
        public static new GetIpamIpInvokeArgs Empty => new GetIpamIpInvokeArgs();
    }


    [OutputType]
    public sealed class GetIpamIpResult
    {
        /// <summary>
        /// The IP address
        /// </summary>
        public readonly string Address;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? MacAddress;
        public readonly string OrganizationId;
        public readonly string? PrivateNetworkId;
        public readonly string ProjectId;
        public readonly string Region;
        public readonly Outputs.GetIpamIpResourceResult? Resource;
        public readonly ImmutableArray<string> Tags;
        public readonly string Type;
        public readonly string Zonal;

        [OutputConstructor]
        private GetIpamIpResult(
            string address,

            string id,

            string? macAddress,

            string organizationId,

            string? privateNetworkId,

            string projectId,

            string region,

            Outputs.GetIpamIpResourceResult? resource,

            ImmutableArray<string> tags,

            string type,

            string zonal)
        {
            Address = address;
            Id = id;
            MacAddress = macAddress;
            OrganizationId = organizationId;
            PrivateNetworkId = privateNetworkId;
            ProjectId = projectId;
            Region = region;
            Resource = resource;
            Tags = tags;
            Type = type;
            Zonal = zonal;
        }
    }
}
