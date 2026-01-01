import { Plugin } from 'vite';
import { UserDefinedOptions, ComplexSafelist } from 'purgecss';

type Options = Partial<UserDefinedOptions> & {
    safelist?: ComplexSafelist;
};
type PurgeOptions = Omit<Options, 'css'>;
declare function purgeCss(purgeOptions?: PurgeOptions): Plugin;

export { purgeCss as default, purgeCss };
