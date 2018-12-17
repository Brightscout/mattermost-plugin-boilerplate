
/**
 * Returns the base url of the plugin
 * installation.
 *
 * @return {string} instance base URL
 */
function getBaseURL() {
    const url = new URL(window.location.href);
    return `${url.protocol}//${url.host}`;
}

function popupCenter(url, title, w, h) {
    // Fixes dual-screen position                            Most browsers       Firefox
    const dualScreenLeft = window.screenLeft === undefined ? window.screenX : window.screenLeft; // eslint-disable-line no-undefined
    const dualScreenTop = window.screenTop === undefined ? window.screenY : window.screenTop; // eslint-disable-line no-undefined

    let width;
    let height;

    if (window.innerWidth) {
        width = window.innerWidth;
    } else {
        width = document.documentElement.clientWidth ? document.documentElement.clientWidth : screen.width;
    }

    if (window.innerHeight) {
        height = window.innerHeight;
    } else {
        height = document.documentElement.clientHeight ? document.documentElement.clientHeight : screen.height;
    }

    const left = ((width / 2) - (w / 2)) + dualScreenLeft;
    const top = ((height / 2) - (h / 2)) + dualScreenTop;
    const newWindow = window.open(url, title, 'scrollbars=yes, width=' + w + ', height=' + h + ', top=' + top + ', left=' + left);

    if (newWindow != null) {
        newWindow.focus();
    }

    return newWindow;
}

export default {
    getBaseURL,
    popupCenter,
};
