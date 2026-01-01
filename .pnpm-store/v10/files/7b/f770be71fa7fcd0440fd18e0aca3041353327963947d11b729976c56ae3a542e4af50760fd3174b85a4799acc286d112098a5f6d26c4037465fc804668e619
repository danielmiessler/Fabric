# Simple Statistics


A JavaScript implementation of descriptive, regression, and inference statistics.

[![Coverage Status](https://coveralls.io/repos/github/simple-statistics/simple-statistics/badge.svg)](https://coveralls.io/github/simple-statistics/simple-statistics)
[![npm version](https://badge.fury.io/js/simple-statistics.svg)](http://badge.fury.io/js/simple-statistics)

Implemented in literate JavaScript with no dependencies, designed to work
in all modern browsers (including IE) as well as in [node.js](https://nodejs.org/).

* :green_book: [API Documentation](https://simple-statistics.github.io/)
* :chart_with_upwards_trend: [Benchmarks](./benchmarks/)
* :kissing: [A list of other statistics libraries](./SEEALSO.md)

## Installation

* **I'm using Node.js, Webpack, Browserify, Rollup, or another module bundler,
  and install packages from npm.**
  * First, install the `simple-statistics` module, using `npm install simple-statistics`,
    then include the code with require or import:
  * **I use the `require` function to use modules in my project. (most likely)**
    * When you use `require`, you have the freedom to assign the module to any
      variable name you want, but you need to specify the module's name exactly:
      in this case, 'simple-statistics'. The `require` method returns an object
      with all of the module's methods attached to it.<br /> <pre>var ss = require('simple-statistics')</pre>
  * **I use `import` to use modules in my project. I'm probably using Babel, `@std/esm`, Webpack, or Rollup.**
    * Import all functions under the ss object: <pre>import * as ss from 'simple-statistics'</pre>
      Include a specific named export: <pre>import {min} from 'simple-statistics'</pre>
      Simple statistics has _only_ named exports for ES6.
* **I'm using Deno.**
  * Simple Statistics is published for Deno: https://deno.land/x/simplestatistics
* **I'm not using a module bundler. I'm writing a web page, and want to include
  simple-statistics using a script tag.**
  * **I want to support all browsers**
    * When you use simple-statistics from a script tag, you don't get to choose
      the variable name it is assigned to: simple-statistics will always become
      available globally as the variable `ss`. You can reassign this variable to
      another name if you want to, but doing so is optional.
      ```HTML
      <script src='https://unpkg.com/simple-statistics@7.8.8/dist/simple-statistics.min.js'>
      </script>
      ```
  * **I want to use ES6 modules in a browser and I'm [willing to only support new browsers](https://caniuse.com/#feat=es6-module) to do it**
    * This module works great with the [`?module`](https://unpkg.com/#/query-parameters) query parameter of unpkg. If you
      specify `type='module'` in your script tag, you'll be able to import simple-statistics
      directly - through `index.js` and with true [ES6 import syntax and behavior](http://exploringjs.com/es6/ch_modules.html).
      ```js
      <script type='module'>
      import {min} from "https://unpkg.com/simple-statistics@7.8.8/index.js?module"
      console.log(min([1, 2, 3]))
      </script>
      ```
      This feature is still experimental in unpkg and very bleeding-edge.
