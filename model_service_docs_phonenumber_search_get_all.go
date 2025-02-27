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
 * ServiceDocsPhonenumberSearchGetAll
 */
@.annotation.Generated(value = "org.openapitools.codegen.languages.GoClientCodegen", comments = "Generator version: 7.11.0-SNAPSHOT")
public class ServiceDocsPhonenumberSearchGetAll {
  @.annotation.Nullable
  private []ServicePhoneNumberSearchOutput Data;

  @.annotation.Nullable
  private string RequestId;

  @.annotation.Nullable
  private int32 StatusCode;

  public ServiceDocsPhonenumberSearchGetAll() {
  }

  public ServiceDocsPhonenumberSearchGetAll Data(@.annotation.Nullable []ServicePhoneNumberSearchOutput Data) {
    
    this.Data = Data;
    return this;
  }

  public ServiceDocsPhonenumberSearchGetAll addDataItem(ServicePhoneNumberSearchOutput DataItem) {
    if (this.Data == null) {
      this.Data = new ArrayList<>();
    }
    this.Data.add(DataItem);
    return this;
  }

  /**
   * Get Data
   * @return Data
   */
  @.annotation.Nullable

  public []ServicePhoneNumberSearchOutput GetData() {
    return Data;
  }


  public void setData(@.annotation.Nullable []ServicePhoneNumberSearchOutput Data) {
    this.Data = Data;
  }

  public ServiceDocsPhonenumberSearchGetAll RequestId(@.annotation.Nullable string RequestId) {
    
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

  public ServiceDocsPhonenumberSearchGetAll StatusCode(@.annotation.Nullable int32 StatusCode) {
    
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
    ServiceDocsPhonenumberSearchGetAll ServiceDocsPhonenumberSearchGetAll = (ServiceDocsPhonenumberSearchGetAll) o;
    return Objects.equals(this.Data, ServiceDocsPhonenumberSearchGetAll.Data) &&
        Objects.equals(this.RequestId, ServiceDocsPhonenumberSearchGetAll.RequestId) &&
        Objects.equals(this.StatusCode, ServiceDocsPhonenumberSearchGetAll.StatusCode);
  }

  @Override
  public int hashCode() {
    return Objects.hash(Data, RequestId, StatusCode);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ServiceDocsPhonenumberSearchGetAll {\n");
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

