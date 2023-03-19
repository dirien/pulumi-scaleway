"use strict";
const pulumi = require("@pulumi/pulumi");
const scaleway = require("@ediri/scaleway");

const registry = new scaleway.RegistryNamespace("registry", {
    name: "pulumi-scaleway-dirien",
    isPublic: true,
})

exports.endpoint = registry.endpoint
