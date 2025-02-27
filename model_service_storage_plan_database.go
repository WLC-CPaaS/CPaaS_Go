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
 * ServiceStoragePlanDatabase
 */
@.annotation.Generated(value = "org.openapitools.codegen.languages.GoClientCodegen", comments = "Generator version: 7.11.0-SNAPSHOT")
public class ServiceStoragePlanDatabase {
  @.annotation.Nullable
  private ServiceStoragePlanDatabaseAttachment Attachments;

  @.annotation.Nullable
  private string Connection;

  @.annotation.Nullable
  private ServiceStoragePlanDatabaseProperties Database;

  @.annotation.Nullable
  private ServiceStoragePlanDatabaseTypes Types;

  public ServiceStoragePlanDatabase() {
  }

  public ServiceStoragePlanDatabase Attachments(@.annotation.Nullable ServiceStoragePlanDatabaseAttachment Attachments) {
    
    this.Attachments = Attachments;
    return this;
  }

  /**
   * Get Attachments
   * @return Attachments
   */
  @.annotation.Nullable

  public ServiceStoragePlanDatabaseAttachment GetAttachments() {
    return Attachments;
  }


  public void setAttachments(@.annotation.Nullable ServiceStoragePlanDatabaseAttachment Attachments) {
    this.Attachments = Attachments;
  }

  public ServiceStoragePlanDatabase Connection(@.annotation.Nullable string Connection) {
    
    this.Connection = Connection;
    return this;
  }

  /**
   * Get Connection
   * @return Connection
   */
  @.annotation.Nullable

  public string GetConnection() {
    return Connection;
  }


  public void setConnection(@.annotation.Nullable string Connection) {
    this.Connection = Connection;
  }

  public ServiceStoragePlanDatabase Database(@.annotation.Nullable ServiceStoragePlanDatabaseProperties Database) {
    
    this.Database = Database;
    return this;
  }

  /**
   * Get Database
   * @return Database
   */
  @.annotation.Nullable

  public ServiceStoragePlanDatabaseProperties GetDatabase() {
    return Database;
  }


  public void setDatabase(@.annotation.Nullable ServiceStoragePlanDatabaseProperties Database) {
    this.Database = Database;
  }

  public ServiceStoragePlanDatabase Types(@.annotation.Nullable ServiceStoragePlanDatabaseTypes Types) {
    
    this.Types = Types;
    return this;
  }

  /**
   * Get Types
   * @return Types
   */
  @.annotation.Nullable

  public ServiceStoragePlanDatabaseTypes GetTypes() {
    return Types;
  }


  public void setTypes(@.annotation.Nullable ServiceStoragePlanDatabaseTypes Types) {
    this.Types = Types;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ServiceStoragePlanDatabase ServiceStoragePlanDatabase = (ServiceStoragePlanDatabase) o;
    return Objects.equals(this.Attachments, ServiceStoragePlanDatabase.Attachments) &&
        Objects.equals(this.Connection, ServiceStoragePlanDatabase.Connection) &&
        Objects.equals(this.Database, ServiceStoragePlanDatabase.Database) &&
        Objects.equals(this.Types, ServiceStoragePlanDatabase.Types);
  }

  @Override
  public int hashCode() {
    return Objects.hash(Attachments, Connection, Database, Types);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ServiceStoragePlanDatabase {\n");
    sb.append("    Attachments: ").append(toIndentedString(Attachments)).append("\n");
    sb.append("    Connection: ").append(toIndentedString(Connection)).append("\n");
    sb.append("    Database: ").append(toIndentedString(Database)).append("\n");
    sb.append("    Types: ").append(toIndentedString(Types)).append("\n");
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

