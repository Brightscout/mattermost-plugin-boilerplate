import React from 'react';
import {ChannelHeaderButtonIcon} from 'components/icons';

class HelloWorldPlugin {
    initialize(registry, store) {
        registry.registerChannelHeaderButtonAction(
            // icon - JSX element to use as the button's icon
            <ChannelHeaderButtonIcon/>,
            // action - a function called when the button is clicked, passed the channel and channel member as arguments
            // null,
            () => {
                console.log('Hello World!');
            },
            // dropdown_text - string or JSX element shown for the dropdown button description
            'Hello World',
        );
    }
}

window.registerPlugin('boilerplate', new HelloWorldPlugin());