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
import bytes;
import fmt;

/**
 * ServiceVOIPCallQueueAddEditData
 */
@.annotation.Generated(value = "org.openapitools.codegen.languages.GoClientCodegen", comments = "Generator version: 7.11.0-SNAPSHOT")
public class ServiceVOIPCallQueueAddEditData {
  @.annotation.Nullable
  private int32 AgentWrapupTime;

  @.annotation.Nullable
  private map[string]interface{} Features;

  @.annotation.Nullable
  private bool ForceAwayOnReject;

  @.annotation.Nonnull
  private string Name;

  @.annotation.Nullable
  private string QueueRouter;

  @.annotation.Nullable
  private string QueueType;

  @.annotation.Nullable
  private int32 RingTimeout;

  @.annotation.Nullable
  private int32 TickTime;

  public ServiceVOIPCallQueueAddEditData() {
  }

  public ServiceVOIPCallQueueAddEditData AgentWrapupTime(@.annotation.Nullable int32 AgentWrapupTime) {
    
    this.AgentWrapupTime = AgentWrapupTime;
    return this;
  }

  /**
   * Get AgentWrapupTime
   * @return AgentWrapupTime
   */
  @.annotation.Nullable

  public int32 GetAgentWrapupTime() {
    return AgentWrapupTime;
  }


  public void setAgentWrapupTime(@.annotation.Nullable int32 AgentWrapupTime) {
    this.AgentWrapupTime = AgentWrapupTime;
  }

  public ServiceVOIPCallQueueAddEditData Features(@.annotation.Nullable map[string]interface{} Features) {
    
    this.Features = Features;
    return this;
  }

  public ServiceVOIPCallQueueAddEditData putFeaturesItem(String key, interface{} FeaturesItem) {
    if (this.Features == null) {
      this.Features = new HashMap<>();
    }
    this.Features.put(key, FeaturesItem);
    return this;
  }

  /**
   * Get Features
   * @return Features
   */
  @.annotation.Nullable

  public map[string]interface{} GetFeatures() {
    return Features;
  }


  public void setFeatures(@.annotation.Nullable map[string]interface{} Features) {
    this.Features = Features;
  }

  public ServiceVOIPCallQueueAddEditData ForceAwayOnReject(@.annotation.Nullable bool ForceAwayOnReject) {
    
    this.ForceAwayOnReject = ForceAwayOnReject;
    return this;
  }

  /**
   * Get ForceAwayOnReject
   * @return ForceAwayOnReject
   */
  @.annotation.Nullable

  public bool getForceAwayOnReject() {
    return ForceAwayOnReject;
  }


  public void setForceAwayOnReject(@.annotation.Nullable bool ForceAwayOnReject) {
    this.ForceAwayOnReject = ForceAwayOnReject;
  }

  public ServiceVOIPCallQueueAddEditData Name(@.annotation.Nonnull string Name) {
    
    this.Name = Name;
    return this;
  }

  /**
   * Get Name
   * @return Name
   */
  @.annotation.Nonnull

  public string GetName() {
    return Name;
  }


  public void setName(@.annotation.Nonnull string Name) {
    this.Name = Name;
  }

  public ServiceVOIPCallQueueAddEditData QueueRouter(@.annotation.Nullable string QueueRouter) {
    
    this.QueueRouter = QueueRouter;
    return this;
  }

  /**
   * Get QueueRouter
   * @return QueueRouter
   */
  @.annotation.Nullable

  public string GetQueueRouter() {
    return QueueRouter;
  }


  public void setQueueRouter(@.annotation.Nullable string QueueRouter) {
    this.QueueRouter = QueueRouter;
  }

  public ServiceVOIPCallQueueAddEditData QueueType(@.annotation.Nullable string QueueType) {
    
    this.QueueType = QueueType;
    return this;
  }

  /**
   * Get QueueType
   * @return QueueType
   */
  @.annotation.Nullable

  public string GetQueueType() {
    return QueueType;
  }


  public void setQueueType(@.annotation.Nullable string QueueType) {
    this.QueueType = QueueType;
  }

  public ServiceVOIPCallQueueAddEditData RingTimeout(@.annotation.Nullable int32 RingTimeout) {
    
    this.RingTimeout = RingTimeout;
    return this;
  }

  /**
   * Get RingTimeout
   * @return RingTimeout
   */
  @.annotation.Nullable

  public int32 GetRingTimeout() {
    return RingTimeout;
  }


  public void setRingTimeout(@.annotation.Nullable int32 RingTimeout) {
    this.RingTimeout = RingTimeout;
  }

  public ServiceVOIPCallQueueAddEditData TickTime(@.annotation.Nullable int32 TickTime) {
    
    this.TickTime = TickTime;
    return this;
  }

  /**
   * Get TickTime
   * @return TickTime
   */
  @.annotation.Nullable

  public int32 GetTickTime() {
    return TickTime;
  }


  public void setTickTime(@.annotation.Nullable int32 TickTime) {
    this.TickTime = TickTime;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ServiceVOIPCallQueueAddEditData ServiceVOIPCallQueueAddEditData = (ServiceVOIPCallQueueAddEditData) o;
    return Objects.equals(this.AgentWrapupTime, ServiceVOIPCallQueueAddEditData.AgentWrapupTime) &&
        Objects.equals(this.Features, ServiceVOIPCallQueueAddEditData.Features) &&
        Objects.equals(this.ForceAwayOnReject, ServiceVOIPCallQueueAddEditData.ForceAwayOnReject) &&
        Objects.equals(this.Name, ServiceVOIPCallQueueAddEditData.Name) &&
        Objects.equals(this.QueueRouter, ServiceVOIPCallQueueAddEditData.QueueRouter) &&
        Objects.equals(this.QueueType, ServiceVOIPCallQueueAddEditData.QueueType) &&
        Objects.equals(this.RingTimeout, ServiceVOIPCallQueueAddEditData.RingTimeout) &&
        Objects.equals(this.TickTime, ServiceVOIPCallQueueAddEditData.TickTime);
  }

  @Override
  public int hashCode() {
    return Objects.hash(AgentWrapupTime, Features, ForceAwayOnReject, Name, QueueRouter, QueueType, RingTimeout, TickTime);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ServiceVOIPCallQueueAddEditData {\n");
    sb.append("    AgentWrapupTime: ").append(toIndentedString(AgentWrapupTime)).append("\n");
    sb.append("    Features: ").append(toIndentedString(Features)).append("\n");
    sb.append("    ForceAwayOnReject: ").append(toIndentedString(ForceAwayOnReject)).append("\n");
    sb.append("    Name: ").append(toIndentedString(Name)).append("\n");
    sb.append("    QueueRouter: ").append(toIndentedString(QueueRouter)).append("\n");
    sb.append("    QueueType: ").append(toIndentedString(QueueType)).append("\n");
    sb.append("    RingTimeout: ").append(toIndentedString(RingTimeout)).append("\n");
    sb.append("    TickTime: ").append(toIndentedString(TickTime)).append("\n");
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

