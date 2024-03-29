// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import io.dirien.scaleway.RdbDatabaseArgs;
import io.dirien.scaleway.Utilities;
import io.dirien.scaleway.inputs.RdbDatabaseState;
import java.lang.Boolean;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Creates and manages Scaleway RDB database.
 * For more information, see [the documentation](https://developers.scaleway.com/en/products/rdb/api).
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
 * import com.pulumi.scaleway.RdbDatabase;
 * import com.pulumi.scaleway.RdbDatabaseArgs;
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
 *         var main = new RdbDatabase(&#34;main&#34;, RdbDatabaseArgs.builder()        
 *             .instanceId(scaleway_rdb_instance.main().id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * RDB Database can be imported using the `{region}/{id}/{DBNAME}`, e.g. bash
 * 
 * ```sh
 *  $ pulumi import scaleway:index/rdbDatabase:RdbDatabase rdb01_mydb fr-par/11111111-1111-1111-1111-111111111111/mydb
 * ```
 * 
 */
@ResourceType(type="scaleway:index/rdbDatabase:RdbDatabase")
public class RdbDatabase extends com.pulumi.resources.CustomResource {
    /**
     * UUID of the rdb instance.
     * 
     * &gt; **Important:** Updates to `instance_id` will recreate the Database.
     * 
     */
    @Export(name="instanceId", refs={String.class}, tree="[0]")
    private Output<String> instanceId;

    /**
     * @return UUID of the rdb instance.
     * 
     * &gt; **Important:** Updates to `instance_id` will recreate the Database.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }
    /**
     * Whether the database is managed or not.
     * 
     */
    @Export(name="managed", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> managed;

    /**
     * @return Whether the database is managed or not.
     * 
     */
    public Output<Boolean> managed() {
        return this.managed;
    }
    /**
     * Name of the database (e.g. `my-new-database`).
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the database (e.g. `my-new-database`).
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The name of the owner of the database.
     * 
     */
    @Export(name="owner", refs={String.class}, tree="[0]")
    private Output<String> owner;

    /**
     * @return The name of the owner of the database.
     * 
     */
    public Output<String> owner() {
        return this.owner;
    }
    /**
     * `region`) The region in which the resource exists.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return `region`) The region in which the resource exists.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * Size of the database (in bytes).
     * 
     */
    @Export(name="size", refs={String.class}, tree="[0]")
    private Output<String> size;

    /**
     * @return Size of the database (in bytes).
     * 
     */
    public Output<String> size() {
        return this.size;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RdbDatabase(String name) {
        this(name, RdbDatabaseArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RdbDatabase(String name, RdbDatabaseArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RdbDatabase(String name, RdbDatabaseArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("scaleway:index/rdbDatabase:RdbDatabase", name, args == null ? RdbDatabaseArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RdbDatabase(String name, Output<String> id, @Nullable RdbDatabaseState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("scaleway:index/rdbDatabase:RdbDatabase", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
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
    public static RdbDatabase get(String name, Output<String> id, @Nullable RdbDatabaseState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RdbDatabase(name, id, state, options);
    }
}
