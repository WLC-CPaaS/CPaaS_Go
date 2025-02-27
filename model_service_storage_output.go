/*
 * White Label Communications CPaas API Documentation
 * A CPaaS platform API
 *
 * The version of the OpenAPI document: 1.1
 * Contact: support@whitelabelcomm.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package openapi;

import java.util.Objects;
import java.util.Arrays;

/**
 * ServiceStorageOutput
 */
@.annotation.Generated(value = "org.openapitools.codegen.languages.GoClientCodegen", comments = "Generator version: 7.11.0-SNAPSHOT")
public class ServiceStorageOutput {
  @.annotation.Nullable
  private map[string]interface{} Attachments;

  @.annotation.Nullable
  private map[string]interface{} Connections;

  @.annotation.Nullable
  private string Id;

  @.annotation.Nullable
  private ServiceStoragePlan Plan;

  public ServiceStorageOutput() {
  }

  public ServiceStorageOutput Attachments(@.annotation.Nullable map[string]interface{} Attachments) {
    
    this.Attachments = Attachments;
    return this;
  }

  public ServiceStorageOutput putAttachmentsItem(String key, interface{} AttachmentsItem) {
    if (this.Attachments == null) {
      this.Attachments = new HashMap<>();
    }
    this.Attachments.put(key, AttachmentsItem);
    return this;
  }

  /**
   * Get Attachments
   * @return Attachments
   */
  @.annotation.Nullable

  public map[string]interface{} GetAttachments() {
    return Attachments;
  }


  public void setAttachments(@.annotation.Nullable map[string]interface{} Attachments) {
    this.Attachments = Attachments;
  }

  public ServiceStorageOutput Connections(@.annotation.Nullable map[string]interface{} Connections) {
    
    this.Connections = Connections;
    return this;
  }

  public ServiceStorageOutput putConnectionsItem(String key, interface{} ConnectionsItem) {
    if (this.Connections == null) {
      this.Connections = new HashMap<>();
    }
    this.Connections.put(key, ConnectionsItem);
    return this;
  }

  /**
   * Get Connections
   * @return Connections
   */
  @.annotation.Nullable

  public map[string]interface{} GetConnections() {
    return Connections;
  }


  public void setConnections(@.annotation.Nullable map[string]interface{} Connections) {
    this.Connections = Connections;
  }

  public ServiceStorageOutput Id(@.annotation.Nullable string Id) {
    
    this.Id = Id;
    return this;
  }

  /**
   * Get Id
   * @return Id
   */
  @.annotation.Nullable

  public string GetId() {
    return Id;
  }


  public void setId(@.annotation.Nullable string Id) {
    this.Id = Id;
  }

  public ServiceStorageOutput Plan(@.annotation.Nullable ServiceStoragePlan Plan) {
    
    this.Plan = Plan;
    return this;
  }

  /**
   * Get Plan
   * @return Plan
   */
  @.annotation.Nullable

  public ServiceStoragePlan GetPlan() {
    return Plan;
  }


  public void setPlan(@.annotation.Nullable ServiceStoragePlan Plan) {
    this.Plan = Plan;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ServiceStorageOutput ServiceStorageOutput = (ServiceStorageOutput) o;
    return Objects.equals(this.Attachments, ServiceStorageOutput.Attachments) &&
        Objects.equals(this.Connections, ServiceStorageOutput.Connections) &&
        Objects.equals(this.Id, ServiceStorageOutput.Id) &&
        Objects.equals(this.Plan, ServiceStorageOutput.Plan);
  }

  @Override
  public int hashCode() {
    return Objects.hash(Attachments, Connections, Id, Plan);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ServiceStorageOutput {\n");
    sb.append("    Attachments: ").append(toIndentedString(Attachments)).append("\n");
    sb.append("    Connections: ").append(toIndentedString(Connections)).append("\n");
    sb.append("    Id: ").append(toIndentedString(Id)).append("\n");
    sb.append("    Plan: ").append(toIndentedString(Plan)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}

