(function() {
  'use strict';

  var S = null;

  if (typeof module !== 'undefined'  && typeof module.exports !== 'undefined')
    S = require('../lib/string');
  else {
    S = window.S;
  }

  function T(v) { if (!v) { throw new Error('Should be true.'); } };
  function F(v) { if (v) { throw new Error('Should be false.'); } };
  function EQ(v1, v2) {
    if (typeof require != 'undefined' && typeof process != 'undefined') //node
      require('assert').equal(v1, v2)
    else
      T (v1 === v2)
  }

  function ARY_EQ(a1, a2) {
    EQ (a1.length, a2.length)
    for (var i = 0; i < a1.length; ++i)
      EQ (a1[i], a2[i])
  }


  /*if (typeof window !== "undefined" && window !== null) {
    S = window.S;
  } else {
    S = require('../lib/string');
  }*/

  describe('string.js', function() {

    describe('- constructor', function() {
      it('should should set the internal "s" property', function() {
        T (S('helo').s === 'helo')
        T (S(5).s === '5')
        T (S(new Date(2012, 1, 1)).s.indexOf('2012') != -1)
        T (S(new RegExp()).s.substr(0,1) === '/')
        T (S({}).s === '[object Object]')
        T (S(null).s === null)
        T (S(undefined).s === undefined)
      })
    })

    describe('- between(left, right)', function() {
      it('should extract string between `left` and `right`', function() {
        T (S('<a>foo</a>').between('<a>', '</a>').s === 'foo')
        T (S('<a>foo</a></a>').between('<a>', '</a>').s === 'foo')
        T (S('<a><a>foo</a></a>').between('<a>', '</a>').s === '<a>foo')
        T (S('<a>foo').between('<a>', '</a>').s === '')
      })
    })

    describe('- camelize()', function() {
      it('should remove any underscores or dashes and convert a string into camel casing', function() {
        T (S('data_rate').camelize().s === 'dataRate');
        T (S('background-color').camelize().s === 'backgroundColor');
        T (S('-moz-something').camelize().s === 'MozSomething');
        T (S('_car_speed_').camelize().s === 'CarSpeed');
        T (S('yes_we_can').camelize().s === 'yesWeCan');
      })
    })

    describe('- capitalize()', function() {
      it('should capitalize the string', function() {
        T (S('jon').capitalize().s === 'Jon');
        T (S('JP').capitalize().s === 'Jp');
      })
    })

    describe('- charAt(index)', function() {
      it('should return a native JavaScript string with the character at the specified position', function() {
        T (S('hi').charAt(1) === 'i')
      })
    })

    describe('- chompLeft(prefix)', function() {
      it('should remove `prefix` from start of string', function() {
        T (S('foobar').chompLeft('foo').s === 'bar')
        T (S('foobar').chompLeft('bar').s === 'foobar')
        T (S('').chompLeft('foo').s === '')
        T (S('').chompLeft('').s === '')
      })
    })

    describe('- chompRight(suffix)', function() {
      it('should remove `suffix` from end of string', function() {
        T (S('foobar').chompRight('foo').s === 'foobar')
        T (S('foobar').chompRight('bar').s === 'foo')
        T (S('').chompRight('foo').s === '')
        T (S('').chompRight('').s === '')
      })
    })

    describe('- collapseWhitespace()', function() {
      it('should convert all adjacent whitespace characters to a single space and trim the ends', function() {
        T (S('  Strings   \t are   \n\n\t fun\n!  ').collapseWhitespace().s === 'Strings are fun !');
      })
    })

    describe('- contains(substring)', function() {
      it('should return true if the string contains the specified input string', function() {
        T (S('JavaScript is one of the best languages!').contains('one'));
        F (S('What do you think?').contains('YES!'));
      })
    })

    describe('- count(substring)', function() {
      it('should return the count of all substrings', function() {
        EQ (S('JP likes to program. JP does not play in the NBA.').count("JP"), 2)
        EQ (S('Does not exist.').count("Flying Spaghetti Monster"), 0)
        EQ (S('Does not exist.').count("Bigfoot"), 0)
        EQ (S('JavaScript is fun, therefore Node.js is fun').count("fun"), 2)
        EQ (S('funfunfun').count("fun"), 3)
      })
    })

    describe('- dasherize()', function() {
      it('should convert a camel cased string into a string delimited by dashes', function() {
        T (S('dataRate').dasherize().s === 'data-rate');
        T (S('CarSpeed').dasherize().s === '-car-speed');
        T (S('yesWeCan').dasherize().s === 'yes-we-can');
        T (S('backgroundColor').dasherize().s === 'background-color');
      })
    })

    describe('- decodeHTMLEntities()', function() {
      it('should decode HTML entities into their proper string representation', function() {
        EQ (S('Ken Thompson &amp; Dennis Ritchie').decodeHTMLEntities().s, 'Ken Thompson & Dennis Ritchie');
        EQ (S('3 &lt; 4').decodeHTMLEntities().s, '3 < 4');
        EQ (S('http:&#47;&#47;').decodeHTMLEntities().s, 'http://')
      })
    })

    describe('- endsWith(suffix)', function() {
      it("should return true if the string ends with the input string", function() {
        T (S("hello jon").endsWith('jon'));
        F (S('ffffaaa').endsWith('jon'));
        T (S("").endsWith(''));
        T (S("hi").endsWith(''));
        T (S("hi").endsWith('hi'));
      })
    })

    describe('- ensureLeft(prefix)', function() {
      it('should prepend `prefix` if string does not start with prefix', function() {
        T (S('foobar').ensureLeft('foo').s === 'foobar')
        T (S('bar').ensureLeft('foo').s === 'foobar')
        T (S('').ensureLeft('foo').s === 'foo')
        T (S('').ensureLeft('').s === '')
      })
    })

    describe('- ensureRight(suffix)', function() {
      it('should append `suffix` if string does not end with suffix', function() {
        T (S('foobar').ensureRight('bar').s === 'foobar')
        T (S('foo').ensureRight('bar').s === 'foobar')
        T (S('').ensureRight('foo').s === 'foo')
        T (S('').ensureRight('').s === '')
      })
    })

    describe('- escapeHTML()', function() {
      it('should escape the html', function() {
        T (S('<div>Blah & "blah" & \'blah\'</div>').escapeHTML().s ===
             '&lt;div&gt;Blah &amp; &quot;blah&quot; &amp; &apos;blah&apos;&lt;/div&gt;');
        T (S('&lt;').escapeHTML().s === '&amp;lt;');
      })
    })

     describe('+ extendPrototype()', function() {
      it('should extend the String prototype with the extra methods', function() {
        S.extendPrototype();
        T (" hello!".endsWith('!'));
        S.restorePrototype();
      })
    })

    describe('- humanize()', function() {
      it('should humanize the string', function() {
        EQ (S('the_humanize_string_method').humanize().s, 'The humanize string method')
        EQ (S('ThehumanizeStringMethod').humanize().s, 'Thehumanize string method')
        EQ (S('the humanize string method').humanize().s, 'The humanize string method')
        EQ (S('the humanize_id string method_id').humanize().s, 'The humanize id string method')
        EQ (S('the  humanize string method  ').humanize().s, 'The humanize string method')
        EQ (S('   capitalize dash-CamelCase_underscore trim  ').humanize().s, 'Capitalize dash camel case underscore trim')
        EQ (S(123).humanize().s, '123')
        EQ (S('').humanize().s, '')
        EQ (S(null).humanize().s, '')
        EQ (S(undefined).humanize().s, '')
      })
    })

    describe('- include(substring)', function() {
      it('should return true if the string contains the specified input string', function() {
        T (S('JavaScript is one of the best languages!').include('one'));
        F (S('What do you think?').include('YES!'));
      })
    })

    describe('- isAlpha()', function() {
      it("should return true if the string contains only letters", function() {
        T (S("afaf").isAlpha());
        T (S("FJslfjkasfs").isAlpha());
        T (S("áéúóúÁÉÍÓÚãõÃÕàèìòùÀÈÌÒÙâêîôûÂÊÎÔÛäëïöüÄËÏÖÜçÇ").isAlpha());
        F (S("adflj43faljsdf").isAlpha());
        F (S("33").isAlpha());
        F (S("TT....TTTafafetstYY").isAlpha());
        F (S("-áéúóúÁÉÍÓÚãõÃÕàèìòùÀÈÌÒÙâêîôûÂÊÎÔÛäëïöüÄËÏÖÜçÇ").isAlpha());
      })
    })

    describe('- isAlphaNumeric()', function() {
      it("should return true if the string contains only letters and digits", function() {
        T (S("afaf35353afaf").isAlphaNumeric());
        T (S("FFFF99fff").isAlphaNumeric());
        T (S("99").isAlphaNumeric());
        T (S("afff").isAlphaNumeric());
        T (S("Infinity").isAlphaNumeric());
        T (S("áéúóúÁÉÍÓÚãõÃÕàèìòùÀÈÌÒÙâêîôûÂÊÎÔÛäëïöüÄËÏÖÜçÇ1234567890").isAlphaNumeric());
        F (S("-Infinity").isAlphaNumeric());
        F (S("-33").isAlphaNumeric());
        F (S("aaff..").isAlphaNumeric());
        F (S(".áéúóúÁÉÍÓÚãõÃÕàèìòùÀÈÌÒÙâêîôûÂÊÎÔÛäëïöüÄËÏÖÜçÇ1234567890").isAlphaNumeric());
      })
    })

    describe('- isEmpty()', function() {
     it('should return true if the string is solely composed of whitespace or is null', function() {
        T (S(' ').isEmpty());
        T (S('\t\t\t    ').isEmpty());
        T (S('\n\n ').isEmpty());
        F (S('hey').isEmpty())
        T (S(null).isEmpty())
        T (S(null).isEmpty())
      })
    })

    describe('- isLower()', function() {
     it('should return true if the character or string is lowercase', function() {
        T (S('a').isLower());
        T (S('z').isLower());
        F (S('B').isLower());
        T (S('hijp').isLower());
        T (S('áéúóúãõàèìòùâêîôûäëïöüç').isLower());
        T (S('áéúóúãõàèìòùâêîôûäëïöüça').isLower());
        F (S('hi jp').isLower());
        F (S('HelLO').isLower());
        F (S('ÁÉÍÓÚÃÕÀÈÌÒÙÂÊÎÔÛÄËÏÖÜÇ').isLower());
        F (S('áéúóúãõàèìòùâêîôûäëïöüçÁ').isLower());
        F (S('áéúóúãõàèìòùâêîôû äëïöüç').isLower());
      })
    })

    describe('- isNumeric()', function() {
      it("should return true if the string only contains digits, this would not include Infinity or -Infinity", function() {
        T (S("3").isNumeric());
        F (S("34.22").isNumeric());
        F (S("-22.33").isNumeric());
        F (S("NaN").isNumeric());
        F (S("Infinity").isNumeric());
        F (S("-Infinity").isNumeric());
        F (S("JP").isNumeric());
        F (S("-5").isNumeric());
        T (S("000992424242").isNumeric());
      })
    })

    describe('- isUpper()', function() {
      it('should return true if the character or string is uppercase', function() {
        F (S('a').isUpper());
        F (S('z').isUpper());
        T (S('B').isUpper());
        T (S('HIJP').isUpper());
        T (S('ÁÉÍÓÚÃÕÀÈÌÒÙÂÊÎÔÛÄËÏÖÜÇ').isUpper());
        F (S('HI JP').isUpper());
        F (S('HelLO').isUpper());
        F (S('áéúóúãõàèìòùâêîôûäëïöüç').isUpper());
        F (S('áéúóúãõàèìòùâêîôûäëïöüçÁ').isUpper());
        F (S('ÁÉÍÓÚÃÕÀÈÌÒÙÂÊÎÔÛÄËÏÖÜÇá').isUpper());
      })
    })

    describe('- length', function() {
      it('should return the length of the string', function() {
        T (S('hello').length === 5);
        T (S('').length === 0);
        T (S(null).length === -1);
        T (S(undefined).length === -1);
      })
    })

    describe('- left(N)', function() {
      it('should return the substring denoted by N positive left-most characters', function() {
        T (S('My name is JP').left(2).s === 'My');
        T (S('Hi').left(0).s === '');
        T (S('Hello').left(1).s === 'H');
      })
      it('should return the substring denoted by N negative left-most characters, equivalent to calling right(-N)', function() {
        T (S('My name is JP').left(-2).s === 'JP');
      })
    })

    describe('- pad(len, [char])', function() {
      it('should pad the string in the center with specified character', function() {
        T (S('hello').pad(5).s === 'hello');
        T (S('hello').pad(10).s === '   hello  ');
        T (S('hey').pad(7).s === '  hey  ');
        T (S('hey').pad(5).s === ' hey ');
        T (S('hey').pad(4).s === ' hey');
        T (S('hey').pad(7, '-').s === '--hey--');
      })
    })

    describe('- padLeft(len, [char])', function() {
      it('should left pad the string', function() {
        T (S('hello').padLeft(5).s === 'hello');
        T (S('hello').padLeft(10).s === '     hello');
        T (S('hello').padLeft(7).s === '  hello');
        T (S('hello').padLeft(6).s === ' hello');
        T (S('hello').padLeft(10, '.').s === '.....hello');
      })
    })

    describe('- padRight(len, [char])', function() {
      it('should right pad the string', function() {
        T (S('hello').padRight(5).s === 'hello');
        T (S('hello').padRight(10).s === 'hello     ');
        T (S('hello').padRight(7).s === 'hello  ');
        T (S('hello').padRight(6).s === 'hello ');
        T (S('hello').padRight(10, '.').s === 'hello.....');
      })
    })

    describe('- parseCSV([delim],[qualifier],[escape],[lineDelimiter])', function() {
      it('should parse a CSV line into an array', function() {
        ARY_EQ (S("'a','b','c'").parseCSV(',', "'"), ['a', 'b', 'c'])
        ARY_EQ (S('"a","b","c"').parseCSV(), ['a', 'b', 'c'])
        ARY_EQ (S('a,b,c').parseCSV(',', null), ['a', 'b', 'c'])
        ARY_EQ (S("'a,','b','c'").parseCSV(',', "'"), ['a,', 'b', 'c'])
        ARY_EQ (S('"a","b",4,"c"').parseCSV(',', null), ['"a"', '"b"', '4', '"c"'])
        ARY_EQ (S('"a","b","4","c"').parseCSV(), ['a', 'b', '4', 'c'])
        ARY_EQ (S('"a","b",       "4","c"').parseCSV(), ['a', 'b', '4', 'c'])
        ARY_EQ (S('"a","b",       4,"c"').parseCSV(",", null), [ '"a"', '"b"', '       4', '"c"' ])
        ARY_EQ (S('"a","b\\"","d","c"').parseCSV(), ['a', 'b"', 'd', 'c'])
        ARY_EQ (S('"jp","really\tlikes to code"').parseCSV(), ['jp', 'really\tlikes to code'])
        ARY_EQ (S('"a","b+"","d","c"').parseCSV(",", "\"", "+"), ['a', 'b"', 'd', 'c'])
        ARY_EQ (S('"a","b""","d","c"').parseCSV(",", "\"", "\""), ['a', 'b"', 'd', 'c'])
        ARY_EQ (S('"a","","c"').parseCSV(), ['a', '', 'c'])
        ARY_EQ (S('"","b","c"').parseCSV(), ['', 'b', 'c'])

        var lines = (S('"a\na","b","c"\n"a", """b\nb", "a"').parseCSV(',', '"', '"', '\n'));
        ARY_EQ(lines[0], [ 'a\na', 'b', 'c' ]);
        ARY_EQ(lines[1], [ 'a', '"b\nb', 'a' ]);
      })
    })

    describe('- repeat(n)', function() {
      it('should return the string concatenated with itself n times', function() {
        T (S(' ').repeat(5).s === '     ');
        T (S('*').repeat(3).s === '***');
      })
    })

    describe('- replaceAll(substring, replacement)', function() {
      it('should return the new string with all occurrences of substring replaced with the replacment string', function() {
        T (S(' does IT work? ').replaceAll(' ', '_').s === '_does_IT_work?_');
        T (S('Yes it does!').replaceAll(' ', '').s === 'Yesitdoes!')
        T (S('lalala.blabla').replaceAll('.', '_').s === 'lalala_blabla')

        var e = '\\', q = '"';
        var r = e + q;
        T (S('a').replaceAll(q, r).s === 'a');
      })
    })

    describe('+ restorePrototype()', function() {
      it('should restore the original String prototype', function() {
        T (typeof ' hi'.endsWith === 'undefined');
        S.extendPrototype();
        T (' hi'.endsWith('hi'));
        S.restorePrototype();
        T (typeof ' hi'.endsWith === 'undefined');
      })
    })

    describe('- right(N)', function() {
      it('should return the substring denoted by N positive right-most characters', function() {
        T (S('I AM CRAZY').right(2).s === 'ZY');
        T (S('Does it work?  ').right(4).s === 'k?  ');
        T (S('Hi').right(0).s === '');
      })
      it('should return the substring denoted by N negative right-most characters, equivalent to calling left(-N)', function() {
        T (S('My name is JP').right(-2).s === 'My');
      })
    })

    describe('- s', function() {
      it('should return the native string', function() {
        T (S('hi').s === 'hi');
        T (S('hi').toString() === S('hi').s);
      })
    })

    describe('- slugify', function() {
      it('should convert the text to url slug', function() {
        T (S('Global Thermonuclear Warfare').slugify().s === 'global-thermonuclear-warfare')
        T (S('Fast JSON Parsing').slugify().s === 'fast-json-parsing')
      })
    })

    describe('- startsWith(prefix)', function() {
      it("should return true if the string starts with the input string", function() {
        T (S("JP is a software engineer").startsWith("JP"));
        F (S('wants to change the world').startsWith("politicians"));
        T (S("").startsWith(""));
        T (S("Hi").startsWith(""));
        T (S("JP").startsWith("JP"));
      })
    })

    describe('- stripPunctuation()', function() {
      it('should strip all of the punctuation', function() {
        T (S('My, st[ring] *full* of %punct)').stripPunctuation().s === 'My string full of punct')
      })
    })

    describe('- stripTags([tag1],[tag2],...)', function() {
      it('should strip all of the html tags or tags specified by the parameters', function() {
        T (S('<p>just <b>some</b> text</p>').stripTags().s === 'just some text')
        T (S('<p>just <b>some</b> text</p>').stripTags('p').s === 'just <b>some</b> text')
      })
    })

    describe('- template(values, [open], [close])', function() {
      it('should return the string replaced with template values', function() {
        var str = "Hello {{name}}! How are you doing during the year of {{date-year}}?"
        var values = {name: 'JP', 'date-year': 2013}
        EQ (S(str).template(values).s, 'Hello JP! How are you doing during the year of 2013?')

        str = "Hello #{name}! How are you doing during the year of #{date-year}?"
        EQ (S(str).template(values, '#{', '}').s, 'Hello JP! How are you doing during the year of 2013?')

        S.TMPL_OPEN = '{'
        S.TMPL_CLOSE = '}'
        str = "Hello {name}! How are you doing during the year of {date-year}?"
        EQ (S(str).template(values).s, 'Hello JP! How are you doing during the year of 2013?')
      })
    })

    describe('- times(n)', function() {
      it('should return the string concatenated with itself n times', function() {
        T (S(' ').times(5).s === '     ');
        T (S('*').times(3).s === '***');
      })
    })

    describe('- toFloat([precision])', function() {
      it('should return the float value, wraps parseFloat', function() {
        T (S('5').toFloat() === 5);
        T (S('5.3').toFloat() === 5.3);
        T (S(5.3).toFloat() === 5.3);
        T (S('-10').toFloat() === -10);
        T (S('55.3 adfafaf').toFloat() === 55.3)
        T (S('afff 44').toFloat().toString() === 'NaN')
        T (S(3.45522222333232).toFloat(2) === 3.46)
      })
    })

    describe('- toBoolean', function() {
      it('should convert a logical truth string to boolean', function() {
        T (S('true').toBoolean());
        F (S('false').toBoolean());
        F (S('hello').toBoolean());
        T (S(true).toBoolean());
        T (S('on').toBoolean());
        T (S('yes').toBoolean());
        T (S('TRUE').toBoolean());
        T (S('TrUe').toBoolean());
        T (S('YES').toBoolean());
        T (S('ON').toBoolean());
        F (S('').toBoolean());
        F (S(undefined).toBoolean())
        F (S('undefined').toBoolean())
        F (S(null).toBoolean())
        F (S(false).toBoolean())
        F (S({}).toBoolean())
        T (S(1).toBoolean())
        F (S(-1).toBoolean())
        F (S(0).toBoolean())
      })
    })

    describe('- toCSV(options)', function() {
      it('should convert the array to csv', function() {
        EQ (S(['a', 'b', 'c']).toCSV().s, '"a","b","c"');
        EQ (S(['a', 'b', 'c']).toCSV(':').s, '"a":"b":"c"');
        EQ (S(['a', 'b', 'c']).toCSV(':', null).s, 'a:b:c');
        EQ (S(['a', 'b', 'c']).toCSV('*', "'").s, "'a'*'b'*'c'");
        EQ (S(['a"', 'b', 4, 'c']).toCSV({delimiter: ',', qualifier: '"', escape: '\\',  encloseNumbers: false}).s, '"a\\"","b",4,"c"');
        EQ (S({firstName: 'JP', lastName: 'Richardson'}).toCSV({keys: true}).s, '"firstName","lastName"');
        EQ (S({firstName: 'JP', lastName: 'Richardson'}).toCSV().s, '"JP","Richardson"');
        EQ (S(['a', null, undefined, 'c']).toCSV().s, '"a","","","c"');
        EQ (S(['my "foo" bar', 'barf']).toCSV({delimiter: ';', qualifier: '"', escape: '"'}).s, '"my ""foo"" bar";"barf"');
      })
    })

    describe('- toInt()', function() {
      it('should return the integer value, wraps parseInt', function() {
        T (S('5').toInt() === 5);
        T (S('5.3').toInt() === 5);
        T (S(5.3).toInt() === 5);
        T (S('-10').toInt() === -10);
        T (S('55 adfafaf').toInt() === 55)
        T (S('afff 44').toInt().toString() === 'NaN')
        T (S('0xff').toInt() == 255)
      })
    })

    describe('- toString()', function() {
      it('should return the native string', function() {
        T (S('hi').toString() === 'hi');
        T (S('hi').toString() === S('hi').s);
      })
    })

    describe('- trim()', function() {
      it('should return the string with leading and trailing whitespace removed', function() {
        T (S('hello ').trim().s === 'hello');
        T (S(' hello ').trim().s === 'hello');
        T (S('\nhello').trim().s === 'hello');
        T (S('\nhello\r\n').trim().s === 'hello');
        T (S('\thello\t').trim().s === 'hello');
      })
    })

    describe('- trimLeft()', function() {
      it('should return the string with leading whitespace removed', function() {
        T (S('  How are you?').trimLeft().s === 'How are you?');
        T (S(' JP ').trimLeft().s === 'JP ');
      })
    })

    describe('- trimRight()', function() {
      it('should return the string with trailing whitespace removed', function() {
        T (S('How are you?  ').trimRight().s === 'How are you?');
        T (S(' JP ').trimRight().s === ' JP');
      })
    })

    describe('- truncate(length, [chars])', function() {
      it('should truncate the string, accounting for word placement and chars count', function() {
        T (S('this is some long text').truncate(3).s === '...')
        T (S('this is some long text').truncate(7).s === 'this is...')
        T (S('this is some long text').truncate(11).s === 'this is...')
        T (S('this is some long text').truncate(12).s === 'this is some...')
        T (S('this is some long text').truncate(11).s === 'this is...')
        T (S('this is some long text').truncate(14, ' read more').s === 'this is some read more')
        EQ (S('some string').truncate(200).s, 'some string')
      })
    })

    describe('- underscore()', function() {
      it('should convert a camel cased string into a string separated by underscores', function() {
        T (S('dataRate').underscore().s === 'data_rate');
        T (S('CarSpeed').underscore().s === '_car_speed');
        T (S('yesWeCan').underscore().s === 'yes_we_can');
      })
    })

    describe('- unescapeHTML', function() {
      it('should unescape the HTML', function() {
        T (S('&lt;div&gt;Blah &amp; &quot;blah&quot; &amp; &apos;blah&apos;&lt;/div&gt;').unescapeHTML().s ===
             '<div>Blah & "blah" & \'blah\'</div>');
        T (S('&amp;lt;').unescapeHTML().s === '&lt;');
      })
    })

    describe('- valueOf()', function() {
      it('should return the primitive value of the string, wraps native valueOf()', function() {
        T (S('hi').valueOf() === 'hi')
      })
    })

    describe('+ VERSION', function() {
      it('should exist', function() {
        T (S.VERSION)
      })
    })

    it('should import native JavaScript string methods', function() {
      T (S('hi    ').substr(0,1).trimRight().startsWith('h'));
      T (S('hello ').concat('jp').indexOf('jp') === 6);
      T (S('this is so cool').substr(0, 4).s === 'this');
    })

  })
}).call(this);
