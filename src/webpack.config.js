const path = require("path");

module.exports = {
  mode: "development",
  entry: "./static/scripts/app.js",
  output: {
    filename: "app.js",
    path: path.resolve(__dirname, "static/view/dist"),
  },
  resolve: {
    alias: {
      NodeModules: path.resolve(__dirname, 'node_modules/'),
    },
  },
};