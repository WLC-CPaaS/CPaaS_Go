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
 * ServiceDocsChannelGetSingle
 */
@.annotation.Generated(value = "org.openapitools.codegen.languages.GoClientCodegen", comments = "Generator version: 7.11.0-SNAPSHOT")
public class ServiceDocsChannelGetSingle {
  @.annotation.Nullable
  private ServiceChannelOutput Data;

  @.annotation.Nullable
  private string RequestId;

  @.annotation.Nullable
  private int32 StatusCode;

  public ServiceDocsChannelGetSingle() {
  }

  public ServiceDocsChannelGetSingle Data(@.annotation.Nullable ServiceChannelOutput Data) {
    
    this.Data = Data;
    return this;
  }

  /**
   * Get Data
   * @return Data
   */
  @.annotation.Nullable

  public ServiceChannelOutput GetData() {
    return Data;
  }


  public void setData(@.annotation.Nullable ServiceChannelOutput Data) {
    this.Data = Data;
  }

  public ServiceDocsChannelGetSingle RequestId(@.annotation.Nullable string RequestId) {
    
    this.RequestId = RequestId;
    return this;
  }

  /**
   * Get RequestId
   * @return RequestId
   */
  @.annotation.Nullable

  public string GetRequestId() {
    return RequestId;
  }


  public void setRequestId(@.annotation.Nullable string RequestId) {
    this.RequestId = RequestId;
  }

  public ServiceDocsChannelGetSingle StatusCode(@.annotation.Nullable int32 StatusCode) {
    
    this.StatusCode = StatusCode;
    return this;
  }

  /**
   * Get StatusCode
   * @return StatusCode
   */
  @.annotation.Nullable

  public int32 GetStatusCode() {
    return StatusCode;
  }


  public void setStatusCode(@.annotation.Nullable int32 StatusCode) {
    this.StatusCode = StatusCode;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ServiceDocsChannelGetSingle ServiceDocsChannelGetSingle = (ServiceDocsChannelGetSingle) o;
    return Objects.equals(this.Data, ServiceDocsChannelGetSingle.Data) &&
        Objects.equals(this.RequestId, ServiceDocsChannelGetSingle.RequestId) &&
        Objects.equals(this.StatusCode, ServiceDocsChannelGetSingle.StatusCode);
  }

  @Override
  public int hashCode() {
    return Objects.hash(Data, RequestId, StatusCode);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ServiceDocsChannelGetSingle {\n");
    sb.append("    Data: ").append(toIndentedString(Data)).append("\n");
    sb.append("    RequestId: ").append(toIndentedString(RequestId)).append("\n");
    sb.append("    StatusCode: ").append(toIndentedString(StatusCode)).append("\n");
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

