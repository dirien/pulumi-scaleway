// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import io.dirien.scaleway.ContainerNamespaceArgs;
import io.dirien.scaleway.Utilities;
import io.dirien.scaleway.inputs.ContainerNamespaceState;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Creates and manages Scaleway Serverless Container Namespace.
 * For more information see [the documentation](https://developers.scaleway.com/en/products/containers/api/#namespaces-cdce79).
 * 
 * ## Examples
 * 
 * ### Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.scaleway.ContainerNamespace;
 * import com.pulumi.scaleway.ContainerNamespaceArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var main = new ContainerNamespace(&#34;main&#34;, ContainerNamespaceArgs.builder()        
 *             .description(&#34;Main container namespace&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Namespaces can be imported using the `{region}/{id}`, e.g. bash
 * 
 * ```sh
 *  $ pulumi import scaleway:index/containerNamespace:ContainerNamespace main fr-par/11111111-1111-1111-1111-111111111111
 * ```
 * 
 */
@ResourceType(type="scaleway:index/containerNamespace:ContainerNamespace")
public class ContainerNamespace extends com.pulumi.resources.CustomResource {
    /**
     * The description of the namespace.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the namespace.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Destroy registry on deletion
     * 
     * @deprecated
     * Registry namespace is automatically destroyed with namespace
     * 
     */
    @Deprecated /* Registry namespace is automatically destroyed with namespace */
    @Export(name="destroyRegistry", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> destroyRegistry;

    /**
     * @return Destroy registry on deletion
     * 
     */
    public Output<Optional<Boolean>> destroyRegistry() {
        return Codegen.optional(this.destroyRegistry);
    }
    /**
     * The environment variables of the namespace.
     * 
     */
    @Export(name="environmentVariables", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> environmentVariables;

    /**
     * @return The environment variables of the namespace.
     * 
     */
    public Output<Optional<Map<String,String>>> environmentVariables() {
        return Codegen.optional(this.environmentVariables);
    }
    /**
     * The unique name of the container namespace.
     * 
     * &gt; **Important** Updates to `name` will recreate the namespace.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The unique name of the container namespace.
     * 
     * &gt; **Important** Updates to `name` will recreate the namespace.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The organization ID the namespace is associated with.
     * 
     */
    @Export(name="organizationId", refs={String.class}, tree="[0]")
    private Output<String> organizationId;

    /**
     * @return The organization ID the namespace is associated with.
     * 
     */
    public Output<String> organizationId() {
        return this.organizationId;
    }
    /**
     * `project_id`) The ID of the project the namespace is associated with.
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output<String> projectId;

    /**
     * @return `project_id`) The ID of the project the namespace is associated with.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * `region`). The region in which the namespace should be created.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return `region`). The region in which the namespace should be created.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The registry endpoint of the namespace.
     * 
     */
    @Export(name="registryEndpoint", refs={String.class}, tree="[0]")
    private Output<String> registryEndpoint;

    /**
     * @return The registry endpoint of the namespace.
     * 
     */
    public Output<String> registryEndpoint() {
        return this.registryEndpoint;
    }
    /**
     * The registry namespace ID of the namespace.
     * 
     */
    @Export(name="registryNamespaceId", refs={String.class}, tree="[0]")
    private Output<String> registryNamespaceId;

    /**
     * @return The registry namespace ID of the namespace.
     * 
     */
    public Output<String> registryNamespaceId() {
        return this.registryNamespaceId;
    }
    /**
     * The secret environment variables of the namespace.
     * 
     */
    @Export(name="secretEnvironmentVariables", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> secretEnvironmentVariables;

    /**
     * @return The secret environment variables of the namespace.
     * 
     */
    public Output<Optional<Map<String,String>>> secretEnvironmentVariables() {
        return Codegen.optional(this.secretEnvironmentVariables);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ContainerNamespace(String name) {
        this(name, ContainerNamespaceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ContainerNamespace(String name, @Nullable ContainerNamespaceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ContainerNamespace(String name, @Nullable ContainerNamespaceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("scaleway:index/containerNamespace:ContainerNamespace", name, args == null ? ContainerNamespaceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ContainerNamespace(String name, Output<String> id, @Nullable ContainerNamespaceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("scaleway:index/containerNamespace:ContainerNamespace", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "secretEnvironmentVariables"
            ))
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static ContainerNamespace get(String name, Output<String> id, @Nullable ContainerNamespaceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ContainerNamespace(name, id, state, options);
    }
}
