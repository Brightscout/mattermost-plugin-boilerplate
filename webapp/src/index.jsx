import React from 'react';

import {ChannelHeaderButtonIcon} from 'components/icons';

import Constants from './constants';

//
// Define the plugin class that will register
// our plugin components.
//
class PluginClass {
    initialize(registry) {
        registry.registerChannelHeaderButtonAction(

            // icon - JSX element to use as the button's icon
            <ChannelHeaderButtonIcon/>,

            // action - a function called when the button is clicked, passed the channel and channel member as arguments
            // null,
            () => {
                console.log('Hello World!'); // eslint-disable-line no-console
            },

            // dropdown_text - string or JSX element shown for the dropdown button description
            'Hello World',
        );
    }
}

//
// To register your plugin you must expose it
// on window.
//
window.registerPlugin(Constants.PLUGIN_NAME, new PluginClass());
