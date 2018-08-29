var path = require('path');

module.exports = {
    entry: [
        './src/index.jsx'
    ],
    resolve: {
        modules: [
            'src',
            'node_modules'
        ],
        extensions: ['*', '.js', '.jsx']
    },
    module: {
        rules: [
            {
                test: /\.(js|jsx)$/,
                exclude: /node_modules/,
                use: {
                    loader: 'babel-loader',
                    options: {
                        presets: ['env', 'react']
                    }
                }
            }
        ]
    },
    externals: {
        react: 'React'
    },
    output: {
        path: path.join(__dirname, '/dist'),
        publicPath: '/',
        filename: 'main.js'
    }
};