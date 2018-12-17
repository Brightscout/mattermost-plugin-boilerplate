const path = require('path');

module.exports = {
    entry: [
        './src/index.jsx',
    ],
    resolve: {
        modules: [
            'src',
            'node_modules',
        ],
        extensions: ['*', '.js', '.jsx'],
    },
    module: {
        rules: [
            {
                test: /\.(js|jsx)$/,
                exclude: /node_modules/,
                use: {
                    loader: 'babel-loader',
                    options: {
                        presets: [
                            'react',
                            'env',
                            'stage-0',
                        ],
                        plugins: [
                            'transform-runtime',
                        ],
                    },
                },
            },
        ],
    },
    externals: {
        react: 'React',
        redux: 'Redux',
        'prop-types': 'PropTypes',
        'post-utils': 'PostUtils',
        'react-bootstrap': 'ReactBootstrap',
        'react-redux': 'ReactRedux',
    },
    output: {
        path: path.join(__dirname, '/dist'),
        publicPath: '/',
        filename: 'main.js',
    },
};
