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
    public static class GetFunction
    {
        public static Task<GetFunctionResult> InvokeAsync(GetFunctionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFunctionResult>("scaleway:index/getFunction:getFunction", args ?? new GetFunctionArgs(), options.WithDefaults());

        public static Output<GetFunctionResult> Invoke(GetFunctionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFunctionResult>("scaleway:index/getFunction:getFunction", args ?? new GetFunctionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFunctionArgs : global::Pulumi.InvokeArgs
    {
        [Input("functionId")]
        public string? FunctionId { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        [Input("namespaceId", required: true)]
        public string NamespaceId { get; set; } = null!;

        public GetFunctionArgs()
        {
        }
        public static new GetFunctionArgs Empty => new GetFunctionArgs();
    }

    public sealed class GetFunctionInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("functionId")]
        public Input<string>? FunctionId { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("namespaceId", required: true)]
        public Input<string> NamespaceId { get; set; } = null!;

        public GetFunctionInvokeArgs()
        {
        }
        public static new GetFunctionInvokeArgs Empty => new GetFunctionInvokeArgs();
    }


    [OutputType]
    public sealed class GetFunctionResult
    {
        public readonly int CpuLimit;
        public readonly bool Deploy;
        public readonly string Description;
        public readonly string DomainName;
        public readonly ImmutableDictionary<string, string> EnvironmentVariables;
        public readonly string? FunctionId;
        public readonly string Handler;
        public readonly string HttpOption;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly int MaxScale;
        public readonly int MemoryLimit;
        public readonly int MinScale;
        public readonly string? Name;
        public readonly string NamespaceId;
        public readonly string OrganizationId;
        public readonly string Privacy;
        public readonly string ProjectId;
        public readonly string Region;
        public readonly string Runtime;
        public readonly ImmutableDictionary<string, string> SecretEnvironmentVariables;
        public readonly int Timeout;
        public readonly string ZipFile;
        public readonly string ZipHash;

        [OutputConstructor]
        private GetFunctionResult(
            int cpuLimit,

            bool deploy,

            string description,

            string domainName,

            ImmutableDictionary<string, string> environmentVariables,

            string? functionId,

            string handler,

            string httpOption,

            string id,

            int maxScale,

            int memoryLimit,

            int minScale,

            string? name,

            string namespaceId,

            string organizationId,

            string privacy,

            string projectId,

            string region,

            string runtime,

            ImmutableDictionary<string, string> secretEnvironmentVariables,

            int timeout,

            string zipFile,

            string zipHash)
        {
            CpuLimit = cpuLimit;
            Deploy = deploy;
            Description = description;
            DomainName = domainName;
            EnvironmentVariables = environmentVariables;
            FunctionId = functionId;
            Handler = handler;
            HttpOption = httpOption;
            Id = id;
            MaxScale = maxScale;
            MemoryLimit = memoryLimit;
            MinScale = minScale;
            Name = name;
            NamespaceId = namespaceId;
            OrganizationId = organizationId;
            Privacy = privacy;
            ProjectId = projectId;
            Region = region;
            Runtime = runtime;
            SecretEnvironmentVariables = secretEnvironmentVariables;
            Timeout = timeout;
            ZipFile = zipFile;
            ZipHash = zipHash;
        }
    }
}
