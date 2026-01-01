/**
 * @fileoverview Main package entrypoint.
 * @author Nicholas C. Zakas
 */

"use strict";

const { version } = require("../package.json");

//------------------------------------------------------------------------------
// Public Interface
//------------------------------------------------------------------------------

module.exports = {
    meta: {
        name: "@eslint/js",
        version
    },
    configs: {
        all: require("./configs/eslint-all"),
        recommended: require("./configs/eslint-recommended")
    }
};
