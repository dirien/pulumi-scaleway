using System.Collections.Generic;
using Pulumi;
using Scaleway = ediri.Scaleway;

return await Deployment.RunAsync(() =>
{
   var registry = new Scaleway.RegistryNamespace("registry", new Scaleway.RegistryNamespaceArgs
   {
      Name = "pulumi-scaleway-dirien",
      IsPublic = true
   });
   return new Dictionary<string, object?>
   {
      ["endpoint"] = registry.Endpoint
   };
});
