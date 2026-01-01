/**
 * @license lucide-svelte v0.309.0 - ISC

This source code is licensed under the ISC license.
See the LICENSE file in the root directory of this source tree.
 */

import type { SvelteComponent } from 'svelte';
import type { SVGAttributes, SvelteHTMLElements } from 'svelte/elements';
export type Attrs = SVGAttributes<SVGSVGElement>;
export type IconNode = [elementName: keyof SvelteHTMLElements, attrs: Attrs][];
export interface IconProps extends Attrs {
    color?: string;
    size?: number | string;
    strokeWidth?: number | string;
    absoluteStrokeWidth?: boolean;
    class?: string;
}
export type IconEvents = {
    [evt: string]: CustomEvent<any>;
};
export type IconSlots = {
    default: {};
};
export type Icon = SvelteComponent<IconProps, IconEvents, IconSlots>;
