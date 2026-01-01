// Table Component Utilities
// Cell Formatters ---
// NOTE: this would require `onMount`, which is too slow, so may just nix this.
// REMINDER: if re-enabled, update `index.ts`!
/** Wrap object key value with an HTML tag.  */
// export function tableCellFormatter(source: any[], key: string, tagName: string, classes?: string) {
// 	return source.map((row) => {
// 		if (row[key]) {
// 			const newElem: HTMLElement = document.createElement(tagName);
// 			newElem.innerHTML = row[key];
// 			if (classes) newElem.setAttribute('class', classes);
// 			row[key] = newElem.outerHTML;
// 		}
// 		return row;
// 	});
// }
// Source Formatters ---
/** Map an object to a defined order. */
export function tableSourceMapper(source, keys) {
    return source.map((row) => {
        const mappedRow = {};
        keys.forEach((key) => (mappedRow[key] = row[key]));
        return mappedRow;
    });
}
/** Map an array of objects to an array of values. */
export function tableSourceValues(source) {
    return source.map((row) => Object.values(row));
}
/** Sets object order and returns values. */
export function tableMapperValues(source, keys) {
    return tableSourceValues(tableSourceMapper(source, keys));
}
