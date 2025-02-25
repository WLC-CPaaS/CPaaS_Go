/**
 * White Label Communications CPaas API Documentation
 * A CPaaS platform API
 *
 * The version of the OpenAPI document: 1.1
 * Contact: support@whitelabelcomm.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */


import ApiClient from "../ApiClient";
import ServiceDocsStorageGet from '../model/ServiceDocsStorageGet';
import ServiceVOIPStorageAddEditData from '../model/ServiceVOIPStorageAddEditData';
import UtilCPAASError from '../model/UtilCPAASError';

/**
* Storage service.
* @module api/StorageApi
* @version 1.1
*/
export default class StorageApi {

    /**
    * Constructs a new StorageApi. 
    * @alias module:api/StorageApi
    * @class
    * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
    * default to {@link module:ApiClient#instance} if unspecified.
    */
    constructor(apiClient) {
        this.apiClient = apiClient || ApiClient.instance;
    }


    /**
     * Callback function to receive the result of the v1AccountAccountIDStorageDelete operation.
     * @callback module:api/StorageApi~v1AccountAccountIDStorageDeleteCallback
     * @param {String} error Error message, if any.
     * @param {module:model/ServiceDocsStorageGet} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Delete Storage
     * Delete items that are stored in an account.
     * @param {String} accountID Account ID, 32 alpha numeric
     * @param {module:api/StorageApi~v1AccountAccountIDStorageDeleteCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/ServiceDocsStorageGet}
     */
    v1AccountAccountIDStorageDelete(accountID, callback) {
      let postBody = null;
      // verify the required parameter 'accountID' is set
      if (accountID === undefined || accountID === null) {
        throw new Error("Missing the required parameter 'accountID' when calling v1AccountAccountIDStorageDelete");
      }

      let pathParams = {
        'accountID': accountID
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['BearerAuth'];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = ServiceDocsStorageGet;
      return this.apiClient.callApi(
        '/v1/account/{accountID}/storage', 'DELETE',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the v1AccountAccountIDStorageGet operation.
     * @callback module:api/StorageApi~v1AccountAccountIDStorageGetCallback
     * @param {String} error Error message, if any.
     * @param {module:model/ServiceDocsStorageGet} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get Storage Details
     * Retrieve storage details for an account.
     * @param {String} accountID Account ID, 32 alpha numeric
     * @param {module:api/StorageApi~v1AccountAccountIDStorageGetCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/ServiceDocsStorageGet}
     */
    v1AccountAccountIDStorageGet(accountID, callback) {
      let postBody = null;
      // verify the required parameter 'accountID' is set
      if (accountID === undefined || accountID === null) {
        throw new Error("Missing the required parameter 'accountID' when calling v1AccountAccountIDStorageGet");
      }

      let pathParams = {
        'accountID': accountID
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['BearerAuth'];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = ServiceDocsStorageGet;
      return this.apiClient.callApi(
        '/v1/account/{accountID}/storage', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the v1AccountAccountIDStoragePost operation.
     * @callback module:api/StorageApi~v1AccountAccountIDStoragePostCallback
     * @param {String} error Error message, if any.
     * @param {module:model/ServiceDocsStorageGet} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Create Storage
     * Create storage in an account for voicemails, call recordings, faxes, etc.
     * @param {String} accountID Account ID, 32 alpha numeric
     * @param {module:model/ServiceVOIPStorageAddEditData} reqBody payload fields
     * @param {module:api/StorageApi~v1AccountAccountIDStoragePostCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/ServiceDocsStorageGet}
     */
    v1AccountAccountIDStoragePost(accountID, reqBody, callback) {
      let postBody = reqBody;
      // verify the required parameter 'accountID' is set
      if (accountID === undefined || accountID === null) {
        throw new Error("Missing the required parameter 'accountID' when calling v1AccountAccountIDStoragePost");
      }
      // verify the required parameter 'reqBody' is set
      if (reqBody === undefined || reqBody === null) {
        throw new Error("Missing the required parameter 'reqBody' when calling v1AccountAccountIDStoragePost");
      }

      let pathParams = {
        'accountID': accountID
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['BearerAuth'];
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = ServiceDocsStorageGet;
      return this.apiClient.callApi(
        '/v1/account/{accountID}/storage', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the v1AccountAccountIDStoragePut operation.
     * @callback module:api/StorageApi~v1AccountAccountIDStoragePutCallback
     * @param {String} error Error message, if any.
     * @param {module:model/ServiceDocsStorageGet} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Update Storage
     * Modify the names of metadata to make it easier to locate (e.g., change the name of voicemail_storage to voicemail_and_callrecordings_storage, etc.).
     * @param {String} accountID Account ID, 32 alpha numeric
     * @param {module:model/ServiceVOIPStorageAddEditData} reqBody payload fields
     * @param {module:api/StorageApi~v1AccountAccountIDStoragePutCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/ServiceDocsStorageGet}
     */
    v1AccountAccountIDStoragePut(accountID, reqBody, callback) {
      let postBody = reqBody;
      // verify the required parameter 'accountID' is set
      if (accountID === undefined || accountID === null) {
        throw new Error("Missing the required parameter 'accountID' when calling v1AccountAccountIDStoragePut");
      }
      // verify the required parameter 'reqBody' is set
      if (reqBody === undefined || reqBody === null) {
        throw new Error("Missing the required parameter 'reqBody' when calling v1AccountAccountIDStoragePut");
      }

      let pathParams = {
        'accountID': accountID
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['BearerAuth'];
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = ServiceDocsStorageGet;
      return this.apiClient.callApi(
        '/v1/account/{accountID}/storage', 'PUT',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }


}
