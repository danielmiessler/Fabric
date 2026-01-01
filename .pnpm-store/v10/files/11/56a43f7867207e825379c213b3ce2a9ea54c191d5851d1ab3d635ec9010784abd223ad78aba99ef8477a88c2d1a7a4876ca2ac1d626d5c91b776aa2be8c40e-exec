#!/usr/bin/env node

let help = `
Usage: mini-svg-data-uri <source> [dest]

Options:
-v, --version  Output the version number
-h, --help     Display help for command

Examples:
  mini-svg-data-uri file.svg           Write to stdout
  mini-svg-data-uri icon.svg icon.uri  Write to file
`;

let [source, dest] = process.argv.slice(2);

switch (source) {
  case '-h':
  case '--help':
  case undefined:
    console.log(help);
    process.exit();

  case '-v':
  case '--version':
    console.log(require('./package').version);
    process.exit();
}

const fs = require('fs');
const svgToMiniDataURI = require('.');

fs.readFile(source, 'utf8', (err, data) => {
  if (err) {
    console.error(err.message);
    console.log(help);
    process.exit(1);
  }
  const out = svgToMiniDataURI(data);
  dest ? fs.writeFileSync(dest, out) : console.log(out);
});
