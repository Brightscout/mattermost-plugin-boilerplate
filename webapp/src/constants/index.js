import Utils from '../utils';

import {PLUGIN_NAME} from './manifest';
import SVGS from './svgs';

//
// Define our URL constants
//
const URL_BASE = `${Utils.getBaseURL()}/plugins/${PLUGIN_NAME}`;

//
// Export the constants
//
export default {
    PLUGIN_NAME,
    URL_BASE,
    SVGS,
};
