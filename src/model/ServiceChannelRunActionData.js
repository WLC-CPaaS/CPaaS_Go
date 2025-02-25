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

/**
 * The ServiceChannelRunActionData model module.
 * @module model/ServiceChannelRunActionData
 * @version 1.1
 */
class ServiceChannelRunActionData {
    /**
     * Constructs a new <code>ServiceChannelRunActionData</code>.
     * @alias module:model/ServiceChannelRunActionData
     * @param action {String} 
     */
    constructor(action) { 
        
        ServiceChannelRunActionData.initialize(this, action);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, action) { 
        obj['action'] = action;
    }

    /**
     * Constructs a <code>ServiceChannelRunActionData</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ServiceChannelRunActionData} obj Optional instance to populate.
     * @return {module:model/ServiceChannelRunActionData} The populated <code>ServiceChannelRunActionData</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ServiceChannelRunActionData();

            if (data.hasOwnProperty('action')) {
                obj['action'] = ApiClient.convertToType(data['action'], 'String');
            }
            if (data.hasOwnProperty('moh')) {
                obj['moh'] = ApiClient.convertToType(data['moh'], 'String');
            }
            if (data.hasOwnProperty('takeback_dtmf')) {
                obj['takeback_dtmf'] = ApiClient.convertToType(data['takeback_dtmf'], 'String');
            }
            if (data.hasOwnProperty('target')) {
                obj['target'] = ApiClient.convertToType(data['target'], 'String');
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>ServiceChannelRunActionData</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>ServiceChannelRunActionData</code>.
     */
    static validateJSON(data) {
        // check to make sure all required properties are present in the JSON string
        for (const property of ServiceChannelRunActionData.RequiredProperties) {
            if (!data.hasOwnProperty(property)) {
                throw new Error("The required field `" + property + "` is not found in the JSON data: " + JSON.stringify(data));
            }
        }
        // ensure the json data is a string
        if (data['action'] && !(typeof data['action'] === 'string' || data['action'] instanceof String)) {
            throw new Error("Expected the field `action` to be a primitive type in the JSON string but got " + data['action']);
        }
        // ensure the json data is a string
        if (data['moh'] && !(typeof data['moh'] === 'string' || data['moh'] instanceof String)) {
            throw new Error("Expected the field `moh` to be a primitive type in the JSON string but got " + data['moh']);
        }
        // ensure the json data is a string
        if (data['takeback_dtmf'] && !(typeof data['takeback_dtmf'] === 'string' || data['takeback_dtmf'] instanceof String)) {
            throw new Error("Expected the field `takeback_dtmf` to be a primitive type in the JSON string but got " + data['takeback_dtmf']);
        }
        // ensure the json data is a string
        if (data['target'] && !(typeof data['target'] === 'string' || data['target'] instanceof String)) {
            throw new Error("Expected the field `target` to be a primitive type in the JSON string but got " + data['target']);
        }

        return true;
    }


}

ServiceChannelRunActionData.RequiredProperties = ["action"];

/**
 * @member {String} action
 */
ServiceChannelRunActionData.prototype['action'] = undefined;

/**
 * @member {String} moh
 */
ServiceChannelRunActionData.prototype['moh'] = undefined;

/**
 * @member {String} takeback_dtmf
 */
ServiceChannelRunActionData.prototype['takeback_dtmf'] = undefined;

/**
 * @member {String} target
 */
ServiceChannelRunActionData.prototype['target'] = undefined;






export default ServiceChannelRunActionData;

