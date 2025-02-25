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
import ServiceCallRecordingSource from './ServiceCallRecordingSource';

/**
 * The ServiceCallRecordingSettings model module.
 * @module model/ServiceCallRecordingSettings
 * @version 1.1
 */
class ServiceCallRecordingSettings {
    /**
     * Constructs a new <code>ServiceCallRecordingSettings</code>.
     * @alias module:model/ServiceCallRecordingSettings
     */
    constructor() { 
        
        ServiceCallRecordingSettings.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ServiceCallRecordingSettings</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ServiceCallRecordingSettings} obj Optional instance to populate.
     * @return {module:model/ServiceCallRecordingSettings} The populated <code>ServiceCallRecordingSettings</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ServiceCallRecordingSettings();

            if (data.hasOwnProperty('any')) {
                obj['any'] = ServiceCallRecordingSource.constructFromObject(data['any']);
            }
            if (data.hasOwnProperty('inbound')) {
                obj['inbound'] = ServiceCallRecordingSource.constructFromObject(data['inbound']);
            }
            if (data.hasOwnProperty('outbound')) {
                obj['outbound'] = ServiceCallRecordingSource.constructFromObject(data['outbound']);
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>ServiceCallRecordingSettings</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>ServiceCallRecordingSettings</code>.
     */
    static validateJSON(data) {
        // validate the optional field `any`
        if (data['any']) { // data not null
          ServiceCallRecordingSource.validateJSON(data['any']);
        }
        // validate the optional field `inbound`
        if (data['inbound']) { // data not null
          ServiceCallRecordingSource.validateJSON(data['inbound']);
        }
        // validate the optional field `outbound`
        if (data['outbound']) { // data not null
          ServiceCallRecordingSource.validateJSON(data['outbound']);
        }

        return true;
    }


}



/**
 * @member {module:model/ServiceCallRecordingSource} any
 */
ServiceCallRecordingSettings.prototype['any'] = undefined;

/**
 * @member {module:model/ServiceCallRecordingSource} inbound
 */
ServiceCallRecordingSettings.prototype['inbound'] = undefined;

/**
 * @member {module:model/ServiceCallRecordingSource} outbound
 */
ServiceCallRecordingSettings.prototype['outbound'] = undefined;






export default ServiceCallRecordingSettings;

