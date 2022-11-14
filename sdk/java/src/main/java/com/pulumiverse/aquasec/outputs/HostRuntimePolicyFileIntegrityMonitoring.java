// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumiverse.aquasec.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class HostRuntimePolicyFileIntegrityMonitoring {
    /**
     * @return List of paths to be excluded from being monitored.
     * 
     */
    private @Nullable List<String> excludedPaths;
    /**
     * @return List of processes to be excluded from being monitored.
     * 
     */
    private @Nullable List<String> excludedProcesses;
    /**
     * @return List of users to be excluded from being monitored.
     * 
     */
    private @Nullable List<String> excludedUsers;
    /**
     * @return If true, add attributes operations will be monitored.
     * 
     */
    private @Nullable Boolean monitorAttributes;
    /**
     * @return If true, create operations will be monitored.
     * 
     */
    private @Nullable Boolean monitorCreate;
    /**
     * @return If true, deletion operations will be monitored.
     * 
     */
    private @Nullable Boolean monitorDelete;
    /**
     * @return If true, modification operations will be monitored.
     * 
     */
    private @Nullable Boolean monitorModify;
    /**
     * @return If true, read operations will be monitored.
     * 
     */
    private @Nullable Boolean monitorRead;
    /**
     * @return List of paths to be monitored.
     * 
     */
    private @Nullable List<String> monitoredPaths;
    /**
     * @return List of processes to be monitored.
     * 
     */
    private @Nullable List<String> monitoredProcesses;
    /**
     * @return List of users to be monitored.
     * 
     */
    private @Nullable List<String> monitoredUsers;

    private HostRuntimePolicyFileIntegrityMonitoring() {}
    /**
     * @return List of paths to be excluded from being monitored.
     * 
     */
    public List<String> excludedPaths() {
        return this.excludedPaths == null ? List.of() : this.excludedPaths;
    }
    /**
     * @return List of processes to be excluded from being monitored.
     * 
     */
    public List<String> excludedProcesses() {
        return this.excludedProcesses == null ? List.of() : this.excludedProcesses;
    }
    /**
     * @return List of users to be excluded from being monitored.
     * 
     */
    public List<String> excludedUsers() {
        return this.excludedUsers == null ? List.of() : this.excludedUsers;
    }
    /**
     * @return If true, add attributes operations will be monitored.
     * 
     */
    public Optional<Boolean> monitorAttributes() {
        return Optional.ofNullable(this.monitorAttributes);
    }
    /**
     * @return If true, create operations will be monitored.
     * 
     */
    public Optional<Boolean> monitorCreate() {
        return Optional.ofNullable(this.monitorCreate);
    }
    /**
     * @return If true, deletion operations will be monitored.
     * 
     */
    public Optional<Boolean> monitorDelete() {
        return Optional.ofNullable(this.monitorDelete);
    }
    /**
     * @return If true, modification operations will be monitored.
     * 
     */
    public Optional<Boolean> monitorModify() {
        return Optional.ofNullable(this.monitorModify);
    }
    /**
     * @return If true, read operations will be monitored.
     * 
     */
    public Optional<Boolean> monitorRead() {
        return Optional.ofNullable(this.monitorRead);
    }
    /**
     * @return List of paths to be monitored.
     * 
     */
    public List<String> monitoredPaths() {
        return this.monitoredPaths == null ? List.of() : this.monitoredPaths;
    }
    /**
     * @return List of processes to be monitored.
     * 
     */
    public List<String> monitoredProcesses() {
        return this.monitoredProcesses == null ? List.of() : this.monitoredProcesses;
    }
    /**
     * @return List of users to be monitored.
     * 
     */
    public List<String> monitoredUsers() {
        return this.monitoredUsers == null ? List.of() : this.monitoredUsers;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(HostRuntimePolicyFileIntegrityMonitoring defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<String> excludedPaths;
        private @Nullable List<String> excludedProcesses;
        private @Nullable List<String> excludedUsers;
        private @Nullable Boolean monitorAttributes;
        private @Nullable Boolean monitorCreate;
        private @Nullable Boolean monitorDelete;
        private @Nullable Boolean monitorModify;
        private @Nullable Boolean monitorRead;
        private @Nullable List<String> monitoredPaths;
        private @Nullable List<String> monitoredProcesses;
        private @Nullable List<String> monitoredUsers;
        public Builder() {}
        public Builder(HostRuntimePolicyFileIntegrityMonitoring defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.excludedPaths = defaults.excludedPaths;
    	      this.excludedProcesses = defaults.excludedProcesses;
    	      this.excludedUsers = defaults.excludedUsers;
    	      this.monitorAttributes = defaults.monitorAttributes;
    	      this.monitorCreate = defaults.monitorCreate;
    	      this.monitorDelete = defaults.monitorDelete;
    	      this.monitorModify = defaults.monitorModify;
    	      this.monitorRead = defaults.monitorRead;
    	      this.monitoredPaths = defaults.monitoredPaths;
    	      this.monitoredProcesses = defaults.monitoredProcesses;
    	      this.monitoredUsers = defaults.monitoredUsers;
        }

        @CustomType.Setter
        public Builder excludedPaths(@Nullable List<String> excludedPaths) {
            this.excludedPaths = excludedPaths;
            return this;
        }
        public Builder excludedPaths(String... excludedPaths) {
            return excludedPaths(List.of(excludedPaths));
        }
        @CustomType.Setter
        public Builder excludedProcesses(@Nullable List<String> excludedProcesses) {
            this.excludedProcesses = excludedProcesses;
            return this;
        }
        public Builder excludedProcesses(String... excludedProcesses) {
            return excludedProcesses(List.of(excludedProcesses));
        }
        @CustomType.Setter
        public Builder excludedUsers(@Nullable List<String> excludedUsers) {
            this.excludedUsers = excludedUsers;
            return this;
        }
        public Builder excludedUsers(String... excludedUsers) {
            return excludedUsers(List.of(excludedUsers));
        }
        @CustomType.Setter
        public Builder monitorAttributes(@Nullable Boolean monitorAttributes) {
            this.monitorAttributes = monitorAttributes;
            return this;
        }
        @CustomType.Setter
        public Builder monitorCreate(@Nullable Boolean monitorCreate) {
            this.monitorCreate = monitorCreate;
            return this;
        }
        @CustomType.Setter
        public Builder monitorDelete(@Nullable Boolean monitorDelete) {
            this.monitorDelete = monitorDelete;
            return this;
        }
        @CustomType.Setter
        public Builder monitorModify(@Nullable Boolean monitorModify) {
            this.monitorModify = monitorModify;
            return this;
        }
        @CustomType.Setter
        public Builder monitorRead(@Nullable Boolean monitorRead) {
            this.monitorRead = monitorRead;
            return this;
        }
        @CustomType.Setter
        public Builder monitoredPaths(@Nullable List<String> monitoredPaths) {
            this.monitoredPaths = monitoredPaths;
            return this;
        }
        public Builder monitoredPaths(String... monitoredPaths) {
            return monitoredPaths(List.of(monitoredPaths));
        }
        @CustomType.Setter
        public Builder monitoredProcesses(@Nullable List<String> monitoredProcesses) {
            this.monitoredProcesses = monitoredProcesses;
            return this;
        }
        public Builder monitoredProcesses(String... monitoredProcesses) {
            return monitoredProcesses(List.of(monitoredProcesses));
        }
        @CustomType.Setter
        public Builder monitoredUsers(@Nullable List<String> monitoredUsers) {
            this.monitoredUsers = monitoredUsers;
            return this;
        }
        public Builder monitoredUsers(String... monitoredUsers) {
            return monitoredUsers(List.of(monitoredUsers));
        }
        public HostRuntimePolicyFileIntegrityMonitoring build() {
            final var o = new HostRuntimePolicyFileIntegrityMonitoring();
            o.excludedPaths = excludedPaths;
            o.excludedProcesses = excludedProcesses;
            o.excludedUsers = excludedUsers;
            o.monitorAttributes = monitorAttributes;
            o.monitorCreate = monitorCreate;
            o.monitorDelete = monitorDelete;
            o.monitorModify = monitorModify;
            o.monitorRead = monitorRead;
            o.monitoredPaths = monitoredPaths;
            o.monitoredProcesses = monitoredProcesses;
            o.monitoredUsers = monitoredUsers;
            return o;
        }
    }
}
