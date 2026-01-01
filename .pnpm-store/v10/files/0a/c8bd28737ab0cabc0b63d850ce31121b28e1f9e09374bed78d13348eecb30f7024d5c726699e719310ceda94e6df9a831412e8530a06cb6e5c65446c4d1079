// Transitions ---
export function dynamicTransition(node, dynParams) {
    const { transition, params, enabled } = dynParams;
    if (enabled)
        return transition(node, params);
    // it's better to just set the `duration` to 0 to prevent flickering
    if ('duration' in params)
        return transition(node, { duration: 0 });
    // if the transition doesn't provide a `duration` prop, then we'll just return this as a last resort
    return { duration: 0 };
}
