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

import ApiClient from '../ApiClient';
import ServiceE911LocationURI from './ServiceE911LocationURI';

/**
 * The ServiceDocE911LocationsURIApiOutput model module.
 * @module model/ServiceDocE911LocationsURIApiOutput
 * @version 1.1
 */
class ServiceDocE911LocationsURIApiOutput {
    /**
     * Constructs a new <code>ServiceDocE911LocationsURIApiOutput</code>.
     * @alias module:model/ServiceDocE911LocationsURIApiOutput
     */
    constructor() { 
        
        ServiceDocE911LocationsURIApiOutput.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ServiceDocE911LocationsURIApiOutput</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ServiceDocE911LocationsURIApiOutput} obj Optional instance to populate.
     * @return {module:model/ServiceDocE911LocationsURIApiOutput} The populated <code>ServiceDocE911LocationsURIApiOutput</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ServiceDocE911LocationsURIApiOutput();

            if (data.hasOwnProperty('data')) {
                obj['data'] = ApiClient.convertToType(data['data'], [ServiceE911LocationURI]);
            }
            if (data.hasOwnProperty('request_id')) {
                obj['request_id'] = ApiClient.convertToType(data['request_id'], 'String');
            }
            if (data.hasOwnProperty('status_code')) {
                obj['status_code'] = ApiClient.convertToType(data['status_code'], 'Number');
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>ServiceDocE911LocationsURIApiOutput</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>ServiceDocE911LocationsURIApiOutput</code>.
     */
    static validateJSON(data) {
        if (data['data']) { // data not null
            // ensure the json data is an array
            if (!Array.isArray(data['data'])) {
                throw new Error("Expected the field `data` to be an array in the JSON data but got " + data['data']);
            }
            // validate the optional field `data` (array)
            for (const item of data['data']) {
                ServiceE911LocationURI.validateJSON(item);
            };
        }
        // ensure the json data is a string
        if (data['request_id'] && !(typeof data['request_id'] === 'string' || data['request_id'] instanceof String)) {
            throw new Error("Expected the field `request_id` to be a primitive type in the JSON string but got " + data['request_id']);
        }

        return true;
    }


}



/**
 * @member {Array.<module:model/ServiceE911LocationURI>} data
 */
ServiceDocE911LocationsURIApiOutput.prototype['data'] = undefined;

/**
 * @member {String} request_id
 */
ServiceDocE911LocationsURIApiOutput.prototype['request_id'] = undefined;

/**
 * @member {Number} status_code
 */
ServiceDocE911LocationsURIApiOutput.prototype['status_code'] = undefined;






export default ServiceDocE911LocationsURIApiOutput;

