// Action: Filter
export function filter(node, filterName) {
    // Return if no filterName provided
    if (filterName === undefined)
        return;
    const applyFilter = () => {
        node.setAttribute('style', `filter: url("${filterName}")`);
    };
    applyFilter();
    return {
        update(newArgs) {
            filterName = newArgs;
            applyFilter();
        }
    };
}
