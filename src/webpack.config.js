const path = require("path");

module.exports = {
  mode: "development",
  entry: "./static/src/app.js",
  output: {
    filename: "app.js",
    path: path.resolve(__dirname, "static/view/dist"),
  },
  resolve: {
    alias: {
      NodeModules: path.resolve(__dirname, 'node_modules/'),
    },
  },
  module: {
    rules: [
      {
        test: /\.css$/i,
        include: path.resolve(__dirname, 'static'),
        use: ['style-loader', 'css-loader', 'postcss-loader'],
      },
    ],
  },
};