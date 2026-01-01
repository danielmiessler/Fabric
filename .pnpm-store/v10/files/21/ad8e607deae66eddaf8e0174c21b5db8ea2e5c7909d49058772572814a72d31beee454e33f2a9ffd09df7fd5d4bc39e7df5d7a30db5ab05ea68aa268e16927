[string.js](http://stringjs.com)
=========

[![build status](https://secure.travis-ci.org/jprichardson/string.js.png)](http://travis-ci.org/jprichardson/string.js)

`string.js`, or simply `S` is a lightweight (**< 5 kb** minified and gzipped) JavaScript library for the browser or for Node.js that provides extra String methods. Originally, it modified the String prototype. But I quickly learned that in JavaScript, this is considered poor practice.



Why?
----

Personally, I prefer the cleanliness of the way code looks when it appears to be native methods. i.e. when you modify native JavaScript prototypes. However, if any app dependency required `string.js`, then the app's string prototype would be modified in every module. This could be troublesome. So I settled on creating a wrapper a la jQuery style. For those of you prototype hatin' fools, there is the method `extendPrototype()`.

Here's a list of alternative frameworks:

* [Prototype Framework's String library](http://prototypejs.org/api/string)
* [Uize.String](http://www.uize.com/reference/Uize.String.html)
* [Google Closure's String](http://closure-library.googlecode.com/svn/docs/namespace_goog_string.html)
* [Underscore.string](http://epeli.github.com/underscore.string/)
* [Sugar.js](http://sugarjs.com)
* [php.js](http://phpjs.org/pages/home)

Why wasn't I happy with any of them? They are all static methods that don't seem to support chaining in a clean way 'OR' they have an odd dependency. Sugar is the notable exception.



Installation
------------

    npm install --save string



Experiment with String.js Now
-----------------------------

Assuming you're on http://stringjs.com, just simply open up the Webkit inspector in either Chrome or Safari, or the web console in Firefox and you'll notice that `string.js` is included in this page so that you can start experimenting with it right away.



Usage
-----

### Node.js

```javascript
var S = require('string');
```

Originally, I was using `$s` but glancing over the code, it was easy to confuse `$s` for string.js with `$` for jQuery. Feel free to use the most convenient variable for you.

### Browsers

```html
<!-- HTML5 -->
<script src="https://raw.github.com/jprichardson/string.js/master/lib/string.min.js"></script>

<!-- Note that in the mime type for Javascript is now officially 'application/javascript'. If you
set the type to application/javascript in IE browsers, your Javscript will fail. Just don't set a
type via the script tag and set the mime type from your server. Most browsers look at the server mime
type anyway -->

<!-- For HTML4/IE -->
<script type="text/javascript" src="https://raw.github.com/jprichardson/lib/string.js/master/string.min.js"></script>
```

A global variable `window.S` or simply `S` is created.


### AMD Support

It now [has AMD support](https://github.com/jprichardson/string.js/pull/20). See [require.js](http://requirejs.org/) on how to use with AMD modules.


### Both

```javascript
var doesIt = S('my cool string').left(2).endsWith('y'); //true
```

Access the wrapped string using `s` variable or `toString()`

```javascript
var name = S('Your name is JP').right(2).s; //'JP'
```

is the same as…

```javascript
var name = S('Your name is JP').right(2).toString(); //'JP'
```

Still like the clean look of calling these methods directly on native Strings? No problem. Call `extendPrototype()`. Make sure to not call this at the module level, at it'll effect the entire application lifecycle. You should really only use this at the method level. The one exception being if your application will not be a dependency of another application.

```javascript
S.extendPrototype();
var name = 'Your name is JP'.right(2); //'JP'
S.restorePrototype(); //be a good citizen and clean up
```


### Browser Compatiblity

`string.js` has been designed to be compatible with Node.js and with IE6+, Firefox 3+, Safari 2+, Chrome 3+. Please [click here][browsertest] to run the tests in your browser. Report any browser issues here: https://github.com/jprichardson/string.js/issues


Native JavaScript Methods
-------------------------

`string.js` imports all of the native JavaScript methods. This is for convenience. The only difference is that the imported methods return `string.js` objects instead of native JavaScript strings. The one exception to this is the method `charAt(index)`. This is because `charAt()` only returns a string of length one. This is typically done for comparisons and a `string.js` object will have little to no value here.

All of the native methods support chaining with the `string.js` methods.

**Example:**

```javascript
var S = require('string');

var phrase = S('JavaScript is the best scripting language ever!');
var sub = 'best scripting';
var pos = phrase.indexOf(sub);
console.log(phrase.substr(pos, sub.length).truncate(8)); //best...
```


Methods
-------

See [test file][testfile] for more details.

I use the same nomenclature as Objective-C regarding methods. **+** means `static` or `class` method. **-** means `non-static` or `instance` method.

### - constructor(nativeJsString) ###

This creates a new `string.js` object. The parameter can be anything. The `toString()` method will be called on any objects. Some native objects are used in some functions such as `toCSV()`.

Example:

```javascript
S('hello').s //'hello'
S(['a,b']).s //"'a','b'"
S({hi: 'jp'}).s //[object Object]
```


### - between(left, right)

Extracts a string between `left` and `right` strings.

Example:

```javascript
S('<a>foobar</a>').between('<a>', '</a>').s; // 'foobar'
```

### - camelize()

Remove any underscores or dashes and convert a string into camel casing.

Example:

```javascript
S('data_rate').camelize().s; //'dataRate'
S('background-color').camelize().s; //'backgroundColor'
S('-moz-something').camelize().s; //'mozSomething'
S('_car_speed_').camelize().s; //'carSpeed'
S('yes_we_can').camelize().s; //'yesWeCan'
```


### - capitalize() ###

Capitalizes the first character of a string.

Example:

```javascript
S('jon').capitalize().s; //'Jon'
S('JP').capitalize().s; //'Jp'
```


### - chompLeft(prefix)

Removes `prefix` from start of string.

Example:

```javascript
S('foobar').chompLeft('foo').s; //'bar'
S('foobar').chompLeft('bar').s; //'foobar'
```


### - chompRight(suffix)

Removes `suffix` from end of string.

Example:

```javascript
S('foobar').chompRight('bar').s; //'foo'
S('foobar').chompRight('foo').s; //'foobar'
```


### - collapseWhitespace() ###

Converts all adjacent whitespace characters to a single space.

Example:

```javascript
var str = S('  String   \t libraries are   \n\n\t fun\n!  ').collapseWhitespace().s; //'String libraries are fun !'
```


### - contains(ss) ###

Returns true if the string contains `ss`.

Alias: `include()`

Example:

```javascript
S('JavaScript is one of the best languages!').contains('one'); //true
```


### - count(substring) ###

Returns the count of the number of occurrences of the substring.

Example:

```javascript
S('JP likes to program. JP does not play in the NBA.').count("JP")// 2
S('Does not exist.').count("Flying Spaghetti Monster") //0
S('Does not exist.').count("Bigfoot") //0
S('JavaScript is fun, therefore Node.js is fun').count("fun") //2
S('funfunfun').count("fun") //3
```


### - dasherize() ###

Returns a converted camel cased string into a string delimited by dashes.

Examples:

```javascript
S('dataRate').dasherize().s; //'data-rate'
S('CarSpeed').dasherize().s; //'-car-speed'
S('yesWeCan').dasherize().s; //'yes-we-can'
S('backgroundColor').dasherize().s; //'background-color'
```


### - decodeHTMLEntities() ###

Decodes HTML entities into their string representation.

```javascript
S('Ken Thompson &amp; Dennis Ritchie').decodeHTMLEntities().s; //'Ken Thompson & Dennis Ritchie'
S('3 &lt; 4').decodeHTMLEntities().s; //'3 < 4'
```


### - endsWith(ss) ###

Returns true if the string ends with `ss`.

Example:

```javascript
S("hello jon").endsWith('jon'); //true
```


### - escapeHTML() ###

Escapes the html.

Example:

```javascript
S('<div>hi</div>').escapeHTML().s; //&lt;div&gt;hi&lt;/div&gt;
```


### + extendPrototype() ###

Modifies `String.prototype` to have all of the methods found in string.js.

Example:

```javascript
S.extendPrototype();
```



### - ensureLeft(prefix)

Ensures string starts with `prefix`.

Example:

```javascript
S('subdir').ensureLeft('/').s; //'/subdir'
S('/subdir').ensureLeft('/').s; //'/subdir'
```


### - ensureRight(suffix)

Ensures string ends with `suffix`.

Example:

```javascript
S('dir').ensureRight('/').s; //'dir/'
S('dir/').ensureRight('/').s; //'dir/'
```

### - humanize() ###

Transforms the input into a human friendly form.

Example:

```javascript
S('the_humanize_string_method').humanize().s  //'The humanize string method'
S('ThehumanizeStringMethod').humanize().s //'Thehumanize string method'
S('the humanize string method').humanize().s  //'The humanize string method'
S('the humanize_id string method_id').humanize().s //'The humanize id string method'
S('the  humanize string method  ').humanize().s //'The humanize string method'
S('   capitalize dash-CamelCase_underscore trim  ').humanize().s //'Capitalize dash camel case underscore trim'
```

### - include(ss) ###

Returns true if the string contains the `ss`.

Alias: `contains()`

Example:

```javascript
S('JavaScript is one of the best languages!').include('one'); //true
```


### - isAlpha() ###

Return true if the string contains only letters.

Example:

```javascript
S("afaf").isAlpha(); //true
S('fdafaf3').isAlpha(); //false
S('dfdf--dfd').isAlpha(); //false
```


### - isAlphaNumeric() ###

Return true if the string contains only letters and numbers

Example:

```javascript
S("afaf35353afaf").isAlphaNumeric(); //true
S("FFFF99fff").isAlphaNumeric(); //true
S("99").isAlphaNumeric(); //true
S("afff").isAlphaNumeric(); //true
S("Infinity").isAlphaNumeric(); //true
S("-Infinity").isAlphaNumeric(); //false
S("-33").isAlphaNumeric(); //false
S("aaff..").isAlphaNumeric(); //false
```


### - isEmpty() ###

Return true if the string is solely composed of whitespace or is `null`/`undefined`.

Example:

```javascript
S(' ').isEmpty(); //true
S('\t\t\t    ').isEmpty(); //true
S('\n\n ').isEmpty(); //true
S('helo').isEmpty(); //false
S(null).isEmpty(); //true
S(undefined).isEmpty(); //true
```


### - isLower() ###

Return true if the character or string is lowercase

Example:

```javascript
S('a').isLower(); //true
S('z').isLower(); //true
S('B').isLower(); //false
S('hijp').isLower(); //true
S('hi jp').isLower(); //false
S('HelLO').isLower(); //false
```


### - isNumeric() ###

Return true if the string only contains digits

Example:

```javascript
S("3").isNumeric(); //true
S("34.22").isNumeric(); //false
S("-22.33").isNumeric(); //false
S("NaN").isNumeric(); //false
S("Infinity").isNumeric(); //false
S("-Infinity").isNumeric(); //false
S("JP").isNumeric(); //false
S("-5").isNumeric(); //false
S("000992424242").isNumeric(); //true
```


### - isUpper() ###

Returns true if the character or string is uppercase

Example:

```javascript
S('a').isUpper() //false
S('z').isUpper()  //false
S('B').isUpper() //true
S('HIJP').isUpper() //true
S('HI JP').isUpper() //false
S('HelLO').isUpper() //true
```


### - left(n) ###

Return the substring denoted by `n` positive left-most characters.

Example:

```javascript
S('My name is JP').left(2).s; //'My'
S('Hi').left(0).s; //''
S('My name is JP').left(-2).s; //'JP', same as right(2)
```


### - length ###

Property to return the length of the string object.

Example:

```javascript
S('hi').length; //2
```


### - pad(len, [char])

Pads the string in the center with specified character

Example:

```javascript
S('hello').pad(5).s //'hello'
S('hello').pad(10).s //'   hello  '
S('hey').pad(7).s //'  hey  '
S('hey').pad(5).s //' hey '
S('hey').pad(4).s //' hey'
S('hey').pad(7, '-').s//'--hey--'
```


### - padLeft(len, [char])

Left pads the string.

Example:

```javascript
S('hello').padLeft(5).s //'hello'
S('hello').padLeft(10).s //'     hello'
S('hello').padLeft(7).s //'  hello'
S('hello').padLeft(6).s //' hello'
S('hello').padLeft(10, '.').s //'.....hello'
```


### - padRight(len, [char])

Right pads the string.

Example:

```javascript
S('hello').padRight(5).s //'hello'
S('hello').padRight(10).s //'hello     '
S('hello').padRight(7).s //'hello  '
S('hello').padRight(6).s //'hello '
S('hello').padRight(10, '.').s //'hello.....'
```


### - parseCSV() ###

Parses a CSV line into an array.

**Arguments:**
- `delimiter`: The character that is separates or delimits fields. Default: `,`
- `qualifier`: The character that encloses fields. Default: `"`
- `escape`: The character that represents the escape character. Default: `\`
- `lineDelimiter`: The character that represents the end of a line. When a lineDelimiter is passed the result will be a multidimensional array. Default: `undefined`

Example:

```javascript
S("'a','b','c'").parseCSV(',', "'") //['a', 'b', 'c'])
S('"a","b","c"').parseCSV() // ['a', 'b', 'c'])
S('a,b,c').parseCSV(',', null)  //['a', 'b', 'c'])
S("'a,','b','c'").parseCSV(',', "'") //['a,', 'b', 'c'])
S('"a","b",4,"c"').parseCSV(',', null) //['"a"', '"b"', '4', '"c"'])
S('"a","b","4","c"').parseCSV() //['a', 'b', '4', 'c'])
S('"a","b",       "4","c"').parseCSV() //['a', 'b', '4', 'c'])
S('"a","b",       4,"c"').parseCSV(",", null) //[ '"a"', '"b"', '       4', '"c"' ])
S('"a","b\\"","d","c"').parseCSV() //['a', 'b"', 'd', 'c'])
S('"a","b\\"","d","c"').parseCSV() //['a', 'b"', 'd', 'c'])
S('"a\na","b","c"\n"a", """b\nb", "a"').parseCSV(',', '"', '"', '\n')) // [ [ 'a\na', 'b', 'c' ], [ 'a', '"b\nb', 'a' ] ]
```

### - repeat(n) ###

Returns a string repeated `n` times.

Alias: `times()`

Example:

```javascript
S(' ').repeat(5).s; //'     '
S('*').repeat(3).s; //'***'
```


### - replaceAll(ss, newstr) ###

Return the new string with all occurrences of `ss` replaced with `newstr`.

Example:

```javascript
S(' does IT work? ').replaceAll(' ', '_').s; //'_does_IT_work?_'
S('Yes it does!').replaceAll(' ', '').s; //'Yesitdoes!'
```


### + restorePrototype() ###

Restore the original String prototype. Typically used in conjunction with `extendPrototype()`.

Example:

```javascript
S.restorePrototype();
```


### - right(n) ###

Return the substring denoted by `n` positive right-most characters.

Example:

```javascript
S('I AM CRAZY').right(2).s; //'ZY'
S('Does it work?  ').right(4).s; //'k?  '
S('Hi').right(0).s; //''
S('My name is JP').right(-2).s; //'My', same as left(2)
```


### - s ###

Alias: `toString()`

The encapsulated native string representation of an `S` object.

Example:

```javascript
S('my name is JP.').capitalize().s; //My name is JP.
var a = "Hello " + S('joe!'); //a = "Hello joe!"
S("Hello").toString() === S("Hello").s; //true
```


### - slugify() ###

Converts the text into a valid url slug.

```javascript
S('Global Thermonuclear Warfare').slugify().s // 'global-thermonuclear-warfare'
```


### - startsWith(prefix)   ###

Return true if the string starts with `prefix`.

Example:

```javascript
S("JP is a software engineer").startsWith("JP"); //true
S('wants to change the world').startsWith("politicians"); //false
```


### - stripPunctuation()

Strip all of the punctuation.

Example:

```javascript
S('My, st[ring] *full* of %punct)').stripPunctuation().s; //My string full of punct
```



### - stripTags([tag1],[tag2],...) ###

Strip all of the HTML tags or tags specified by the parameters.

Example:

```javascript
S('<p>just <b>some</b> text</p>').stripTags().s //'just some text'
S('<p>just <b>some</b> text</p>').stripTags('p').s //'just <b>some</b> text'
```


### - template(values, [open], [close])

Takes a string and interpolates the values. Defaults to `{{` and `}}` for Mustache compatible templates. However, you can change this default by modifying `S.TMPL_OPEN` and `S.TMPL_CLOSE`.

Example:

```js
var str = "Hello {{name}}! How are you doing during the year of {{date-year}}?"
var values = {name: 'JP', 'date-year': 2013}
console.log(S(str).template(values).s) //'Hello JP! How are you doing during the year of 2013?'

str = "Hello #{name}! How are you doing during the year of #{date-year}?"
console.log(S(str).template(values, '#{', '}').s) //'Hello JP! How are you doing during the year of 2013?'

S.TMPL_OPEN = '{'
S.TMPL_CLOSE = '}'
str = "Hello {name}! How are you doing during the year of {date-year}?"
console.log(S(str).template(values).s) //'Hello JP! How are you doing during the year of 2013?'
```


### - times(n) ###

Returns a string repeated `n` times.

Alias: `repeat()`

Example:

```javascript
S(' ').times(5).s //'     '
S('*').times(3).s //'***'
```


### - toBoolean() / toBool()

Converts a a logical truth string to boolean. That is: `true`, `1`, `'true'`, `'on'`, or `'yes'`.

JavaScript Note: You can easily convert truthy values to `booleans` by prefixing them with `!!`. e.g.
`!!'hi' === true` or `!!'' === false` or `!!{} === true`.

Example:

```javascript
S('true').toBoolean() //true
S('false').toBoolean() //false
S('hello').toBoolean() //false
S(true).toBoolean() //true
S('on').toBoolean() //true
S('yes').toBoolean() //true
S('TRUE').toBoolean() //true
S('TrUe').toBoolean() //true
S('YES').toBoolean() //true
S('ON').toBoolean() //true
S('').toBoolean() //false
S(undefined).toBoolean() //false
S('undefined').toBoolean() //false
S(null).toBoolean() //false
S(false).toBoolean() //false
S({}).toBoolean() //false
S(1).toBoolean() //true
S(-1).toBoolean() //false
S(0).toBoolean() //false
```



### - toCSV(options) ###

Converts an array or object to a CSV line.

You can either optionally pass in two string arguments or pass in a configuration object.

**String Arguments:**
- `delimiter`: The character that is separates or delimits fields. Default: `,`
- `qualifier`: The character that encloses fields. Default: `"`


**Object Configuration:**
- `delimiter`: The character that is separates or delimits fields. Default: `,`
- `qualifier`: The character that encloses fields. Default: `"`
- `escape`: The character that escapes any incline `qualifier` characters. Default: `\`, in JS this is `\\`
- `encloseNumbers`: Enclose number objects with the `qualifier` chracter. Default: `true`
- `keys`: If the input isn't an array, but an object, then if this is set to true, the keys will be output to the CSV line, otherwise it's the object's values. Default: `false`.

Example:

```javascript
S(['a', 'b', 'c']).toCSV().s //'"a","b","c"'
S(['a', 'b', 'c']).toCSV(':').s //'"a":"b":"c"'
S(['a', 'b', 'c']).toCSV(':', null).s //'a:b:c')
S(['a', 'b', 'c']).toCSV('*', "'").s //"'a'*'b'*'c'"
S(['a"', 'b', 4, 'c']).toCSV({delimiter: ',', qualifier: '"', escape: '\\',  encloseNumbers: false}).s //'"a\\"","b",4,"c"'
S({firstName: 'JP', lastName: 'Richardson'}).toCSV({keys: true}).s //'"firstName","lastName"'
S({firstName: 'JP', lastName: 'Richardson'}).toCSV().s //'"JP","Richardson"'
```


### - toFloat([precision]) ###

Return the float value, wraps parseFloat.

Example:

```javascript
S('5').toFloat() // 5
S('5.3').toFloat()  //5.3
S(5.3).toFloat()  //5.3
S('-10').toFloat()  //-10
S('55.3 adfafaf').toFloat() // 55.3
S('afff 44').toFloat()  //NaN
S(3.45522222333232).toFloat(2) // 3.46
```


### - toInt() / toInteger() ###

Return the number value in integer form. Wrapper for `parseInt()`. Can also parse hex values.

Example:

```javascript
S('5').toInt(); //5
S('5.3').toInt(); //5;
S(5.3).toInt(); //5;
S('-10').toInt(); //-10
S('55 adfafaf').toInt(); //55
S('afff 44').toInt(); //NaN
S('0xff').toInt() //255
```



### - toString() ###

Alias: `s`

Return the string representation of an `S` object. Not really necessary to use. However, JS engines will look at an object and display its `toString()` result.

Example:

```javascript
S('my name is JP.').capitalize().toString(); //My name is JP.
var a = "Hello " + S('joe!'); //a = "Hello joe!"
S("Hello").toString() === S("Hello").s; //true
```


### - trim() ###

Return the string with leading and trailing whitespace removed. Reverts to native `trim()` if it exists.

Example:

```javascript
S('hello ').trim().s; //'hello'
S(' hello ').trim().s; //'hello'
S('\nhello').trim().s; //'hello'
S('\nhello\r\n').trim().s; //'hello'
S('\thello\t').trim().s; //'hello'
```


### - trimLeft() ###

Return the string with leading and whitespace removed

Example:

```javascript
S('  How are you?').trimLeft().s; //'How are you?';
```


### - trimRight() ###

Return the string with trailing whitespace removed.

Example:

```javascript
S('How are you?   ').trimRight().s; //'How are you?';
```


### - truncate(length, [chars]) ###

Truncates the string, accounting for word placement and character count.

Example:

```javascript
S('this is some long text').truncate(3).s //'...'
S('this is some long text').truncate(7).s //'this is...'
S('this is some long text').truncate(11).s //'this is...'
S('this is some long text').truncate(12).s //'this is some...'
S('this is some long text').truncate(11).s //'this is...'
S('this is some long text').truncate(14, ' read more').s //'this is some read more'
```



### - underscore()

Returns converted camel cased string into a string delimited by underscores.

Example:

```javascript
S('dataRate').underscore().s; //'data_rate'
S('CarSpeed').underscore().s; //'_car_speed'
S('yesWeCan').underscore().s; //'yes_we_can'
```


### - unescapeHTML() ###

Unescapes the html.

Example:

```javascript
S('&lt;div&gt;hi&lt;/div&gt;').unescapeHTML().s; //<div>hi</div>
```



### + VERSION ###

Returns native JavaScript string containing the version of `string.js`.

Example:

```javascript
S.VERSION; //1.0.0
```


Quirks
------

`decodeHtmlEntities()` converts `&nbsp;` to **0xa0** (160) and not **0x10** (20). Most browsers consider 0xa0 to be whitespace characters, Internet Explorer does not despite it being part of the ECMA standard. Google Closure does a good job of normalizing this behavior. This may need to fixed in `string.js` at some point in time.



Testing
-------

### Node.js

Install the dev dependencies:

    $ npm install string --development

Then navigate to the installed directory:

    $ cd node_modules/string/

Run test package:

    $ mocha test



### Browser ###

[Click here to run the tests in your web browser.][browsertest]



Credits
-------

I have looked at the code by the creators in the libraries mentioned in **Motivation**. As noted in the source code, I've specifically used code from Google Closure (Google Inc), Underscore String [Esa-Matti Suuronen](http://esa-matti.suuronen.org/), and php.js (http://phpjs.org/authors/index), [Substack](https://github.com/substack/node-ent) and [TJ Holowaychuk](https://github.com/component/pad).



Contributions
-------------

If you contribute to this library, just modify `string.js`, `string.test.js`, and update `README.md`. I'll update the website docs and generate the new `string.min.js`, changelog and version.


### Contributors

(You can add your name, or I'll add it if you forget)

- [JP Richardson](https://github.com/jprichardson)
- [Leonardo Otero](https://github.com/oteroleonardo)
- [Jordan Scales](https://github.com/prezjordan)
- [Eduardo de Matos](https://github.com/eduardo-matos)
- [Christian Maughan Tegnér](https://github.com/CMTegner)
- [Mario Gutierrez](https://github.com/mgutz)
- [Sean O'Dell](https://github.com/seanodell)
- [Tim de Koning](https://github.com/Reggino)
- `<your name here>`


License
-------

Licensed under MIT.

Copyright (C) 2012-2013 JP Richardson <jprichardson@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.




[testfile]: https://github.com/jprichardson/string.js/blob/master/test/string.test.js
[browsertest]: http://stringjs.com/browser.test.html

[aboutjp]: http://about.me/jprichardson
[twitter]: http://twitter.com/jprichardson
[procbits]: http://procbits.com



