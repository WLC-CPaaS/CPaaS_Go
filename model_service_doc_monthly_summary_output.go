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
 * ServiceDocMonthlySummaryOutput
 */
@.annotation.Generated(value = "org.openapitools.codegen.languages.GoClientCodegen", comments = "Generator version: 7.11.0-SNAPSHOT")
public class ServiceDocMonthlySummaryOutput {
  @.annotation.Nullable
  private ModelMonthlySummaryReport Data;

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

  public ServiceDocMonthlySummaryOutput() {
  }

  public ServiceDocMonthlySummaryOutput Data(@.annotation.Nullable ModelMonthlySummaryReport Data) {
    
    this.Data = Data;
    return this;
  }

  /**
   * Get Data
   * @return Data
   */
  @.annotation.Nullable

  public ModelMonthlySummaryReport GetData() {
    return Data;
  }


  public void setData(@.annotation.Nullable ModelMonthlySummaryReport Data) {
    this.Data = Data;
  }

  public ServiceDocMonthlySummaryOutput NextStartKey(@.annotation.Nullable string NextStartKey) {
    
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

  public ServiceDocMonthlySummaryOutput PageSize(@.annotation.Nullable int32 PageSize) {
    
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

  public ServiceDocMonthlySummaryOutput RequestId(@.annotation.Nullable string RequestId) {
    
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

  public ServiceDocMonthlySummaryOutput StartKey(@.annotation.Nullable string StartKey) {
    
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

  public ServiceDocMonthlySummaryOutput StatusCode(@.annotation.Nullable int32 StatusCode) {
    
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
    ServiceDocMonthlySummaryOutput ServiceDocMonthlySummaryOutput = (ServiceDocMonthlySummaryOutput) o;
    return Objects.equals(this.Data, ServiceDocMonthlySummaryOutput.Data) &&
        Objects.equals(this.NextStartKey, ServiceDocMonthlySummaryOutput.NextStartKey) &&
        Objects.equals(this.PageSize, ServiceDocMonthlySummaryOutput.PageSize) &&
        Objects.equals(this.RequestId, ServiceDocMonthlySummaryOutput.RequestId) &&
        Objects.equals(this.StartKey, ServiceDocMonthlySummaryOutput.StartKey) &&
        Objects.equals(this.StatusCode, ServiceDocMonthlySummaryOutput.StatusCode);
  }

  @Override
  public int hashCode() {
    return Objects.hash(Data, NextStartKey, PageSize, RequestId, StartKey, StatusCode);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ServiceDocMonthlySummaryOutput {\n");
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

