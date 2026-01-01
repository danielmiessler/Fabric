export const emulationGroupMarker: "$E$";
/**
Works the same as JavaScript's native `RegExp` constructor in all contexts, but automatically
adjusts matches and subpattern indices (with flag `d`) to account for injected emulation groups.
*/
export class RegExpSubclass extends RegExp {
    /**
    @param {string | RegExpSubclass} expression
    @param {string} [flags]
    @param {{useEmulationGroups: boolean;}} [options]
    */
    constructor(expression: string | RegExpSubclass, flags?: string, options?: {
        useEmulationGroups: boolean;
    });
    /**
    @private
    @type {Array<{
      exclude: boolean;
      transfer?: number;
    }> | undefined}
    */
    private _captureMap;
    /**
    @private
    @type {Record<number, string> | undefined}
    */
    private _namesByIndex;
}
