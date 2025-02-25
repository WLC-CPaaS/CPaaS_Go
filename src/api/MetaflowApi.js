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
import ServiceDocMetaflowGet from '../model/ServiceDocMetaflowGet';
import ServiceVOIPMetaflowAddData from '../model/ServiceVOIPMetaflowAddData';
import UtilCPAASError from '../model/UtilCPAASError';

/**
* Metaflow service.
* @module api/MetaflowApi
* @version 1.1
*/
export default class MetaflowApi {

    /**
    * Constructs a new MetaflowApi. 
    * @alias module:api/MetaflowApi
    * @class
    * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
    * default to {@link module:ApiClient#instance} if unspecified.
    */
    constructor(apiClient) {
        this.apiClient = apiClient || ApiClient.instance;
    }


    /**
     * Callback function to receive the result of the v1AccountAccountIDDeviceDeviceIDMetaflowDelete operation.
     * @callback module:api/MetaflowApi~v1AccountAccountIDDeviceDeviceIDMetaflowDeleteCallback
     * @param {String} error Error message, if any.
     * @param {module:model/ServiceDocMetaflowGet} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Delete Device Metaflow
     * Delete all metaflows associated with a device.
     * @param {String} accountID Account ID, 32 alpha numeric
     * @param {String} deviceID Device ID, 32 alpha numeric
     * @param {module:api/MetaflowApi~v1AccountAccountIDDeviceDeviceIDMetaflowDeleteCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/ServiceDocMetaflowGet}
     */
    v1AccountAccountIDDeviceDeviceIDMetaflowDelete(accountID, deviceID, callback) {
      let postBody = null;
      // verify the required parameter 'accountID' is set
      if (accountID === undefined || accountID === null) {
        throw new Error("Missing the required parameter 'accountID' when calling v1AccountAccountIDDeviceDeviceIDMetaflowDelete");
      }
      // verify the required parameter 'deviceID' is set
      if (deviceID === undefined || deviceID === null) {
        throw new Error("Missing the required parameter 'deviceID' when calling v1AccountAccountIDDeviceDeviceIDMetaflowDelete");
      }

      let pathParams = {
        'accountID': accountID,
        'deviceID': deviceID
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
      let returnType = ServiceDocMetaflowGet;
      return this.apiClient.callApi(
        '/v1/account/{accountID}/device/{deviceID}/metaflow', 'DELETE',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the v1AccountAccountIDDeviceDeviceIDMetaflowGet operation.
     * @callback module:api/MetaflowApi~v1AccountAccountIDDeviceDeviceIDMetaflowGetCallback
     * @param {String} error Error message, if any.
     * @param {module:model/ServiceDocMetaflowGet} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get Device Metaflow List
     * Get the list of metaflows for a device.
     * @param {String} accountID Account ID, 32 alpha numeric
     * @param {String} deviceID Device ID, 32 alpha numeric
     * @param {module:api/MetaflowApi~v1AccountAccountIDDeviceDeviceIDMetaflowGetCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/ServiceDocMetaflowGet}
     */
    v1AccountAccountIDDeviceDeviceIDMetaflowGet(accountID, deviceID, callback) {
      let postBody = null;
      // verify the required parameter 'accountID' is set
      if (accountID === undefined || accountID === null) {
        throw new Error("Missing the required parameter 'accountID' when calling v1AccountAccountIDDeviceDeviceIDMetaflowGet");
      }
      // verify the required parameter 'deviceID' is set
      if (deviceID === undefined || deviceID === null) {
        throw new Error("Missing the required parameter 'deviceID' when calling v1AccountAccountIDDeviceDeviceIDMetaflowGet");
      }

      let pathParams = {
        'accountID': accountID,
        'deviceID': deviceID
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
      let returnType = ServiceDocMetaflowGet;
      return this.apiClient.callApi(
        '/v1/account/{accountID}/device/{deviceID}/metaflow', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the v1AccountAccountIDDeviceDeviceIDMetaflowPost operation.
     * @callback module:api/MetaflowApi~v1AccountAccountIDDeviceDeviceIDMetaflowPostCallback
     * @param {String} error Error message, if any.
     * @param {module:model/ServiceDocMetaflowGet} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Create Device Metaflow
     * Create a metaflow or multiple metaflows for a device.
     * @param {String} accountID Account ID, 32 alpha numeric
     * @param {String} deviceID Device ID, 32 alpha numeric
     * @param {module:model/ServiceVOIPMetaflowAddData} reqBody payload fields
     * @param {module:api/MetaflowApi~v1AccountAccountIDDeviceDeviceIDMetaflowPostCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/ServiceDocMetaflowGet}
     */
    v1AccountAccountIDDeviceDeviceIDMetaflowPost(accountID, deviceID, reqBody, callback) {
      let postBody = reqBody;
      // verify the required parameter 'accountID' is set
      if (accountID === undefined || accountID === null) {
        throw new Error("Missing the required parameter 'accountID' when calling v1AccountAccountIDDeviceDeviceIDMetaflowPost");
      }
      // verify the required parameter 'deviceID' is set
      if (deviceID === undefined || deviceID === null) {
        throw new Error("Missing the required parameter 'deviceID' when calling v1AccountAccountIDDeviceDeviceIDMetaflowPost");
      }
      // verify the required parameter 'reqBody' is set
      if (reqBody === undefined || reqBody === null) {
        throw new Error("Missing the required parameter 'reqBody' when calling v1AccountAccountIDDeviceDeviceIDMetaflowPost");
      }

      let pathParams = {
        'accountID': accountID,
        'deviceID': deviceID
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
      let returnType = ServiceDocMetaflowGet;
      return this.apiClient.callApi(
        '/v1/account/{accountID}/device/{deviceID}/metaflow', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the v1AccountAccountIDMetaflowDelete operation.
     * @callback module:api/MetaflowApi~v1AccountAccountIDMetaflowDeleteCallback
     * @param {String} error Error message, if any.
     * @param {module:model/ServiceDocMetaflowGet} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Delete Account Metaflow
     * Remove all metaflows from an account.
     * @param {String} accountID Account ID, 32 alpha numeric
     * @param {module:api/MetaflowApi~v1AccountAccountIDMetaflowDeleteCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/ServiceDocMetaflowGet}
     */
    v1AccountAccountIDMetaflowDelete(accountID, callback) {
      let postBody = null;
      // verify the required parameter 'accountID' is set
      if (accountID === undefined || accountID === null) {
        throw new Error("Missing the required parameter 'accountID' when calling v1AccountAccountIDMetaflowDelete");
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
      let returnType = ServiceDocMetaflowGet;
      return this.apiClient.callApi(
        '/v1/account/{accountID}/metaflow', 'DELETE',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the v1AccountAccountIDMetaflowGet operation.
     * @callback module:api/MetaflowApi~v1AccountAccountIDMetaflowGetCallback
     * @param {String} error Error message, if any.
     * @param {module:model/ServiceDocMetaflowGet} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get Account Metaflow List
     * Get an account's metaflow list.
     * @param {String} accountID Account ID, 32 alpha numeric
     * @param {module:api/MetaflowApi~v1AccountAccountIDMetaflowGetCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/ServiceDocMetaflowGet}
     */
    v1AccountAccountIDMetaflowGet(accountID, callback) {
      let postBody = null;
      // verify the required parameter 'accountID' is set
      if (accountID === undefined || accountID === null) {
        throw new Error("Missing the required parameter 'accountID' when calling v1AccountAccountIDMetaflowGet");
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
      let returnType = ServiceDocMetaflowGet;
      return this.apiClient.callApi(
        '/v1/account/{accountID}/metaflow', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the v1AccountAccountIDMetaflowPost operation.
     * @callback module:api/MetaflowApi~v1AccountAccountIDMetaflowPostCallback
     * @param {String} error Error message, if any.
     * @param {module:model/ServiceDocMetaflowGet} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Create Account Metaflow
     * Generate a metaflow for an account.
     * @param {String} accountID Account ID
     * @param {module:model/ServiceVOIPMetaflowAddData} metaflow Metaflow fields
     * @param {module:api/MetaflowApi~v1AccountAccountIDMetaflowPostCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/ServiceDocMetaflowGet}
     */
    v1AccountAccountIDMetaflowPost(accountID, metaflow, callback) {
      let postBody = metaflow;
      // verify the required parameter 'accountID' is set
      if (accountID === undefined || accountID === null) {
        throw new Error("Missing the required parameter 'accountID' when calling v1AccountAccountIDMetaflowPost");
      }
      // verify the required parameter 'metaflow' is set
      if (metaflow === undefined || metaflow === null) {
        throw new Error("Missing the required parameter 'metaflow' when calling v1AccountAccountIDMetaflowPost");
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
      let returnType = ServiceDocMetaflowGet;
      return this.apiClient.callApi(
        '/v1/account/{accountID}/metaflow', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the v1AccountAccountIDUserUserIDMetaflowDelete operation.
     * @callback module:api/MetaflowApi~v1AccountAccountIDUserUserIDMetaflowDeleteCallback
     * @param {String} error Error message, if any.
     * @param {module:model/ServiceDocMetaflowGet} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Delete User Metaflow
     * Delete all metaflows associated with a user.
     * @param {String} accountID Account ID, 32 alpha numeric
     * @param {String} userID user ID, 32 alpha numeric
     * @param {module:api/MetaflowApi~v1AccountAccountIDUserUserIDMetaflowDeleteCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/ServiceDocMetaflowGet}
     */
    v1AccountAccountIDUserUserIDMetaflowDelete(accountID, userID, callback) {
      let postBody = null;
      // verify the required parameter 'accountID' is set
      if (accountID === undefined || accountID === null) {
        throw new Error("Missing the required parameter 'accountID' when calling v1AccountAccountIDUserUserIDMetaflowDelete");
      }
      // verify the required parameter 'userID' is set
      if (userID === undefined || userID === null) {
        throw new Error("Missing the required parameter 'userID' when calling v1AccountAccountIDUserUserIDMetaflowDelete");
      }

      let pathParams = {
        'accountID': accountID,
        'userID': userID
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
      let returnType = ServiceDocMetaflowGet;
      return this.apiClient.callApi(
        '/v1/account/{accountID}/user/{userID}/metaflow', 'DELETE',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the v1AccountAccountIDUserUserIDMetaflowGet operation.
     * @callback module:api/MetaflowApi~v1AccountAccountIDUserUserIDMetaflowGetCallback
     * @param {String} error Error message, if any.
     * @param {module:model/ServiceDocMetaflowGet} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get User Metaflow List
     * Get the list of metaflows for a user.
     * @param {String} accountID Account ID, 32 alpha numeric
     * @param {String} userID user ID, 32 alpha numeric
     * @param {module:api/MetaflowApi~v1AccountAccountIDUserUserIDMetaflowGetCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/ServiceDocMetaflowGet}
     */
    v1AccountAccountIDUserUserIDMetaflowGet(accountID, userID, callback) {
      let postBody = null;
      // verify the required parameter 'accountID' is set
      if (accountID === undefined || accountID === null) {
        throw new Error("Missing the required parameter 'accountID' when calling v1AccountAccountIDUserUserIDMetaflowGet");
      }
      // verify the required parameter 'userID' is set
      if (userID === undefined || userID === null) {
        throw new Error("Missing the required parameter 'userID' when calling v1AccountAccountIDUserUserIDMetaflowGet");
      }

      let pathParams = {
        'accountID': accountID,
        'userID': userID
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
      let returnType = ServiceDocMetaflowGet;
      return this.apiClient.callApi(
        '/v1/account/{accountID}/user/{userID}/metaflow', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the v1AccountAccountIDUserUserIDMetaflowPost operation.
     * @callback module:api/MetaflowApi~v1AccountAccountIDUserUserIDMetaflowPostCallback
     * @param {String} error Error message, if any.
     * @param {module:model/ServiceDocMetaflowGet} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Create User Metaflow
     * Add a metaflow or multiple metaflows for a user in an account.
     * @param {String} accountID Account ID, 32 alpha numeric
     * @param {String} userID user ID, 32 alpha numeric
     * @param {module:model/ServiceVOIPMetaflowAddData} reqBody payload fields
     * @param {module:api/MetaflowApi~v1AccountAccountIDUserUserIDMetaflowPostCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/ServiceDocMetaflowGet}
     */
    v1AccountAccountIDUserUserIDMetaflowPost(accountID, userID, reqBody, callback) {
      let postBody = reqBody;
      // verify the required parameter 'accountID' is set
      if (accountID === undefined || accountID === null) {
        throw new Error("Missing the required parameter 'accountID' when calling v1AccountAccountIDUserUserIDMetaflowPost");
      }
      // verify the required parameter 'userID' is set
      if (userID === undefined || userID === null) {
        throw new Error("Missing the required parameter 'userID' when calling v1AccountAccountIDUserUserIDMetaflowPost");
      }
      // verify the required parameter 'reqBody' is set
      if (reqBody === undefined || reqBody === null) {
        throw new Error("Missing the required parameter 'reqBody' when calling v1AccountAccountIDUserUserIDMetaflowPost");
      }

      let pathParams = {
        'accountID': accountID,
        'userID': userID
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
      let returnType = ServiceDocMetaflowGet;
      return this.apiClient.callApi(
        '/v1/account/{accountID}/user/{userID}/metaflow', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }


}
