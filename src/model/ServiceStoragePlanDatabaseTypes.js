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
import ServiceStoragePlanDatabaseDocument from './ServiceStoragePlanDatabaseDocument';

/**
 * The ServiceStoragePlanDatabaseTypes model module.
 * @module model/ServiceStoragePlanDatabaseTypes
 * @version 1.1
 */
class ServiceStoragePlanDatabaseTypes {
    /**
     * Constructs a new <code>ServiceStoragePlanDatabaseTypes</code>.
     * @alias module:model/ServiceStoragePlanDatabaseTypes
     */
    constructor() { 
        
        ServiceStoragePlanDatabaseTypes.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ServiceStoragePlanDatabaseTypes</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ServiceStoragePlanDatabaseTypes} obj Optional instance to populate.
     * @return {module:model/ServiceStoragePlanDatabaseTypes} The populated <code>ServiceStoragePlanDatabaseTypes</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ServiceStoragePlanDatabaseTypes();

            if (data.hasOwnProperty('call_recording')) {
                obj['call_recording'] = ServiceStoragePlanDatabaseDocument.constructFromObject(data['call_recording']);
            }
            if (data.hasOwnProperty('fax')) {
                obj['fax'] = ServiceStoragePlanDatabaseDocument.constructFromObject(data['fax']);
            }
            if (data.hasOwnProperty('function')) {
                obj['function'] = ServiceStoragePlanDatabaseDocument.constructFromObject(data['function']);
            }
            if (data.hasOwnProperty('mailbox_message')) {
                obj['mailbox_message'] = ServiceStoragePlanDatabaseDocument.constructFromObject(data['mailbox_message']);
            }
            if (data.hasOwnProperty('media')) {
                obj['media'] = ServiceStoragePlanDatabaseDocument.constructFromObject(data['media']);
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>ServiceStoragePlanDatabaseTypes</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>ServiceStoragePlanDatabaseTypes</code>.
     */
    static validateJSON(data) {
        // validate the optional field `call_recording`
        if (data['call_recording']) { // data not null
          ServiceStoragePlanDatabaseDocument.validateJSON(data['call_recording']);
        }
        // validate the optional field `fax`
        if (data['fax']) { // data not null
          ServiceStoragePlanDatabaseDocument.validateJSON(data['fax']);
        }
        // validate the optional field `function`
        if (data['function']) { // data not null
          ServiceStoragePlanDatabaseDocument.validateJSON(data['function']);
        }
        // validate the optional field `mailbox_message`
        if (data['mailbox_message']) { // data not null
          ServiceStoragePlanDatabaseDocument.validateJSON(data['mailbox_message']);
        }
        // validate the optional field `media`
        if (data['media']) { // data not null
          ServiceStoragePlanDatabaseDocument.validateJSON(data['media']);
        }

        return true;
    }


}



/**
 * @member {module:model/ServiceStoragePlanDatabaseDocument} call_recording
 */
ServiceStoragePlanDatabaseTypes.prototype['call_recording'] = undefined;

/**
 * @member {module:model/ServiceStoragePlanDatabaseDocument} fax
 */
ServiceStoragePlanDatabaseTypes.prototype['fax'] = undefined;

/**
 * @member {module:model/ServiceStoragePlanDatabaseDocument} function
 */
ServiceStoragePlanDatabaseTypes.prototype['function'] = undefined;

/**
 * @member {module:model/ServiceStoragePlanDatabaseDocument} mailbox_message
 */
ServiceStoragePlanDatabaseTypes.prototype['mailbox_message'] = undefined;

/**
 * @member {module:model/ServiceStoragePlanDatabaseDocument} media
 */
ServiceStoragePlanDatabaseTypes.prototype['media'] = undefined;






export default ServiceStoragePlanDatabaseTypes;

