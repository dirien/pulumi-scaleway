import * as pulumi from "@pulumi/pulumi";
import * as scaleway from "@ediri/scaleway";

const registry = new scaleway.RegistryNamespace("registry", {
    name: "pulumi-scaleway-dirien",
    isPublic: true,
})

export const registryId = registry.endpoint
