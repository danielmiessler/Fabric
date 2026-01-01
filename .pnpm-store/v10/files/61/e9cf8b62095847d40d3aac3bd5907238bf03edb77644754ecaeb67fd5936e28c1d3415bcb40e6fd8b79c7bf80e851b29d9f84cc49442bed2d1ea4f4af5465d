// Shared Data/Table Actions
// Data Table (only) ---
/** Svelte Action for applying sort asc/dsc classes. */
export function tableInteraction(node) {
    const classAsc = 'table-sort-asc';
    const classDsc = 'table-sort-dsc';
    // Click Handler
    const onClick = (e) => {
        if (!(e.target instanceof Element))
            return;
        const sortTarget = e.target;
        // Get target state before modification
        const targetAscSorted = sortTarget.classList.contains(classAsc);
        const sortTargetKey = sortTarget.getAttribute('data-sort');
        // Clear asc class
        const elemAsc = node.querySelector(`.${classAsc}`);
        if (elemAsc)
            elemAsc.classList.remove(classAsc);
        // Clear dsc class
        const elemDsc = node.querySelector(`.${classDsc}`);
        if (elemDsc)
            elemDsc.classList.remove(classDsc);
        // Set new sort class
        if (sortTargetKey) {
            const classToApply = targetAscSorted ? classDsc : classAsc;
            e.target.classList.add(classToApply);
        }
    };
    // Events
    node.addEventListener('click', onClick);
    // Lifecycle
    return {
        destroy() {
            node.removeEventListener('click', onClick);
        }
    };
}
// Shared ---
/** Svelte Action for handling table a11y keyboard interactions. */
export function tableA11y(node) {
    const keyWhitelist = ['ArrowRight', 'ArrowUp', 'ArrowLeft', 'ArrowDown', 'Home', 'End'];
    // on:keydown
    const onKeyDown = (event) => {
        // console.log('keydown triggered');
        if (keyWhitelist.includes(event.code)) {
            event.preventDefault();
            // prettier-ignore
            switch (event.code) {
                case 'ArrowUp':
                    a11ySetActiveCell(node, 0, -1);
                    break;
                case 'ArrowDown':
                    a11ySetActiveCell(node, 0, 1);
                    break;
                case 'ArrowLeft':
                    a11ySetActiveCell(node, -1, 0);
                    break;
                case 'ArrowRight':
                    a11ySetActiveCell(node, 1, 0);
                    break;
                case 'Home':
                    a11yJumpToOuterColumn(node, 'first');
                    break;
                case 'End':
                    a11yJumpToOuterColumn(node, 'last');
                    break;
                default: break;
            }
        }
    };
    // Event Listener
    node.addEventListener('keydown', onKeyDown);
    // Lifecycle
    return {
        destroy() {
            node.removeEventListener('keydown', onKeyDown);
        }
    };
}
function a11ySetActiveCell(node, x, y) {
    // Focused Element
    const focusedElem = document.activeElement;
    if (!focusedElem || !focusedElem.parentElement || !focusedElem.parentElement.ariaRowIndex || !focusedElem.ariaColIndex)
        return;
    const focusedElemRowIndex = parseInt(focusedElem.parentElement.ariaRowIndex);
    const focusedElemColIndex = parseInt(focusedElem.ariaColIndex);
    // Target Element
    const targetRowElement = node.querySelector(`[aria-rowindex="${focusedElemRowIndex + y}"]`);
    if (targetRowElement !== null) {
        const targetColElement = targetRowElement.querySelector(`[aria-colindex="${focusedElemColIndex + x}"]`);
        if (targetColElement !== null)
            targetColElement.focus();
    }
}
function a11yGetTargetElem(node) {
    // Focused Element
    const focusedElem = document.activeElement;
    if (!focusedElem || !focusedElem.parentElement || !focusedElem.parentElement.ariaRowIndex)
        return null;
    const focusedElemRowIndex = parseInt(focusedElem.parentElement.ariaRowIndex);
    // Return Target Element
    return node.querySelector(`[aria-rowindex="${focusedElemRowIndex}"]`);
}
function a11yJumpToOuterColumn(node, type = 'first') {
    const targetRowElement = a11yGetTargetElem(node);
    if (targetRowElement === null)
        return;
    const lastIndex = targetRowElement.children.length;
    const selected = type === 'first' ? 1 : lastIndex;
    const targetColElement = targetRowElement.querySelector(`[aria-colindex="${selected}"]`);
    if (targetColElement === null)
        return;
    targetColElement.focus();
}
