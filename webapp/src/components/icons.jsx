import React from 'react';

import Constants from '../constants';

export class ChannelHeaderButtonIcon extends React.PureComponent {
    render() {
        return (
            <span
                dangerouslySetInnerHTML={{
                    __html: Constants.SVGS.CHANNEL_HEADER_ICON,
                }}
            />
        );
    }
}
