import {Context, replaceUnescaped} from 'regex-utilities';

// This marker was chosen because it's impossible to match (so its extemely unlikely to be used in
// a user-provided regex); it's not at risk of being optimized away, transformed, or flagged as an
// error by a plugin; and it ends with an unquantifiable token
const emulationGroupMarker = '$E$';
// Note: Emulation groups with transfer are also supported. They look like `($N$E$…)` where `N` is
// an integer 1 or greater. They're not used directly by Regex+ but can be used by plugins and
// libraries that use Regex+ internals. Emulation groups with transfer are not only excluded from
// match results, but additionally transfer their match to the group specified by `N`

/**
Works the same as JavaScript's native `RegExp` constructor in all contexts, but automatically
adjusts matches and subpattern indices (with flag `d`) to account for injected emulation groups.
*/
class RegExpSubclass extends RegExp {
  // Avoid `#private` to allow for subclassing
  /**
  @private
  @type {Array<{
    exclude: boolean;
    transfer?: number;
  }> | undefined}
  */
  _captureMap;
  /**
  @private
  @type {Record<number, string> | undefined}
  */
  _namesByIndex;
  /**
  @param {string | RegExpSubclass} expression
  @param {string} [flags]
  @param {{useEmulationGroups: boolean;}} [options]
  */
  constructor(expression, flags, options) {
    if (expression instanceof RegExp && options) {
      throw new Error('Cannot provide options when copying a regexp');
    }
    const useEmulationGroups = !!options?.useEmulationGroups;
    const unmarked = useEmulationGroups ? unmarkEmulationGroups(expression) : null;
    super(unmarked?.expression || expression, flags);
    // The third argument `options` isn't provided when regexes are copied as part of the internal
    // handling of string methods `matchAll` and `split`
    const src = useEmulationGroups ? unmarked : (expression instanceof RegExpSubclass ? expression : null);
    if (src) {
      this._captureMap = src._captureMap;
      this._namesByIndex = src._namesByIndex;
    }
  }
  /**
  Called internally by all String/RegExp methods that use regexes.
  @override
  @param {string} str
  @returns {RegExpExecArray | null}
  */
  exec(str) {
    const match = RegExp.prototype.exec.call(this, str);
    if (!match || !this._captureMap) {
      return match;
    }
    const matchCopy = [...match];
    // Empty all but the first value of the array while preserving its other properties
    match.length = 1;
    let indicesCopy;
    if (this.hasIndices) {
      indicesCopy = [...match.indices];
      match.indices.length = 1;
    }
    for (let i = 1; i < matchCopy.length; i++) {
      if (this._captureMap[i].exclude) {
        const transfer = this._captureMap[i].transfer;
        if (transfer && match.length > transfer) {
          match[transfer] = matchCopy[i];
          const transferName = this._namesByIndex[transfer];
          if (transferName) {
            match.groups[transferName] = matchCopy[i];
            if (this.hasIndices) {
              match.indices.groups[transferName] = indicesCopy[i];
            }
          }
          if (this.hasIndices) {
            match.indices[transfer] = indicesCopy[i];
          }
        }
      } else {
        match.push(matchCopy[i]);
        if (this.hasIndices) {
          match.indices.push(indicesCopy[i]);
        }
      }
    }
    return match;
  }
}

/**
Build the capturing group map (with emulation groups marked to indicate their submatches shouldn't
appear in results), and remove the markers for captures that were added to emulate extended syntax.
@param {string} expression
@returns {{
  _captureMap: Array<{
    exclude: boolean;
    transfer?: number;
  }>;
  _namesByIndex: Record<number, string>;
  expression: string;
}}
*/
function unmarkEmulationGroups(expression) {
  const marker = emulationGroupMarker.replace(/\$/g, '\\$');
  const _captureMap = [{exclude: false}];
  const _namesByIndex = {0: ''};
  let realCaptureNum = 0;
  expression = replaceUnescaped(
    expression,
    String.raw`\((?:(?!\?)|\?<(?![=!])(?<name>[^>]+)>)(?<mark>(?:\$(?<transfer>[1-9]\d*))?${marker})?`,
    ({0: m, groups: {name, mark, transfer}}) => {
      if (mark) {
        _captureMap.push({
          exclude: true,
          transfer: transfer && +transfer,
        });
        return m.slice(0, -mark.length);
      }
      realCaptureNum++;
      if (name) {
        _namesByIndex[realCaptureNum] = name;
      }
      _captureMap.push({
        exclude: false,
      });
      return m;
    },
    Context.DEFAULT
  );
  return {
    _captureMap,
    _namesByIndex,
    expression,
  };
}

export {
  emulationGroupMarker,
  RegExpSubclass,
};
