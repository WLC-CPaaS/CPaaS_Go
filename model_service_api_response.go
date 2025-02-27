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
 * ServiceAPIResponse
 */
@.annotation.Generated(value = "org.openapitools.codegen.languages.GoClientCodegen", comments = "Generator version: 7.11.0-SNAPSHOT")
public class ServiceAPIResponse {
  @.annotation.Nullable
  private map[string]interface{} Data;

  @.annotation.Nullable
  private string RequestId;

  @.annotation.Nullable
  private int32 StatusCode;

  public ServiceAPIResponse() {
  }

  public ServiceAPIResponse Data(@.annotation.Nullable map[string]interface{} Data) {
    
    this.Data = Data;
    return this;
  }

  /**
   * Get Data
   * @return Data
   */
  @.annotation.Nullable

  public map[string]interface{} GetData() {
    return Data;
  }


  public void setData(@.annotation.Nullable map[string]interface{} Data) {
    this.Data = Data;
  }

  public ServiceAPIResponse RequestId(@.annotation.Nullable string RequestId) {
    
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

  public ServiceAPIResponse StatusCode(@.annotation.Nullable int32 StatusCode) {
    
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
    ServiceAPIResponse ServiceAPIResponse = (ServiceAPIResponse) o;
    return Objects.equals(this.Data, ServiceAPIResponse.Data) &&
        Objects.equals(this.RequestId, ServiceAPIResponse.RequestId) &&
        Objects.equals(this.StatusCode, ServiceAPIResponse.StatusCode);
  }

  @Override
  public int hashCode() {
    return Objects.hash(Data, RequestId, StatusCode);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ServiceAPIResponse {\n");
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

