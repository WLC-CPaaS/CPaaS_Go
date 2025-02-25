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
import ServiceMetaflowPattern from './ServiceMetaflowPattern';

/**
 * The ServiceMetaflowOutput model module.
 * @module model/ServiceMetaflowOutput
 * @version 1.1
 */
class ServiceMetaflowOutput {
    /**
     * Constructs a new <code>ServiceMetaflowOutput</code>.
     * @alias module:model/ServiceMetaflowOutput
     */
    constructor() { 
        
        ServiceMetaflowOutput.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ServiceMetaflowOutput</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ServiceMetaflowOutput} obj Optional instance to populate.
     * @return {module:model/ServiceMetaflowOutput} The populated <code>ServiceMetaflowOutput</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ServiceMetaflowOutput();

            if (data.hasOwnProperty('binding_digit')) {
                obj['binding_digit'] = ApiClient.convertToType(data['binding_digit'], 'String');
            }
            if (data.hasOwnProperty('numbers')) {
                obj['numbers'] = ApiClient.convertToType(data['numbers'], {'String': ServiceMetaflowPattern});
            }
            if (data.hasOwnProperty('patterns')) {
                obj['patterns'] = ApiClient.convertToType(data['patterns'], {'String': ServiceMetaflowPattern});
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>ServiceMetaflowOutput</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>ServiceMetaflowOutput</code>.
     */
    static validateJSON(data) {
        // ensure the json data is a string
        if (data['binding_digit'] && !(typeof data['binding_digit'] === 'string' || data['binding_digit'] instanceof String)) {
            throw new Error("Expected the field `binding_digit` to be a primitive type in the JSON string but got " + data['binding_digit']);
        }

        return true;
    }


}



/**
 * @member {String} binding_digit
 */
ServiceMetaflowOutput.prototype['binding_digit'] = undefined;

/**
 * @member {Object.<String, module:model/ServiceMetaflowPattern>} numbers
 */
ServiceMetaflowOutput.prototype['numbers'] = undefined;

/**
 * @member {Object.<String, module:model/ServiceMetaflowPattern>} patterns
 */
ServiceMetaflowOutput.prototype['patterns'] = undefined;






export default ServiceMetaflowOutput;

