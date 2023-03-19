// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ContainerState extends com.pulumi.resources.ResourceArgs {

    public static final ContainerState Empty = new ContainerState();

    /**
     * The amount of vCPU computing resources to allocate to each container. Defaults to 70.
     * 
     */
    @Import(name="cpuLimit")
    private @Nullable Output<Integer> cpuLimit;

    /**
     * @return The amount of vCPU computing resources to allocate to each container. Defaults to 70.
     * 
     */
    public Optional<Output<Integer>> cpuLimit() {
        return Optional.ofNullable(this.cpuLimit);
    }

    /**
     * The cron status of the container.
     * 
     */
    @Import(name="cronStatus")
    private @Nullable Output<String> cronStatus;

    /**
     * @return The cron status of the container.
     * 
     */
    public Optional<Output<String>> cronStatus() {
        return Optional.ofNullable(this.cronStatus);
    }

    /**
     * Boolean controlling whether the container is on a production environment.
     * 
     */
    @Import(name="deploy")
    private @Nullable Output<Boolean> deploy;

    /**
     * @return Boolean controlling whether the container is on a production environment.
     * 
     */
    public Optional<Output<Boolean>> deploy() {
        return Optional.ofNullable(this.deploy);
    }

    /**
     * The description of the container.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the container.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The native domain name of the container
     * 
     */
    @Import(name="domainName")
    private @Nullable Output<String> domainName;

    /**
     * @return The native domain name of the container
     * 
     */
    public Optional<Output<String>> domainName() {
        return Optional.ofNullable(this.domainName);
    }

    /**
     * The [environment](https://www.scaleway.com/en/docs/compute/containers/concepts/#environment-variables) variables of the container.
     * 
     */
    @Import(name="environmentVariables")
    private @Nullable Output<Map<String,String>> environmentVariables;

    /**
     * @return The [environment](https://www.scaleway.com/en/docs/compute/containers/concepts/#environment-variables) variables of the container.
     * 
     */
    public Optional<Output<Map<String,String>>> environmentVariables() {
        return Optional.ofNullable(this.environmentVariables);
    }

    /**
     * The error message of the container.
     * 
     */
    @Import(name="errorMessage")
    private @Nullable Output<String> errorMessage;

    /**
     * @return The error message of the container.
     * 
     */
    public Optional<Output<String>> errorMessage() {
        return Optional.ofNullable(this.errorMessage);
    }

    /**
     * HTTP traffic configuration
     * 
     */
    @Import(name="httpOption")
    private @Nullable Output<String> httpOption;

    /**
     * @return HTTP traffic configuration
     * 
     */
    public Optional<Output<String>> httpOption() {
        return Optional.ofNullable(this.httpOption);
    }

    /**
     * The maximum number of simultaneous requests your container can handle at the same time. Defaults to 50.
     * 
     */
    @Import(name="maxConcurrency")
    private @Nullable Output<Integer> maxConcurrency;

    /**
     * @return The maximum number of simultaneous requests your container can handle at the same time. Defaults to 50.
     * 
     */
    public Optional<Output<Integer>> maxConcurrency() {
        return Optional.ofNullable(this.maxConcurrency);
    }

    /**
     * The maximum of number of instances this container can scale to. Default to 20.
     * 
     */
    @Import(name="maxScale")
    private @Nullable Output<Integer> maxScale;

    /**
     * @return The maximum of number of instances this container can scale to. Default to 20.
     * 
     */
    public Optional<Output<Integer>> maxScale() {
        return Optional.ofNullable(this.maxScale);
    }

    /**
     * The memory computing resources in MB to allocate to each container. Defaults to 128.
     * 
     */
    @Import(name="memoryLimit")
    private @Nullable Output<Integer> memoryLimit;

    /**
     * @return The memory computing resources in MB to allocate to each container. Defaults to 128.
     * 
     */
    public Optional<Output<Integer>> memoryLimit() {
        return Optional.ofNullable(this.memoryLimit);
    }

    /**
     * The minimum of running container instances continuously. Defaults to 0.
     * 
     */
    @Import(name="minScale")
    private @Nullable Output<Integer> minScale;

    /**
     * @return The minimum of running container instances continuously. Defaults to 0.
     * 
     */
    public Optional<Output<Integer>> minScale() {
        return Optional.ofNullable(this.minScale);
    }

    /**
     * The unique name of the container name.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The unique name of the container name.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The container namespace ID of the container.
     * 
     */
    @Import(name="namespaceId")
    private @Nullable Output<String> namespaceId;

    /**
     * @return The container namespace ID of the container.
     * 
     */
    public Optional<Output<String>> namespaceId() {
        return Optional.ofNullable(this.namespaceId);
    }

    /**
     * The port to expose the container. Defaults to 8080.
     * 
     */
    @Import(name="port")
    private @Nullable Output<Integer> port;

    /**
     * @return The port to expose the container. Defaults to 8080.
     * 
     */
    public Optional<Output<Integer>> port() {
        return Optional.ofNullable(this.port);
    }

    /**
     * The privacy type define the way to authenticate to your container. Please check our dedicated [section](https://developers.scaleway.com/en/products/containers/api/#protocol-9dd4c8).
     * 
     */
    @Import(name="privacy")
    private @Nullable Output<String> privacy;

    /**
     * @return The privacy type define the way to authenticate to your container. Please check our dedicated [section](https://developers.scaleway.com/en/products/containers/api/#protocol-9dd4c8).
     * 
     */
    public Optional<Output<String>> privacy() {
        return Optional.ofNullable(this.privacy);
    }

    /**
     * The communication [protocol](https://developers.scaleway.com/en/products/containers/api/#protocol-9dd4c8) http1 or h2c. Defaults to http1.
     * 
     */
    @Import(name="protocol")
    private @Nullable Output<String> protocol;

    /**
     * @return The communication [protocol](https://developers.scaleway.com/en/products/containers/api/#protocol-9dd4c8) http1 or h2c. Defaults to http1.
     * 
     */
    public Optional<Output<String>> protocol() {
        return Optional.ofNullable(this.protocol);
    }

    /**
     * (Defaults to provider `region`) The region in which the container was created.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return (Defaults to provider `region`) The region in which the container was created.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * The registry image address. e.g: **&#34;rg.fr-par.scw.cloud/$NAMESPACE/$IMAGE&#34;**.
     * 
     */
    @Import(name="registryImage")
    private @Nullable Output<String> registryImage;

    /**
     * @return The registry image address. e.g: **&#34;rg.fr-par.scw.cloud/$NAMESPACE/$IMAGE&#34;**.
     * 
     */
    public Optional<Output<String>> registryImage() {
        return Optional.ofNullable(this.registryImage);
    }

    /**
     * The sha256 of your source registry image, changing it will re-apply the deployment. Can be any string
     * 
     */
    @Import(name="registrySha256")
    private @Nullable Output<String> registrySha256;

    /**
     * @return The sha256 of your source registry image, changing it will re-apply the deployment. Can be any string
     * 
     */
    public Optional<Output<String>> registrySha256() {
        return Optional.ofNullable(this.registrySha256);
    }

    /**
     * The [secret environment](https://www.scaleway.com/en/docs/compute/containers/concepts/#secrets) variables of the container.
     * 
     */
    @Import(name="secretEnvironmentVariables")
    private @Nullable Output<Map<String,String>> secretEnvironmentVariables;

    /**
     * @return The [secret environment](https://www.scaleway.com/en/docs/compute/containers/concepts/#secrets) variables of the container.
     * 
     */
    public Optional<Output<Map<String,String>>> secretEnvironmentVariables() {
        return Optional.ofNullable(this.secretEnvironmentVariables);
    }

    /**
     * The container status.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The container status.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * The maximum amount of time in seconds during which your container can process a request before we stop it. Defaults to 300s.
     * 
     */
    @Import(name="timeout")
    private @Nullable Output<Integer> timeout;

    /**
     * @return The maximum amount of time in seconds during which your container can process a request before we stop it. Defaults to 300s.
     * 
     */
    public Optional<Output<Integer>> timeout() {
        return Optional.ofNullable(this.timeout);
    }

    private ContainerState() {}

    private ContainerState(ContainerState $) {
        this.cpuLimit = $.cpuLimit;
        this.cronStatus = $.cronStatus;
        this.deploy = $.deploy;
        this.description = $.description;
        this.domainName = $.domainName;
        this.environmentVariables = $.environmentVariables;
        this.errorMessage = $.errorMessage;
        this.httpOption = $.httpOption;
        this.maxConcurrency = $.maxConcurrency;
        this.maxScale = $.maxScale;
        this.memoryLimit = $.memoryLimit;
        this.minScale = $.minScale;
        this.name = $.name;
        this.namespaceId = $.namespaceId;
        this.port = $.port;
        this.privacy = $.privacy;
        this.protocol = $.protocol;
        this.region = $.region;
        this.registryImage = $.registryImage;
        this.registrySha256 = $.registrySha256;
        this.secretEnvironmentVariables = $.secretEnvironmentVariables;
        this.status = $.status;
        this.timeout = $.timeout;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ContainerState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ContainerState $;

        public Builder() {
            $ = new ContainerState();
        }

        public Builder(ContainerState defaults) {
            $ = new ContainerState(Objects.requireNonNull(defaults));
        }

        /**
         * @param cpuLimit The amount of vCPU computing resources to allocate to each container. Defaults to 70.
         * 
         * @return builder
         * 
         */
        public Builder cpuLimit(@Nullable Output<Integer> cpuLimit) {
            $.cpuLimit = cpuLimit;
            return this;
        }

        /**
         * @param cpuLimit The amount of vCPU computing resources to allocate to each container. Defaults to 70.
         * 
         * @return builder
         * 
         */
        public Builder cpuLimit(Integer cpuLimit) {
            return cpuLimit(Output.of(cpuLimit));
        }

        /**
         * @param cronStatus The cron status of the container.
         * 
         * @return builder
         * 
         */
        public Builder cronStatus(@Nullable Output<String> cronStatus) {
            $.cronStatus = cronStatus;
            return this;
        }

        /**
         * @param cronStatus The cron status of the container.
         * 
         * @return builder
         * 
         */
        public Builder cronStatus(String cronStatus) {
            return cronStatus(Output.of(cronStatus));
        }

        /**
         * @param deploy Boolean controlling whether the container is on a production environment.
         * 
         * @return builder
         * 
         */
        public Builder deploy(@Nullable Output<Boolean> deploy) {
            $.deploy = deploy;
            return this;
        }

        /**
         * @param deploy Boolean controlling whether the container is on a production environment.
         * 
         * @return builder
         * 
         */
        public Builder deploy(Boolean deploy) {
            return deploy(Output.of(deploy));
        }

        /**
         * @param description The description of the container.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the container.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param domainName The native domain name of the container
         * 
         * @return builder
         * 
         */
        public Builder domainName(@Nullable Output<String> domainName) {
            $.domainName = domainName;
            return this;
        }

        /**
         * @param domainName The native domain name of the container
         * 
         * @return builder
         * 
         */
        public Builder domainName(String domainName) {
            return domainName(Output.of(domainName));
        }

        /**
         * @param environmentVariables The [environment](https://www.scaleway.com/en/docs/compute/containers/concepts/#environment-variables) variables of the container.
         * 
         * @return builder
         * 
         */
        public Builder environmentVariables(@Nullable Output<Map<String,String>> environmentVariables) {
            $.environmentVariables = environmentVariables;
            return this;
        }

        /**
         * @param environmentVariables The [environment](https://www.scaleway.com/en/docs/compute/containers/concepts/#environment-variables) variables of the container.
         * 
         * @return builder
         * 
         */
        public Builder environmentVariables(Map<String,String> environmentVariables) {
            return environmentVariables(Output.of(environmentVariables));
        }

        /**
         * @param errorMessage The error message of the container.
         * 
         * @return builder
         * 
         */
        public Builder errorMessage(@Nullable Output<String> errorMessage) {
            $.errorMessage = errorMessage;
            return this;
        }

        /**
         * @param errorMessage The error message of the container.
         * 
         * @return builder
         * 
         */
        public Builder errorMessage(String errorMessage) {
            return errorMessage(Output.of(errorMessage));
        }

        /**
         * @param httpOption HTTP traffic configuration
         * 
         * @return builder
         * 
         */
        public Builder httpOption(@Nullable Output<String> httpOption) {
            $.httpOption = httpOption;
            return this;
        }

        /**
         * @param httpOption HTTP traffic configuration
         * 
         * @return builder
         * 
         */
        public Builder httpOption(String httpOption) {
            return httpOption(Output.of(httpOption));
        }

        /**
         * @param maxConcurrency The maximum number of simultaneous requests your container can handle at the same time. Defaults to 50.
         * 
         * @return builder
         * 
         */
        public Builder maxConcurrency(@Nullable Output<Integer> maxConcurrency) {
            $.maxConcurrency = maxConcurrency;
            return this;
        }

        /**
         * @param maxConcurrency The maximum number of simultaneous requests your container can handle at the same time. Defaults to 50.
         * 
         * @return builder
         * 
         */
        public Builder maxConcurrency(Integer maxConcurrency) {
            return maxConcurrency(Output.of(maxConcurrency));
        }

        /**
         * @param maxScale The maximum of number of instances this container can scale to. Default to 20.
         * 
         * @return builder
         * 
         */
        public Builder maxScale(@Nullable Output<Integer> maxScale) {
            $.maxScale = maxScale;
            return this;
        }

        /**
         * @param maxScale The maximum of number of instances this container can scale to. Default to 20.
         * 
         * @return builder
         * 
         */
        public Builder maxScale(Integer maxScale) {
            return maxScale(Output.of(maxScale));
        }

        /**
         * @param memoryLimit The memory computing resources in MB to allocate to each container. Defaults to 128.
         * 
         * @return builder
         * 
         */
        public Builder memoryLimit(@Nullable Output<Integer> memoryLimit) {
            $.memoryLimit = memoryLimit;
            return this;
        }

        /**
         * @param memoryLimit The memory computing resources in MB to allocate to each container. Defaults to 128.
         * 
         * @return builder
         * 
         */
        public Builder memoryLimit(Integer memoryLimit) {
            return memoryLimit(Output.of(memoryLimit));
        }

        /**
         * @param minScale The minimum of running container instances continuously. Defaults to 0.
         * 
         * @return builder
         * 
         */
        public Builder minScale(@Nullable Output<Integer> minScale) {
            $.minScale = minScale;
            return this;
        }

        /**
         * @param minScale The minimum of running container instances continuously. Defaults to 0.
         * 
         * @return builder
         * 
         */
        public Builder minScale(Integer minScale) {
            return minScale(Output.of(minScale));
        }

        /**
         * @param name The unique name of the container name.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The unique name of the container name.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param namespaceId The container namespace ID of the container.
         * 
         * @return builder
         * 
         */
        public Builder namespaceId(@Nullable Output<String> namespaceId) {
            $.namespaceId = namespaceId;
            return this;
        }

        /**
         * @param namespaceId The container namespace ID of the container.
         * 
         * @return builder
         * 
         */
        public Builder namespaceId(String namespaceId) {
            return namespaceId(Output.of(namespaceId));
        }

        /**
         * @param port The port to expose the container. Defaults to 8080.
         * 
         * @return builder
         * 
         */
        public Builder port(@Nullable Output<Integer> port) {
            $.port = port;
            return this;
        }

        /**
         * @param port The port to expose the container. Defaults to 8080.
         * 
         * @return builder
         * 
         */
        public Builder port(Integer port) {
            return port(Output.of(port));
        }

        /**
         * @param privacy The privacy type define the way to authenticate to your container. Please check our dedicated [section](https://developers.scaleway.com/en/products/containers/api/#protocol-9dd4c8).
         * 
         * @return builder
         * 
         */
        public Builder privacy(@Nullable Output<String> privacy) {
            $.privacy = privacy;
            return this;
        }

        /**
         * @param privacy The privacy type define the way to authenticate to your container. Please check our dedicated [section](https://developers.scaleway.com/en/products/containers/api/#protocol-9dd4c8).
         * 
         * @return builder
         * 
         */
        public Builder privacy(String privacy) {
            return privacy(Output.of(privacy));
        }

        /**
         * @param protocol The communication [protocol](https://developers.scaleway.com/en/products/containers/api/#protocol-9dd4c8) http1 or h2c. Defaults to http1.
         * 
         * @return builder
         * 
         */
        public Builder protocol(@Nullable Output<String> protocol) {
            $.protocol = protocol;
            return this;
        }

        /**
         * @param protocol The communication [protocol](https://developers.scaleway.com/en/products/containers/api/#protocol-9dd4c8) http1 or h2c. Defaults to http1.
         * 
         * @return builder
         * 
         */
        public Builder protocol(String protocol) {
            return protocol(Output.of(protocol));
        }

        /**
         * @param region (Defaults to provider `region`) The region in which the container was created.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region (Defaults to provider `region`) The region in which the container was created.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param registryImage The registry image address. e.g: **&#34;rg.fr-par.scw.cloud/$NAMESPACE/$IMAGE&#34;**.
         * 
         * @return builder
         * 
         */
        public Builder registryImage(@Nullable Output<String> registryImage) {
            $.registryImage = registryImage;
            return this;
        }

        /**
         * @param registryImage The registry image address. e.g: **&#34;rg.fr-par.scw.cloud/$NAMESPACE/$IMAGE&#34;**.
         * 
         * @return builder
         * 
         */
        public Builder registryImage(String registryImage) {
            return registryImage(Output.of(registryImage));
        }

        /**
         * @param registrySha256 The sha256 of your source registry image, changing it will re-apply the deployment. Can be any string
         * 
         * @return builder
         * 
         */
        public Builder registrySha256(@Nullable Output<String> registrySha256) {
            $.registrySha256 = registrySha256;
            return this;
        }

        /**
         * @param registrySha256 The sha256 of your source registry image, changing it will re-apply the deployment. Can be any string
         * 
         * @return builder
         * 
         */
        public Builder registrySha256(String registrySha256) {
            return registrySha256(Output.of(registrySha256));
        }

        /**
         * @param secretEnvironmentVariables The [secret environment](https://www.scaleway.com/en/docs/compute/containers/concepts/#secrets) variables of the container.
         * 
         * @return builder
         * 
         */
        public Builder secretEnvironmentVariables(@Nullable Output<Map<String,String>> secretEnvironmentVariables) {
            $.secretEnvironmentVariables = secretEnvironmentVariables;
            return this;
        }

        /**
         * @param secretEnvironmentVariables The [secret environment](https://www.scaleway.com/en/docs/compute/containers/concepts/#secrets) variables of the container.
         * 
         * @return builder
         * 
         */
        public Builder secretEnvironmentVariables(Map<String,String> secretEnvironmentVariables) {
            return secretEnvironmentVariables(Output.of(secretEnvironmentVariables));
        }

        /**
         * @param status The container status.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The container status.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param timeout The maximum amount of time in seconds during which your container can process a request before we stop it. Defaults to 300s.
         * 
         * @return builder
         * 
         */
        public Builder timeout(@Nullable Output<Integer> timeout) {
            $.timeout = timeout;
            return this;
        }

        /**
         * @param timeout The maximum amount of time in seconds during which your container can process a request before we stop it. Defaults to 300s.
         * 
         * @return builder
         * 
         */
        public Builder timeout(Integer timeout) {
            return timeout(Output.of(timeout));
        }

        public ContainerState build() {
            return $;
        }
    }

}
