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
 * ServiceDocsAccountGetAll
 */
@.annotation.Generated(value = "org.openapitools.codegen.languages.GoClientCodegen", comments = "Generator version: 7.11.0-SNAPSHOT")
public class ServiceDocsAccountGetAll {
  @.annotation.Nullable
  private []ServiceAccountOutput Data;

  @.annotation.Nullable
  private string NextStartKey;

  @.annotation.Nullable
  private int32 PageSize;

  @.annotation.Nullable
  private string RequestId;

  @.annotation.Nullable
  private string StartKey;

  @.annotation.Nullable
  private int32 StatusCode;

  public ServiceDocsAccountGetAll() {
  }

  public ServiceDocsAccountGetAll Data(@.annotation.Nullable []ServiceAccountOutput Data) {
    
    this.Data = Data;
    return this;
  }

  public ServiceDocsAccountGetAll addDataItem(ServiceAccountOutput DataItem) {
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

  public []ServiceAccountOutput GetData() {
    return Data;
  }


  public void setData(@.annotation.Nullable []ServiceAccountOutput Data) {
    this.Data = Data;
  }

  public ServiceDocsAccountGetAll NextStartKey(@.annotation.Nullable string NextStartKey) {
    
    this.NextStartKey = NextStartKey;
    return this;
  }

  /**
   * Get NextStartKey
   * @return NextStartKey
   */
  @.annotation.Nullable

  public string GetNextStartKey() {
    return NextStartKey;
  }


  public void setNextStartKey(@.annotation.Nullable string NextStartKey) {
    this.NextStartKey = NextStartKey;
  }

  public ServiceDocsAccountGetAll PageSize(@.annotation.Nullable int32 PageSize) {
    
    this.PageSize = PageSize;
    return this;
  }

  /**
   * Get PageSize
   * @return PageSize
   */
  @.annotation.Nullable

  public int32 GetPageSize() {
    return PageSize;
  }


  public void setPageSize(@.annotation.Nullable int32 PageSize) {
    this.PageSize = PageSize;
  }

  public ServiceDocsAccountGetAll RequestId(@.annotation.Nullable string RequestId) {
    
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

  public ServiceDocsAccountGetAll StartKey(@.annotation.Nullable string StartKey) {
    
    this.StartKey = StartKey;
    return this;
  }

  /**
   * Get StartKey
   * @return StartKey
   */
  @.annotation.Nullable

  public string GetStartKey() {
    return StartKey;
  }


  public void setStartKey(@.annotation.Nullable string StartKey) {
    this.StartKey = StartKey;
  }

  public ServiceDocsAccountGetAll StatusCode(@.annotation.Nullable int32 StatusCode) {
    
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
    ServiceDocsAccountGetAll ServiceDocsAccountGetAll = (ServiceDocsAccountGetAll) o;
    return Objects.equals(this.Data, ServiceDocsAccountGetAll.Data) &&
        Objects.equals(this.NextStartKey, ServiceDocsAccountGetAll.NextStartKey) &&
        Objects.equals(this.PageSize, ServiceDocsAccountGetAll.PageSize) &&
        Objects.equals(this.RequestId, ServiceDocsAccountGetAll.RequestId) &&
        Objects.equals(this.StartKey, ServiceDocsAccountGetAll.StartKey) &&
        Objects.equals(this.StatusCode, ServiceDocsAccountGetAll.StatusCode);
  }

  @Override
  public int hashCode() {
    return Objects.hash(Data, NextStartKey, PageSize, RequestId, StartKey, StatusCode);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ServiceDocsAccountGetAll {\n");
    sb.append("    Data: ").append(toIndentedString(Data)).append("\n");
    sb.append("    NextStartKey: ").append(toIndentedString(NextStartKey)).append("\n");
    sb.append("    PageSize: ").append(toIndentedString(PageSize)).append("\n");
    sb.append("    RequestId: ").append(toIndentedString(RequestId)).append("\n");
    sb.append("    StartKey: ").append(toIndentedString(StartKey)).append("\n");
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

