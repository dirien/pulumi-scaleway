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
    public static class GetLbBackend
    {
        /// <summary>
        /// Get information about Scaleway Load-Balancer Backends.
        /// For more information, see [the documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-backends).
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
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
        ///     var mainLbIp = new Scaleway.LbIp("mainLbIp");
        /// 
        ///     var mainLb = new Scaleway.Lb("mainLb", new()
        ///     {
        ///         IpId = mainLbIp.Id,
        ///         Type = "LB-S",
        ///     });
        /// 
        ///     var mainLbBackend = new Scaleway.LbBackend("mainLbBackend", new()
        ///     {
        ///         LbId = mainLb.Id,
        ///         ForwardProtocol = "http",
        ///         ForwardPort = 80,
        ///     });
        /// 
        ///     var byID = Scaleway.GetLbBackend.Invoke(new()
        ///     {
        ///         BackendId = mainLbBackend.Id,
        ///     });
        /// 
        ///     var byName = Scaleway.GetLbBackend.Invoke(new()
        ///     {
        ///         Name = mainLbBackend.Name,
        ///         LbId = mainLb.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetLbBackendResult> InvokeAsync(GetLbBackendArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLbBackendResult>("scaleway:index/getLbBackend:getLbBackend", args ?? new GetLbBackendArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about Scaleway Load-Balancer Backends.
        /// For more information, see [the documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-backends).
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
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
        ///     var mainLbIp = new Scaleway.LbIp("mainLbIp");
        /// 
        ///     var mainLb = new Scaleway.Lb("mainLb", new()
        ///     {
        ///         IpId = mainLbIp.Id,
        ///         Type = "LB-S",
        ///     });
        /// 
        ///     var mainLbBackend = new Scaleway.LbBackend("mainLbBackend", new()
        ///     {
        ///         LbId = mainLb.Id,
        ///         ForwardProtocol = "http",
        ///         ForwardPort = 80,
        ///     });
        /// 
        ///     var byID = Scaleway.GetLbBackend.Invoke(new()
        ///     {
        ///         BackendId = mainLbBackend.Id,
        ///     });
        /// 
        ///     var byName = Scaleway.GetLbBackend.Invoke(new()
        ///     {
        ///         Name = mainLbBackend.Name,
        ///         LbId = mainLb.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetLbBackendResult> Invoke(GetLbBackendInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLbBackendResult>("scaleway:index/getLbBackend:getLbBackend", args ?? new GetLbBackendInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLbBackendArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The backend id.
        /// - Only one of `name` and `backend_id` should be specified.
        /// </summary>
        [Input("backendId")]
        public string? BackendId { get; set; }

        /// <summary>
        /// The load-balancer ID this backend is attached to.
        /// </summary>
        [Input("lbId")]
        public string? LbId { get; set; }

        /// <summary>
        /// The name of the backend.
        /// - When using the `name` you should specify the `lb-id`
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetLbBackendArgs()
        {
        }
        public static new GetLbBackendArgs Empty => new GetLbBackendArgs();
    }

    public sealed class GetLbBackendInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The backend id.
        /// - Only one of `name` and `backend_id` should be specified.
        /// </summary>
        [Input("backendId")]
        public Input<string>? BackendId { get; set; }

        /// <summary>
        /// The load-balancer ID this backend is attached to.
        /// </summary>
        [Input("lbId")]
        public Input<string>? LbId { get; set; }

        /// <summary>
        /// The name of the backend.
        /// - When using the `name` you should specify the `lb-id`
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetLbBackendInvokeArgs()
        {
        }
        public static new GetLbBackendInvokeArgs Empty => new GetLbBackendInvokeArgs();
    }


    [OutputType]
    public sealed class GetLbBackendResult
    {
        public readonly string? BackendId;
        public readonly string FailoverHost;
        public readonly int ForwardPort;
        public readonly string ForwardPortAlgorithm;
        public readonly string ForwardProtocol;
        public readonly string HealthCheckDelay;
        public readonly ImmutableArray<Outputs.GetLbBackendHealthCheckHttpResult> HealthCheckHttp;
        public readonly ImmutableArray<Outputs.GetLbBackendHealthCheckHttpResult> HealthCheckHttps;
        public readonly int HealthCheckMaxRetries;
        public readonly int HealthCheckPort;
        public readonly ImmutableArray<Outputs.GetLbBackendHealthCheckTcpResult> HealthCheckTcps;
        public readonly string HealthCheckTimeout;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool IgnoreSslServerVerify;
        public readonly string? LbId;
        public readonly string? Name;
        public readonly string OnMarkedDownAction;
        public readonly string ProxyProtocol;
        public readonly bool SendProxyV2;
        public readonly ImmutableArray<string> ServerIps;
        public readonly bool SslBridging;
        public readonly string StickySessions;
        public readonly string StickySessionsCookieName;
        public readonly string TimeoutConnect;
        public readonly string TimeoutServer;
        public readonly string TimeoutTunnel;

        [OutputConstructor]
        private GetLbBackendResult(
            string? backendId,

            string failoverHost,

            int forwardPort,

            string forwardPortAlgorithm,

            string forwardProtocol,

            string healthCheckDelay,

            ImmutableArray<Outputs.GetLbBackendHealthCheckHttpResult> healthCheckHttp,

            ImmutableArray<Outputs.GetLbBackendHealthCheckHttpResult> healthCheckHttps,

            int healthCheckMaxRetries,

            int healthCheckPort,

            ImmutableArray<Outputs.GetLbBackendHealthCheckTcpResult> healthCheckTcps,

            string healthCheckTimeout,

            string id,

            bool ignoreSslServerVerify,

            string? lbId,

            string? name,

            string onMarkedDownAction,

            string proxyProtocol,

            bool sendProxyV2,

            ImmutableArray<string> serverIps,

            bool sslBridging,

            string stickySessions,

            string stickySessionsCookieName,

            string timeoutConnect,

            string timeoutServer,

            string timeoutTunnel)
        {
            BackendId = backendId;
            FailoverHost = failoverHost;
            ForwardPort = forwardPort;
            ForwardPortAlgorithm = forwardPortAlgorithm;
            ForwardProtocol = forwardProtocol;
            HealthCheckDelay = healthCheckDelay;
            HealthCheckHttp = healthCheckHttp;
            HealthCheckHttps = healthCheckHttps;
            HealthCheckMaxRetries = healthCheckMaxRetries;
            HealthCheckPort = healthCheckPort;
            HealthCheckTcps = healthCheckTcps;
            HealthCheckTimeout = healthCheckTimeout;
            Id = id;
            IgnoreSslServerVerify = ignoreSslServerVerify;
            LbId = lbId;
            Name = name;
            OnMarkedDownAction = onMarkedDownAction;
            ProxyProtocol = proxyProtocol;
            SendProxyV2 = sendProxyV2;
            ServerIps = serverIps;
            SslBridging = sslBridging;
            StickySessions = stickySessions;
            StickySessionsCookieName = stickySessionsCookieName;
            TimeoutConnect = timeoutConnect;
            TimeoutServer = timeoutServer;
            TimeoutTunnel = timeoutTunnel;
        }
    }
}
