const { babel } = require("@rollup/plugin-babel");
const nodeResolve = require("@rollup/plugin-node-resolve");
const commonjs = require("@rollup/plugin-commonjs");
const replace = require("@rollup/plugin-replace");

const globals = require("rollup-plugin-node-globals");
const builtins = require("rollup-plugin-node-builtins");
const plugin_async = require("rollup-plugin-async");
const postcss = require("rollup-plugin-postcss");
const svg = require("rollup-plugin-svg");

const babelOptions = {
  babelrc: false,
  presets: ["@babel/preset-react"],
  babelHelpers: "bundled",
};

module.exports = [
  {
    input: "./embed/index.js",
    output: {
      file: "../react-embed/index.js",
      format: "esm",
    },
    external: [],
    plugins: [
      plugin_async(),
      babel(babelOptions),
      nodeResolve({ preferBuiltins: true }),
      commonjs({ ignoreDynamicRequires: false }),
      postcss({ minimize: true }),
      svg({ base64: true, extract: true }),
      globals(),
      builtins(),
      replace({
        preventAssignment: true,
        "process.env.NODE_ENV": JSON.stringify("production"),
        "process.env.NODE_DEBUG": JSON.stringify(""),
      }),
    ],
  },
];
