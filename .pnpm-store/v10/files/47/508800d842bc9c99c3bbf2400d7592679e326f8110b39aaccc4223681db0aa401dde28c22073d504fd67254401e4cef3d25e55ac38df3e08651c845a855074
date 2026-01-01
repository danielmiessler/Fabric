import { cubicInOut } from 'svelte/easing';
import { crossfade } from 'svelte/transition';
// TODO: convert to js transitions on revisit.
// eslint-disable-next-line @typescript-eslint/no-unused-vars
function swoop(duration, x, y) {
    return crossfade({
        duration: (d) => Math.sqrt(d * duration),
        fallback(node) {
            const style = getComputedStyle(node);
            const transform = style.transform === 'none' ? '' : style.transform;
            return {
                duration,
                easing: cubicInOut,
                css: (t, u) => `
                    transform: ${transform} scale(${t}) translate(${u * x}%, ${u * y}%);
                    opacity: ${t}
                `
            };
        }
    });
}
