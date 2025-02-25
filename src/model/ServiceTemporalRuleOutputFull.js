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
 * The ServiceTemporalRuleOutputFull model module.
 * @module model/ServiceTemporalRuleOutputFull
 * @version 1.1
 */
class ServiceTemporalRuleOutputFull {
    /**
     * Constructs a new <code>ServiceTemporalRuleOutputFull</code>.
     * @alias module:model/ServiceTemporalRuleOutputFull
     */
    constructor() { 
        
        ServiceTemporalRuleOutputFull.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ServiceTemporalRuleOutputFull</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ServiceTemporalRuleOutputFull} obj Optional instance to populate.
     * @return {module:model/ServiceTemporalRuleOutputFull} The populated <code>ServiceTemporalRuleOutputFull</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ServiceTemporalRuleOutputFull();

            if (data.hasOwnProperty('cycle')) {
                obj['cycle'] = ApiClient.convertToType(data['cycle'], 'String');
            }
            if (data.hasOwnProperty('days')) {
                obj['days'] = ApiClient.convertToType(data['days'], ['Number']);
            }
            if (data.hasOwnProperty('enabled')) {
                obj['enabled'] = ApiClient.convertToType(data['enabled'], 'Boolean');
            }
            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'String');
            }
            if (data.hasOwnProperty('interval')) {
                obj['interval'] = ApiClient.convertToType(data['interval'], 'Number');
            }
            if (data.hasOwnProperty('month')) {
                obj['month'] = ApiClient.convertToType(data['month'], 'Number');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('ordinal')) {
                obj['ordinal'] = ApiClient.convertToType(data['ordinal'], 'String');
            }
            if (data.hasOwnProperty('start_date')) {
                obj['start_date'] = ApiClient.convertToType(data['start_date'], 'Number');
            }
            if (data.hasOwnProperty('time_window_start')) {
                obj['time_window_start'] = ApiClient.convertToType(data['time_window_start'], 'Number');
            }
            if (data.hasOwnProperty('time_window_stop')) {
                obj['time_window_stop'] = ApiClient.convertToType(data['time_window_stop'], 'Number');
            }
            if (data.hasOwnProperty('wdays')) {
                obj['wdays'] = ApiClient.convertToType(data['wdays'], ['String']);
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>ServiceTemporalRuleOutputFull</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>ServiceTemporalRuleOutputFull</code>.
     */
    static validateJSON(data) {
        // ensure the json data is a string
        if (data['cycle'] && !(typeof data['cycle'] === 'string' || data['cycle'] instanceof String)) {
            throw new Error("Expected the field `cycle` to be a primitive type in the JSON string but got " + data['cycle']);
        }
        // ensure the json data is an array
        if (!Array.isArray(data['days'])) {
            throw new Error("Expected the field `days` to be an array in the JSON data but got " + data['days']);
        }
        // ensure the json data is a string
        if (data['id'] && !(typeof data['id'] === 'string' || data['id'] instanceof String)) {
            throw new Error("Expected the field `id` to be a primitive type in the JSON string but got " + data['id']);
        }
        // ensure the json data is a string
        if (data['name'] && !(typeof data['name'] === 'string' || data['name'] instanceof String)) {
            throw new Error("Expected the field `name` to be a primitive type in the JSON string but got " + data['name']);
        }
        // ensure the json data is a string
        if (data['ordinal'] && !(typeof data['ordinal'] === 'string' || data['ordinal'] instanceof String)) {
            throw new Error("Expected the field `ordinal` to be a primitive type in the JSON string but got " + data['ordinal']);
        }
        // ensure the json data is an array
        if (!Array.isArray(data['wdays'])) {
            throw new Error("Expected the field `wdays` to be an array in the JSON data but got " + data['wdays']);
        }

        return true;
    }


}



/**
 * @member {String} cycle
 */
ServiceTemporalRuleOutputFull.prototype['cycle'] = undefined;

/**
 * @member {Array.<Number>} days
 */
ServiceTemporalRuleOutputFull.prototype['days'] = undefined;

/**
 * @member {Boolean} enabled
 */
ServiceTemporalRuleOutputFull.prototype['enabled'] = undefined;

/**
 * @member {String} id
 */
ServiceTemporalRuleOutputFull.prototype['id'] = undefined;

/**
 * @member {Number} interval
 */
ServiceTemporalRuleOutputFull.prototype['interval'] = undefined;

/**
 * @member {Number} month
 */
ServiceTemporalRuleOutputFull.prototype['month'] = undefined;

/**
 * @member {String} name
 */
ServiceTemporalRuleOutputFull.prototype['name'] = undefined;

/**
 * @member {String} ordinal
 */
ServiceTemporalRuleOutputFull.prototype['ordinal'] = undefined;

/**
 * @member {Number} start_date
 */
ServiceTemporalRuleOutputFull.prototype['start_date'] = undefined;

/**
 * @member {Number} time_window_start
 */
ServiceTemporalRuleOutputFull.prototype['time_window_start'] = undefined;

/**
 * @member {Number} time_window_stop
 */
ServiceTemporalRuleOutputFull.prototype['time_window_stop'] = undefined;

/**
 * @member {Array.<String>} wdays
 */
ServiceTemporalRuleOutputFull.prototype['wdays'] = undefined;






export default ServiceTemporalRuleOutputFull;

