// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway.outputs;

import com.pulumi.core.annotations.CustomType;
import io.dirien.scaleway.outputs.GetLbBackendsBackendHealthCheckHttp;
import io.dirien.scaleway.outputs.GetLbBackendsBackendHealthCheckTcp;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetLbBackendsBackend {
    /**
     * @return The date at which the backend was created (RFC 3339 format).
     * 
     */
    private String createdAt;
    /**
     * @return Scaleway S3 bucket website to be served in case all backend servers are down.
     * 
     */
    private String failoverHost;
    /**
     * @return User sessions will be forwarded to this port of backend servers.
     * 
     */
    private Integer forwardPort;
    /**
     * @return Load balancing algorithm.
     * 
     */
    private String forwardPortAlgorithm;
    /**
     * @return Backend protocol.
     * 
     */
    private String forwardProtocol;
    /**
     * @return Interval between two HC requests.
     * 
     */
    private String healthCheckDelay;
    /**
     * @return This block enable HTTP health check.
     * 
     */
    private List<GetLbBackendsBackendHealthCheckHttp> healthCheckHttp;
    /**
     * @return This block enable HTTPS health check.
     * 
     */
    private List<GetLbBackendsBackendHealthCheckHttp> healthCheckHttps;
    /**
     * @return Number of allowed failed HC requests before the backend server is marked down.
     * 
     */
    private Integer healthCheckMaxRetries;
    /**
     * @return Port the HC requests will be sent to.
     * 
     */
    private Integer healthCheckPort;
    /**
     * @return This block enable TCP health check.
     * 
     */
    private List<GetLbBackendsBackendHealthCheckTcp> healthCheckTcps;
    /**
     * @return Timeout before we consider a HC request failed.
     * 
     */
    private String healthCheckTimeout;
    /**
     * @return The associated backend ID.
     * 
     */
    private String id;
    /**
     * @return Specifies whether the Load Balancer should check the backend server’s certificate before initiating a connection.
     * 
     */
    private Boolean ignoreSslServerVerify;
    /**
     * @return The load-balancer ID this backend is attached to. backends with a LB ID like it are listed.
     * 
     */
    private String lbId;
    /**
     * @return The backend name used as filter. Backends with a name like it are listed.
     * 
     */
    private String name;
    /**
     * @return Modify what occurs when a backend server is marked down.
     * 
     */
    private String onMarkedDownAction;
    /**
     * @return The type of PROXY protocol.
     * 
     */
    private String proxyProtocol;
    /**
     * @return List of backend server IP addresses.
     * 
     */
    private List<String> serverIps;
    /**
     * @return Enables SSL between load balancer and backend servers.
     * 
     */
    private Boolean sslBridging;
    /**
     * @return Enables cookie-based session persistence.
     * 
     */
    private String stickySessions;
    /**
     * @return Cookie name for sticky sessions.
     * 
     */
    private String stickySessionsCookieName;
    /**
     * @return Maximum initial server connection establishment time.
     * 
     */
    private String timeoutConnect;
    /**
     * @return Maximum server connection inactivity time.
     * 
     */
    private String timeoutServer;
    /**
     * @return Maximum tunnel inactivity time.
     * 
     */
    private String timeoutTunnel;
    /**
     * @return The date at which the backend was last updated (RFC 3339 format).
     * 
     */
    private String updateAt;

    private GetLbBackendsBackend() {}
    /**
     * @return The date at which the backend was created (RFC 3339 format).
     * 
     */
    public String createdAt() {
        return this.createdAt;
    }
    /**
     * @return Scaleway S3 bucket website to be served in case all backend servers are down.
     * 
     */
    public String failoverHost() {
        return this.failoverHost;
    }
    /**
     * @return User sessions will be forwarded to this port of backend servers.
     * 
     */
    public Integer forwardPort() {
        return this.forwardPort;
    }
    /**
     * @return Load balancing algorithm.
     * 
     */
    public String forwardPortAlgorithm() {
        return this.forwardPortAlgorithm;
    }
    /**
     * @return Backend protocol.
     * 
     */
    public String forwardProtocol() {
        return this.forwardProtocol;
    }
    /**
     * @return Interval between two HC requests.
     * 
     */
    public String healthCheckDelay() {
        return this.healthCheckDelay;
    }
    /**
     * @return This block enable HTTP health check.
     * 
     */
    public List<GetLbBackendsBackendHealthCheckHttp> healthCheckHttp() {
        return this.healthCheckHttp;
    }
    /**
     * @return This block enable HTTPS health check.
     * 
     */
    public List<GetLbBackendsBackendHealthCheckHttp> healthCheckHttps() {
        return this.healthCheckHttps;
    }
    /**
     * @return Number of allowed failed HC requests before the backend server is marked down.
     * 
     */
    public Integer healthCheckMaxRetries() {
        return this.healthCheckMaxRetries;
    }
    /**
     * @return Port the HC requests will be sent to.
     * 
     */
    public Integer healthCheckPort() {
        return this.healthCheckPort;
    }
    /**
     * @return This block enable TCP health check.
     * 
     */
    public List<GetLbBackendsBackendHealthCheckTcp> healthCheckTcps() {
        return this.healthCheckTcps;
    }
    /**
     * @return Timeout before we consider a HC request failed.
     * 
     */
    public String healthCheckTimeout() {
        return this.healthCheckTimeout;
    }
    /**
     * @return The associated backend ID.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Specifies whether the Load Balancer should check the backend server’s certificate before initiating a connection.
     * 
     */
    public Boolean ignoreSslServerVerify() {
        return this.ignoreSslServerVerify;
    }
    /**
     * @return The load-balancer ID this backend is attached to. backends with a LB ID like it are listed.
     * 
     */
    public String lbId() {
        return this.lbId;
    }
    /**
     * @return The backend name used as filter. Backends with a name like it are listed.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Modify what occurs when a backend server is marked down.
     * 
     */
    public String onMarkedDownAction() {
        return this.onMarkedDownAction;
    }
    /**
     * @return The type of PROXY protocol.
     * 
     */
    public String proxyProtocol() {
        return this.proxyProtocol;
    }
    /**
     * @return List of backend server IP addresses.
     * 
     */
    public List<String> serverIps() {
        return this.serverIps;
    }
    /**
     * @return Enables SSL between load balancer and backend servers.
     * 
     */
    public Boolean sslBridging() {
        return this.sslBridging;
    }
    /**
     * @return Enables cookie-based session persistence.
     * 
     */
    public String stickySessions() {
        return this.stickySessions;
    }
    /**
     * @return Cookie name for sticky sessions.
     * 
     */
    public String stickySessionsCookieName() {
        return this.stickySessionsCookieName;
    }
    /**
     * @return Maximum initial server connection establishment time.
     * 
     */
    public String timeoutConnect() {
        return this.timeoutConnect;
    }
    /**
     * @return Maximum server connection inactivity time.
     * 
     */
    public String timeoutServer() {
        return this.timeoutServer;
    }
    /**
     * @return Maximum tunnel inactivity time.
     * 
     */
    public String timeoutTunnel() {
        return this.timeoutTunnel;
    }
    /**
     * @return The date at which the backend was last updated (RFC 3339 format).
     * 
     */
    public String updateAt() {
        return this.updateAt;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetLbBackendsBackend defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String createdAt;
        private String failoverHost;
        private Integer forwardPort;
        private String forwardPortAlgorithm;
        private String forwardProtocol;
        private String healthCheckDelay;
        private List<GetLbBackendsBackendHealthCheckHttp> healthCheckHttp;
        private List<GetLbBackendsBackendHealthCheckHttp> healthCheckHttps;
        private Integer healthCheckMaxRetries;
        private Integer healthCheckPort;
        private List<GetLbBackendsBackendHealthCheckTcp> healthCheckTcps;
        private String healthCheckTimeout;
        private String id;
        private Boolean ignoreSslServerVerify;
        private String lbId;
        private String name;
        private String onMarkedDownAction;
        private String proxyProtocol;
        private List<String> serverIps;
        private Boolean sslBridging;
        private String stickySessions;
        private String stickySessionsCookieName;
        private String timeoutConnect;
        private String timeoutServer;
        private String timeoutTunnel;
        private String updateAt;
        public Builder() {}
        public Builder(GetLbBackendsBackend defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.createdAt = defaults.createdAt;
    	      this.failoverHost = defaults.failoverHost;
    	      this.forwardPort = defaults.forwardPort;
    	      this.forwardPortAlgorithm = defaults.forwardPortAlgorithm;
    	      this.forwardProtocol = defaults.forwardProtocol;
    	      this.healthCheckDelay = defaults.healthCheckDelay;
    	      this.healthCheckHttp = defaults.healthCheckHttp;
    	      this.healthCheckHttps = defaults.healthCheckHttps;
    	      this.healthCheckMaxRetries = defaults.healthCheckMaxRetries;
    	      this.healthCheckPort = defaults.healthCheckPort;
    	      this.healthCheckTcps = defaults.healthCheckTcps;
    	      this.healthCheckTimeout = defaults.healthCheckTimeout;
    	      this.id = defaults.id;
    	      this.ignoreSslServerVerify = defaults.ignoreSslServerVerify;
    	      this.lbId = defaults.lbId;
    	      this.name = defaults.name;
    	      this.onMarkedDownAction = defaults.onMarkedDownAction;
    	      this.proxyProtocol = defaults.proxyProtocol;
    	      this.serverIps = defaults.serverIps;
    	      this.sslBridging = defaults.sslBridging;
    	      this.stickySessions = defaults.stickySessions;
    	      this.stickySessionsCookieName = defaults.stickySessionsCookieName;
    	      this.timeoutConnect = defaults.timeoutConnect;
    	      this.timeoutServer = defaults.timeoutServer;
    	      this.timeoutTunnel = defaults.timeoutTunnel;
    	      this.updateAt = defaults.updateAt;
        }

        @CustomType.Setter
        public Builder createdAt(String createdAt) {
            this.createdAt = Objects.requireNonNull(createdAt);
            return this;
        }
        @CustomType.Setter
        public Builder failoverHost(String failoverHost) {
            this.failoverHost = Objects.requireNonNull(failoverHost);
            return this;
        }
        @CustomType.Setter
        public Builder forwardPort(Integer forwardPort) {
            this.forwardPort = Objects.requireNonNull(forwardPort);
            return this;
        }
        @CustomType.Setter
        public Builder forwardPortAlgorithm(String forwardPortAlgorithm) {
            this.forwardPortAlgorithm = Objects.requireNonNull(forwardPortAlgorithm);
            return this;
        }
        @CustomType.Setter
        public Builder forwardProtocol(String forwardProtocol) {
            this.forwardProtocol = Objects.requireNonNull(forwardProtocol);
            return this;
        }
        @CustomType.Setter
        public Builder healthCheckDelay(String healthCheckDelay) {
            this.healthCheckDelay = Objects.requireNonNull(healthCheckDelay);
            return this;
        }
        @CustomType.Setter
        public Builder healthCheckHttp(List<GetLbBackendsBackendHealthCheckHttp> healthCheckHttp) {
            this.healthCheckHttp = Objects.requireNonNull(healthCheckHttp);
            return this;
        }
        public Builder healthCheckHttp(GetLbBackendsBackendHealthCheckHttp... healthCheckHttp) {
            return healthCheckHttp(List.of(healthCheckHttp));
        }
        @CustomType.Setter
        public Builder healthCheckHttps(List<GetLbBackendsBackendHealthCheckHttp> healthCheckHttps) {
            this.healthCheckHttps = Objects.requireNonNull(healthCheckHttps);
            return this;
        }
        public Builder healthCheckHttps(GetLbBackendsBackendHealthCheckHttp... healthCheckHttps) {
            return healthCheckHttps(List.of(healthCheckHttps));
        }
        @CustomType.Setter
        public Builder healthCheckMaxRetries(Integer healthCheckMaxRetries) {
            this.healthCheckMaxRetries = Objects.requireNonNull(healthCheckMaxRetries);
            return this;
        }
        @CustomType.Setter
        public Builder healthCheckPort(Integer healthCheckPort) {
            this.healthCheckPort = Objects.requireNonNull(healthCheckPort);
            return this;
        }
        @CustomType.Setter
        public Builder healthCheckTcps(List<GetLbBackendsBackendHealthCheckTcp> healthCheckTcps) {
            this.healthCheckTcps = Objects.requireNonNull(healthCheckTcps);
            return this;
        }
        public Builder healthCheckTcps(GetLbBackendsBackendHealthCheckTcp... healthCheckTcps) {
            return healthCheckTcps(List.of(healthCheckTcps));
        }
        @CustomType.Setter
        public Builder healthCheckTimeout(String healthCheckTimeout) {
            this.healthCheckTimeout = Objects.requireNonNull(healthCheckTimeout);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder ignoreSslServerVerify(Boolean ignoreSslServerVerify) {
            this.ignoreSslServerVerify = Objects.requireNonNull(ignoreSslServerVerify);
            return this;
        }
        @CustomType.Setter
        public Builder lbId(String lbId) {
            this.lbId = Objects.requireNonNull(lbId);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder onMarkedDownAction(String onMarkedDownAction) {
            this.onMarkedDownAction = Objects.requireNonNull(onMarkedDownAction);
            return this;
        }
        @CustomType.Setter
        public Builder proxyProtocol(String proxyProtocol) {
            this.proxyProtocol = Objects.requireNonNull(proxyProtocol);
            return this;
        }
        @CustomType.Setter
        public Builder serverIps(List<String> serverIps) {
            this.serverIps = Objects.requireNonNull(serverIps);
            return this;
        }
        public Builder serverIps(String... serverIps) {
            return serverIps(List.of(serverIps));
        }
        @CustomType.Setter
        public Builder sslBridging(Boolean sslBridging) {
            this.sslBridging = Objects.requireNonNull(sslBridging);
            return this;
        }
        @CustomType.Setter
        public Builder stickySessions(String stickySessions) {
            this.stickySessions = Objects.requireNonNull(stickySessions);
            return this;
        }
        @CustomType.Setter
        public Builder stickySessionsCookieName(String stickySessionsCookieName) {
            this.stickySessionsCookieName = Objects.requireNonNull(stickySessionsCookieName);
            return this;
        }
        @CustomType.Setter
        public Builder timeoutConnect(String timeoutConnect) {
            this.timeoutConnect = Objects.requireNonNull(timeoutConnect);
            return this;
        }
        @CustomType.Setter
        public Builder timeoutServer(String timeoutServer) {
            this.timeoutServer = Objects.requireNonNull(timeoutServer);
            return this;
        }
        @CustomType.Setter
        public Builder timeoutTunnel(String timeoutTunnel) {
            this.timeoutTunnel = Objects.requireNonNull(timeoutTunnel);
            return this;
        }
        @CustomType.Setter
        public Builder updateAt(String updateAt) {
            this.updateAt = Objects.requireNonNull(updateAt);
            return this;
        }
        public GetLbBackendsBackend build() {
            final var o = new GetLbBackendsBackend();
            o.createdAt = createdAt;
            o.failoverHost = failoverHost;
            o.forwardPort = forwardPort;
            o.forwardPortAlgorithm = forwardPortAlgorithm;
            o.forwardProtocol = forwardProtocol;
            o.healthCheckDelay = healthCheckDelay;
            o.healthCheckHttp = healthCheckHttp;
            o.healthCheckHttps = healthCheckHttps;
            o.healthCheckMaxRetries = healthCheckMaxRetries;
            o.healthCheckPort = healthCheckPort;
            o.healthCheckTcps = healthCheckTcps;
            o.healthCheckTimeout = healthCheckTimeout;
            o.id = id;
            o.ignoreSslServerVerify = ignoreSslServerVerify;
            o.lbId = lbId;
            o.name = name;
            o.onMarkedDownAction = onMarkedDownAction;
            o.proxyProtocol = proxyProtocol;
            o.serverIps = serverIps;
            o.sslBridging = sslBridging;
            o.stickySessions = stickySessions;
            o.stickySessionsCookieName = stickySessionsCookieName;
            o.timeoutConnect = timeoutConnect;
            o.timeoutServer = timeoutServer;
            o.timeoutTunnel = timeoutTunnel;
            o.updateAt = updateAt;
            return o;
        }
    }
}
