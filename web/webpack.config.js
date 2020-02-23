const path = require("path");
const HtmlWebpackPlugin = require("html-webpack-plugin");

module.exports = {

    entry: "./src/js/index.jsx",

    output: {
        path: path.resolve(__dirname, "./dist"),
        filename: "bundle.js"
    },

    plugins: [
        new HtmlWebpackPlugin({
            template: "./src/html/index.html"
        }),
    ],

    module: {
        rules: [
            {
                test: /\.js(x?)$/,
                exclude: /node_modules/,
                use: [
                    {
                        loader: "babel-loader"
                    }
                ]
            },
            {
                test: /\.css$/,
                use: [
                    {
                        loader: "style-loader"
                    },
                    {
                        loader: "css-loader"
                    }]
            },
            {
                test: /\.html$/,
                use: [
                    {
                        loader: "html-loader"
                    }
                ]
            }
        ]
    }
};