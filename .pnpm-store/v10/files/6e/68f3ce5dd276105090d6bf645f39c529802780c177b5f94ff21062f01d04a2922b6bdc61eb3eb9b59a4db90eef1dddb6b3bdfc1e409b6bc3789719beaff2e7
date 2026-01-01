import {Context, forEachUnescaped, getGroupContents, hasUnescaped, replaceUnescaped} from 'regex-utilities';
import {emulationGroupMarker} from 'regex/internals';

const r = String.raw;
const gRToken = r`\\g<(?<gRNameOrNum>[^>&]+)&R=(?<gRDepth>[^>]+)>`;
const recursiveToken = r`\(\?R=(?<rDepth>[^\)]+)\)|${gRToken}`;
const namedCapturingDelim = r`\(\?<(?![=!])(?<captureName>[^>]+)>`;
const token = new RegExp(r`${namedCapturingDelim}|${recursiveToken}|\(\?|\\?.`, 'gsu');
const overlappingRecursionMsg = 'Cannot use multiple overlapping recursions';
// Support emulation groups with transfer marker prefix
const emulationGroupMarkerRe = new RegExp(r`(?:\$[1-9]\d*)?${emulationGroupMarker.replace(/\$/g, r`\$`)}`, 'y');

/**
@param {string} expression
@param {{
  flags?: string;
  useEmulationGroups?: boolean;
}} [data]
@returns {string}
*/
export function recursion(expression, data) {
  // Keep the initial fail-check (which avoids unneeded processing) as fast as possible by testing
  // without the accuracy improvement of using `hasUnescaped` with default `Context`
  if (!(new RegExp(recursiveToken, 'su').test(expression))) {
    return expression;
  }
  if (hasUnescaped(expression, r`\(\?\(DEFINE\)`, Context.DEFAULT)) {
    throw new Error('DEFINE groups cannot be used with recursion');
  }
  const useEmulationGroups = !!data?.useEmulationGroups;
  const hasNumberedBackref = hasUnescaped(expression, r`\\[1-9]`, Context.DEFAULT);
  const groupContentsStartPos = new Map();
  const openGroups = [];
  let hasRecursed = false;
  let numCharClassesOpen = 0;
  let numCaptures = 0;
  let match;
  token.lastIndex = 0;
  while ((match = token.exec(expression))) {
    const {0: m, groups: {captureName, rDepth, gRNameOrNum, gRDepth}} = match;
    if (m === '[') {
      numCharClassesOpen++;
    } else if (!numCharClassesOpen) {

      // `(?R=N)`
      if (rDepth) {
        assertMaxInBounds(rDepth);
        if (hasRecursed) {
          throw new Error(overlappingRecursionMsg);
        }
        if (hasNumberedBackref) {
          // Could add support for numbered backrefs with extra effort, but it's probably not worth
          // it. To trigger this error, the regex must include recursion and one of the following:
          // - An interpolated regex that contains a numbered backref (since other numbered
          //   backrefs are prevented by implicit flag n).
          // - A numbered backref, when flag n is explicitly disabled.
          // Note that Regex+'s extended syntax (atomic groups and sometimes subroutines) can also
          // add numbered backrefs, but those work fine because external plugins like this one run
          // *before* the transformation of built-in syntax extensions
          throw new Error('Numbered backrefs cannot be used with global recursion');
        }
        const pre = expression.slice(0, match.index);
        const post = expression.slice(token.lastIndex);
        if (hasUnescaped(post, recursiveToken, Context.DEFAULT)) {
          throw new Error(overlappingRecursionMsg);
        }
        // No need to parse further
        return makeRecursive(pre, post, +rDepth, false, useEmulationGroups);
      // `\g<name&R=N>`, `\g<number&R=N>`
      } else if (gRNameOrNum) {
        assertMaxInBounds(gRDepth);
        let isWithinReffedGroup = false;
        for (const g of openGroups) {
          if (g.name === gRNameOrNum || g.num === +gRNameOrNum) {
            isWithinReffedGroup = true;
            if (g.hasRecursedWithin) {
              throw new Error(overlappingRecursionMsg);
            }
            break;
          }
        }
        if (!isWithinReffedGroup) {
          throw new Error(r`Recursive \g cannot be used outside the referenced group "\g<${gRNameOrNum}&R=${gRDepth}>"`);
        }
        const startPos = groupContentsStartPos.get(gRNameOrNum);
        const groupContents = getGroupContents(expression, startPos);
        if (
          hasNumberedBackref &&
          hasUnescaped(groupContents, r`${namedCapturingDelim}|\((?!\?)`, Context.DEFAULT)
        ) {
          throw new Error('Numbered backrefs cannot be used with recursion of capturing groups');
        }
        const groupContentsPre = expression.slice(startPos, match.index);
        const groupContentsPost = groupContents.slice(groupContentsPre.length + m.length);
        const expansion = makeRecursive(groupContentsPre, groupContentsPost, +gRDepth, true, useEmulationGroups);
        const pre = expression.slice(0, startPos);
        const post = expression.slice(startPos + groupContents.length);
        // Modify the string we're looping over
        expression = `${pre}${expansion}${post}`;
        // Step forward for the next loop iteration
        token.lastIndex += expansion.length - m.length - groupContentsPre.length - groupContentsPost.length;
        openGroups.forEach(g => g.hasRecursedWithin = true);
        hasRecursed = true;
      } else if (captureName) {
        numCaptures++;
        // NOTE: Not currently handling *named* emulation groups that already exist in the pattern
        groupContentsStartPos.set(String(numCaptures), token.lastIndex);
        groupContentsStartPos.set(captureName, token.lastIndex);
        openGroups.push({
          num: numCaptures,
          name: captureName,
        });
      } else if (m.startsWith('(')) {
        const isUnnamedCapture = m === '(';
        if (isUnnamedCapture) {
          numCaptures++;
          groupContentsStartPos.set(
            String(numCaptures),
            token.lastIndex + (useEmulationGroups ? emulationGroupMarkerLength(expression, token.lastIndex) : 0)
          );
        }
        openGroups.push(isUnnamedCapture ? {num: numCaptures} : {});
      } else if (m === ')') {
        openGroups.pop();
      }

    } else if (m === ']') {
      numCharClassesOpen--;
    }
  }

  return expression;
}

/**
@param {string} max
*/
function assertMaxInBounds(max) {
  const errMsg = `Max depth must be integer between 2 and 100; used ${max}`;
  if (!/^[1-9]\d*$/.test(max)) {
    throw new Error(errMsg);
  }
  max = +max;
  if (max < 2 || max > 100) {
    throw new Error(errMsg);
  }
}

/**
@param {string} pre
@param {string} post
@param {number} maxDepth
@param {boolean} isSubpattern
@param {boolean} useEmulationGroups
@returns {string}
*/
function makeRecursive(pre, post, maxDepth, isSubpattern, useEmulationGroups) {
  const namesInRecursed = new Set();
  // Avoid this work if not needed
  if (isSubpattern) {
    forEachUnescaped(pre + post, namedCapturingDelim, ({groups: {captureName}}) => {
      namesInRecursed.add(captureName);
    }, Context.DEFAULT);
  }
  const reps = maxDepth - 1;
  // Depth 2: 'pre(?:pre(?:)post)post'
  // Depth 3: 'pre(?:pre(?:pre(?:)post)post)post'
  return `${pre}${
    repeatWithDepth(`(?:${pre}`, reps, (isSubpattern ? namesInRecursed : null), 'forward', useEmulationGroups)
  }(?:)${
    repeatWithDepth(`${post})`, reps, (isSubpattern ? namesInRecursed : null), 'backward', useEmulationGroups)
  }${post}`;
}

/**
@param {string} expression
@param {number} reps
@param {Set<string> | null} namesInRecursed
@param {'forward' | 'backward'} direction
@param {boolean} useEmulationGroups
@returns {string}
*/
function repeatWithDepth(expression, reps, namesInRecursed, direction, useEmulationGroups) {
  const startNum = 2;
  const depthNum = i => direction === 'backward' ? reps - i + startNum - 1 : i + startNum;
  let result = '';
  for (let i = 0; i < reps; i++) {
    const captureNum = depthNum(i);
    result += replaceUnescaped(
      expression,
      // NOTE: Not currently handling *named* emulation groups that already exist in the pattern
      r`${namedCapturingDelim}|\\k<(?<backref>[^>]+)>${
        useEmulationGroups ? r`|(?<unnamed>\()(?!\?)(?:${emulationGroupMarkerRe.source})?` : ''
      }`,
      ({0: m, index, groups: {captureName, backref, unnamed}}) => {
        if (backref && namesInRecursed && !namesInRecursed.has(backref)) {
          // Don't alter backrefs to groups outside the recursed subpattern
          return m;
        }
        // Only matches unnamed capture delim if `useEmulationGroups`
        if (unnamed) {
          // Add an emulation group marker, possibly replacing an existing marker (removes any
          // transfer prefix)
          return `(${emulationGroupMarker}`;
        }
        const suffix = `_$${captureNum}`;
        return captureName ?
          `(?<${captureName}${suffix}>${useEmulationGroups ? emulationGroupMarker : ''}` :
          r`\k<${backref}${suffix}>`;
      },
      Context.DEFAULT
    );
  }
  return result;
}

function emulationGroupMarkerLength(expression, index) {
  emulationGroupMarkerRe.lastIndex = index;
  const match = emulationGroupMarkerRe.exec(expression);
  return match ? match[0].length : 0;
}
