import React from 'react';

import {ChannelHeaderButtonIcon} from 'components/icons';

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

window.registerPlugin('boilerplate', new PluginClass());
