/** @import { RemoteForm } from '@sveltejs/kit' */
/** @import { BinaryFormMeta, InternalRemoteFormIssue } from 'types' */
/** @import { StandardSchemaV1 } from '@standard-schema/spec' */

import { DEV } from 'esm-env';
import * as devalue from 'devalue';
import { text_decoder, text_encoder } from './utils.js';

/**
 * Sets a value in a nested object using a path string, mutating the original object
 * @param {Record<string, any>} object
 * @param {string} path_string
 * @param {any} value
 */
export function set_nested_value(object, path_string, value) {
	if (path_string.startsWith('n:')) {
		path_string = path_string.slice(2);
		value = value === '' ? undefined : parseFloat(value);
	} else if (path_string.startsWith('b:')) {
		path_string = path_string.slice(2);
		value = value === 'on';
	}

	deep_set(object, split_path(path_string), value);
}

/**
 * Convert `FormData` into a POJO
 * @param {FormData} data
 */
export function convert_formdata(data) {
	/** @type {Record<string, any>} */
	const result = {};

	for (let key of data.keys()) {
		const is_array = key.endsWith('[]');
		/** @type {any[]} */
		let values = data.getAll(key);

		if (is_array) key = key.slice(0, -2);

		if (values.length > 1 && !is_array) {
			throw new Error(`Form cannot contain duplicated keys — "${key}" has ${values.length} values`);
		}

		// an empty `<input type="file">` will submit a non-existent file, bizarrely
		values = values.filter(
			(entry) => typeof entry === 'string' || entry.name !== '' || entry.size > 0
		);

		if (key.startsWith('n:')) {
			key = key.slice(2);
			values = values.map((v) => (v === '' ? undefined : parseFloat(/** @type {string} */ (v))));
		} else if (key.startsWith('b:')) {
			key = key.slice(2);
			values = values.map((v) => v === 'on');
		}

		set_nested_value(result, key, is_array ? values : values[0]);
	}

	return result;
}

export const BINARY_FORM_CONTENT_TYPE = 'application/x-sveltekit-formdata';
const BINARY_FORM_VERSION = 0;

/**
 * The binary format is as follows:
 * - 1 byte: Format version
 * - 4 bytes: Length of the header (u32)
 * - 2 bytes: Length of the file offset table (u16)
 * - header: devalue.stringify([data, meta])
 * - file offset table: JSON.stringify([offset1, offset2, ...]) (empty if no files) (offsets start from the end of the table)
 * - file1, file2, ...
 * @param {Record<string, any>} data
 * @param {BinaryFormMeta} meta
 */
export function serialize_binary_form(data, meta) {
	/** @type {Array<BlobPart>} */
	const blob_parts = [new Uint8Array([BINARY_FORM_VERSION])];

	/** @type {Array<[file: File, index: number]>} */
	const files = [];

	if (!meta.remote_refreshes?.length) {
		delete meta.remote_refreshes;
	}

	const encoded_header = devalue.stringify([data, meta], {
		File: (file) => {
			if (!(file instanceof File)) return;

			files.push([file, files.length]);
			return [file.name, file.type, file.size, file.lastModified, files.length - 1];
		}
	});

	const encoded_header_buffer = text_encoder.encode(encoded_header);

	let encoded_file_offsets = '';
	if (files.length) {
		// Sort small files to the front
		files.sort(([a], [b]) => a.size - b.size);

		/** @type {Array<number>} */
		const file_offsets = new Array(files.length);
		let start = 0;
		for (const [file, index] of files) {
			file_offsets[index] = start;
			start += file.size;
		}
		encoded_file_offsets = JSON.stringify(file_offsets);
	}

	const length_buffer = new Uint8Array(4);
	const length_view = new DataView(length_buffer.buffer);

	length_view.setUint32(0, encoded_header_buffer.byteLength, true);
	blob_parts.push(length_buffer.slice());

	length_view.setUint16(0, encoded_file_offsets.length, true);
	blob_parts.push(length_buffer.slice(0, 2));

	blob_parts.push(encoded_header_buffer);
	blob_parts.push(encoded_file_offsets);

	for (const [file] of files) {
		blob_parts.push(file);
	}

	return {
		blob: new Blob(blob_parts)
	};
}

/**
 * @param {Request} request
 * @returns {Promise<{ data: Record<string, any>; meta: BinaryFormMeta; form_data: FormData | null }>}
 */
export async function deserialize_binary_form(request) {
	if (request.headers.get('content-type') !== BINARY_FORM_CONTENT_TYPE) {
		const form_data = await request.formData();
		return { data: convert_formdata(form_data), meta: {}, form_data };
	}
	if (!request.body) {
		throw new Error('Could not deserialize binary form: no body');
	}

	const reader = request.body.getReader();

	/** @type {Array<Promise<Uint8Array<ArrayBuffer> | undefined>>} */
	const chunks = [];

	/**
	 * @param {number} index
	 * @returns {Promise<Uint8Array<ArrayBuffer> | undefined>}
	 */
	async function get_chunk(index) {
		if (index in chunks) return chunks[index];

		let i = chunks.length;
		while (i <= index) {
			chunks[i] = reader.read().then((chunk) => chunk.value);
			i++;
		}
		return chunks[index];
	}

	/**
	 * @param {number} offset
	 * @param {number} length
	 * @returns {Promise<Uint8Array | null>}
	 */
	async function get_buffer(offset, length) {
		/** @type {Uint8Array} */
		let start_chunk;
		let chunk_start = 0;
		/** @type {number} */
		let chunk_index;
		for (chunk_index = 0; ; chunk_index++) {
			const chunk = await get_chunk(chunk_index);
			if (!chunk) return null;

			const chunk_end = chunk_start + chunk.byteLength;
			// If this chunk contains the target offset
			if (offset >= chunk_start && offset < chunk_end) {
				start_chunk = chunk;
				break;
			}
			chunk_start = chunk_end;
		}
		// If the buffer is completely contained in one chunk, do a subarray
		if (offset + length <= chunk_start + start_chunk.byteLength) {
			return start_chunk.subarray(offset - chunk_start, offset + length - chunk_start);
		}
		// Otherwise, copy the data into a new buffer
		const buffer = new Uint8Array(length);
		buffer.set(start_chunk.subarray(offset - chunk_start));
		let cursor = start_chunk.byteLength - offset + chunk_start;
		while (cursor < length) {
			chunk_index++;
			let chunk = await get_chunk(chunk_index);
			if (!chunk) return null;
			if (chunk.byteLength > length - cursor) {
				chunk = chunk.subarray(0, length - cursor);
			}
			buffer.set(chunk, cursor);
			cursor += chunk.byteLength;
		}

		return buffer;
	}

	const header = await get_buffer(0, 1 + 4 + 2);
	if (!header) throw new Error('Could not deserialize binary form: too short');

	if (header[0] !== BINARY_FORM_VERSION) {
		throw new Error(
			`Could not deserialize binary form: got version ${header[0]}, expected version ${BINARY_FORM_VERSION}`
		);
	}
	const header_view = new DataView(header.buffer, header.byteOffset, header.byteLength);
	const data_length = header_view.getUint32(1, true);
	const file_offsets_length = header_view.getUint16(5, true);

	// Read the form data
	const data_buffer = await get_buffer(1 + 4 + 2, data_length);
	if (!data_buffer) throw new Error('Could not deserialize binary form: data too short');

	/** @type {Array<number>} */
	let file_offsets;
	/** @type {number} */
	let files_start_offset;
	if (file_offsets_length > 0) {
		// Read the file offset table
		const file_offsets_buffer = await get_buffer(1 + 4 + 2 + data_length, file_offsets_length);
		if (!file_offsets_buffer)
			throw new Error('Could not deserialize binary form: file offset table too short');

		file_offsets = /** @type {Array<number>} */ (
			JSON.parse(text_decoder.decode(file_offsets_buffer))
		);
		files_start_offset = 1 + 4 + 2 + data_length + file_offsets_length;
	}

	const [data, meta] = devalue.parse(text_decoder.decode(data_buffer), {
		File: ([name, type, size, last_modified, index]) => {
			return new Proxy(
				new LazyFile(
					name,
					type,
					size,
					last_modified,
					get_chunk,
					files_start_offset + file_offsets[index]
				),
				{
					getPrototypeOf() {
						// Trick validators into thinking this is a normal File
						return File.prototype;
					}
				}
			);
		}
	});

	// Read the request body asyncronously so it doesn't stall
	void (async () => {
		let has_more = true;
		while (has_more) {
			const chunk = await get_chunk(chunks.length);
			has_more = !!chunk;
		}
	})();

	return { data, meta, form_data: null };
}

/** @implements {File} */
class LazyFile {
	/** @type {(index: number) => Promise<Uint8Array<ArrayBuffer> | undefined>} */
	#get_chunk;
	/** @type {number} */
	#offset;
	/**
	 * @param {string} name
	 * @param {string} type
	 * @param {number} size
	 * @param {number} last_modified
	 * @param {(index: number) => Promise<Uint8Array<ArrayBuffer> | undefined>} get_chunk
	 * @param {number} offset
	 */
	constructor(name, type, size, last_modified, get_chunk, offset) {
		this.name = name;
		this.type = type;
		this.size = size;
		this.lastModified = last_modified;
		this.webkitRelativePath = '';
		this.#get_chunk = get_chunk;
		this.#offset = offset;

		// TODO - hacky, required for private members to be accessed on proxy
		this.arrayBuffer = this.arrayBuffer.bind(this);
		this.bytes = this.bytes.bind(this);
		this.slice = this.slice.bind(this);
		this.stream = this.stream.bind(this);
		this.text = this.text.bind(this);
	}
	/** @type {ArrayBuffer | undefined} */
	#buffer;
	async arrayBuffer() {
		this.#buffer ??= await new Response(this.stream()).arrayBuffer();
		return this.#buffer;
	}
	async bytes() {
		return new Uint8Array(await this.arrayBuffer());
	}
	/**
	 * @param {number=} start
	 * @param {number=} end
	 * @param {string=} contentType
	 */
	slice(start = 0, end = this.size, contentType = this.type) {
		// https://github.com/nodejs/node/blob/a5f3cd8cb5ba9e7911d93c5fd3ebc6d781220dd8/lib/internal/blob.js#L240
		if (start < 0) {
			start = Math.max(this.size + start, 0);
		} else {
			start = Math.min(start, this.size);
		}

		if (end < 0) {
			end = Math.max(this.size + end, 0);
		} else {
			end = Math.min(end, this.size);
		}
		const size = Math.max(end - start, 0);
		const file = new LazyFile(
			this.name,
			contentType,
			size,
			this.lastModified,
			this.#get_chunk,
			this.#offset + start
		);

		return file;
	}
	stream() {
		let cursor = 0;
		let chunk_index = 0;
		return new ReadableStream({
			start: async (controller) => {
				let chunk_start = 0;
				let start_chunk = null;
				for (chunk_index = 0; ; chunk_index++) {
					const chunk = await this.#get_chunk(chunk_index);
					if (!chunk) return null;

					const chunk_end = chunk_start + chunk.byteLength;
					// If this chunk contains the target offset
					if (this.#offset >= chunk_start && this.#offset < chunk_end) {
						start_chunk = chunk;
						break;
					}
					chunk_start = chunk_end;
				}
				// If the buffer is completely contained in one chunk, do a subarray
				if (this.#offset + this.size <= chunk_start + start_chunk.byteLength) {
					controller.enqueue(
						start_chunk.subarray(this.#offset - chunk_start, this.#offset + this.size - chunk_start)
					);
					controller.close();
				} else {
					controller.enqueue(start_chunk.subarray(this.#offset - chunk_start));
					cursor = start_chunk.byteLength - this.#offset + chunk_start;
				}
			},
			pull: async (controller) => {
				chunk_index++;
				let chunk = await this.#get_chunk(chunk_index);
				if (!chunk) {
					controller.error('Could not deserialize binary form: incomplete file data');
					controller.close();
					return;
				}
				if (chunk.byteLength > this.size - cursor) {
					chunk = chunk.subarray(0, this.size - cursor);
				}
				controller.enqueue(chunk);
				cursor += chunk.byteLength;
				if (cursor >= this.size) {
					controller.close();
				}
			}
		});
	}
	async text() {
		return text_decoder.decode(await this.arrayBuffer());
	}
}

const path_regex = /^[a-zA-Z_$]\w*(\.[a-zA-Z_$]\w*|\[\d+\])*$/;

/**
 * @param {string} path
 */
export function split_path(path) {
	if (!path_regex.test(path)) {
		throw new Error(`Invalid path ${path}`);
	}

	return path.split(/\.|\[|\]/).filter(Boolean);
}

/**
 * Check if a property key is dangerous and could lead to prototype pollution
 * @param {string} key
 */
function check_prototype_pollution(key) {
	if (key === '__proto__' || key === 'constructor' || key === 'prototype') {
		throw new Error(
			`Invalid key "${key}"` +
				(DEV ? ': This key is not allowed to prevent prototype pollution.' : '')
		);
	}
}

/**
 * Sets a value in a nested object using an array of keys, mutating the original object.
 * @param {Record<string, any>} object
 * @param {string[]} keys
 * @param {any} value
 */
export function deep_set(object, keys, value) {
	let current = object;

	for (let i = 0; i < keys.length - 1; i += 1) {
		const key = keys[i];

		check_prototype_pollution(key);

		const is_array = /^\d+$/.test(keys[i + 1]);
		const exists = key in current;
		const inner = current[key];

		if (exists && is_array !== Array.isArray(inner)) {
			throw new Error(`Invalid array key ${keys[i + 1]}`);
		}

		if (!exists) {
			current[key] = is_array ? [] : {};
		}

		current = current[key];
	}

	const final_key = keys[keys.length - 1];
	check_prototype_pollution(final_key);
	current[final_key] = value;
}

/**
 * @param {StandardSchemaV1.Issue} issue
 * @param {boolean} server Whether this issue came from server validation
 */
export function normalize_issue(issue, server = false) {
	/** @type {InternalRemoteFormIssue} */
	const normalized = { name: '', path: [], message: issue.message, server };

	if (issue.path !== undefined) {
		let name = '';

		for (const segment of issue.path) {
			const key = /** @type {string | number} */ (
				typeof segment === 'object' ? segment.key : segment
			);

			normalized.path.push(key);

			if (typeof key === 'number') {
				name += `[${key}]`;
			} else if (typeof key === 'string') {
				name += name === '' ? key : '.' + key;
			}
		}

		normalized.name = name;
	}

	return normalized;
}

/**
 * @param {InternalRemoteFormIssue[]} issues
 */
export function flatten_issues(issues) {
	/** @type {Record<string, InternalRemoteFormIssue[]>} */
	const result = {};

	for (const issue of issues) {
		(result.$ ??= []).push(issue);

		let name = '';

		if (issue.path !== undefined) {
			for (const key of issue.path) {
				if (typeof key === 'number') {
					name += `[${key}]`;
				} else if (typeof key === 'string') {
					name += name === '' ? key : '.' + key;
				}

				(result[name] ??= []).push(issue);
			}
		}
	}

	return result;
}

/**
 * Gets a nested value from an object using a path array
 * @param {Record<string, any>} object
 * @param {(string | number)[]} path
 * @returns {any}
 */

export function deep_get(object, path) {
	let current = object;
	for (const key of path) {
		if (current == null || typeof current !== 'object') {
			return current;
		}
		current = current[key];
	}
	return current;
}

/**
 * Creates a proxy-based field accessor for form data
 * @param {any} target - Function or empty POJO
 * @param {() => Record<string, any>} get_input - Function to get current input data
 * @param {(path: (string | number)[], value: any) => void} set_input - Function to set input data
 * @param {() => Record<string, InternalRemoteFormIssue[]>} get_issues - Function to get current issues
 * @param {(string | number)[]} path - Current access path
 * @returns {any} Proxy object with name(), value(), and issues() methods
 */
export function create_field_proxy(target, get_input, set_input, get_issues, path = []) {
	const get_value = () => {
		return deep_get(get_input(), path);
	};

	return new Proxy(target, {
		get(target, prop) {
			if (typeof prop === 'symbol') return target[prop];

			// Handle array access like jobs[0]
			if (/^\d+$/.test(prop)) {
				return create_field_proxy({}, get_input, set_input, get_issues, [
					...path,
					parseInt(prop, 10)
				]);
			}

			const key = build_path_string(path);

			if (prop === 'set') {
				const set_func = function (/** @type {any} */ newValue) {
					set_input(path, newValue);
					return newValue;
				};
				return create_field_proxy(set_func, get_input, set_input, get_issues, [...path, prop]);
			}

			if (prop === 'value') {
				return create_field_proxy(get_value, get_input, set_input, get_issues, [...path, prop]);
			}

			if (prop === 'issues' || prop === 'allIssues') {
				const issues_func = () => {
					const all_issues = get_issues()[key === '' ? '$' : key];

					if (prop === 'allIssues') {
						return all_issues?.map((issue) => ({
							path: issue.path,
							message: issue.message
						}));
					}

					return all_issues
						?.filter((issue) => issue.name === key)
						?.map((issue) => ({
							path: issue.path,
							message: issue.message
						}));
				};

				return create_field_proxy(issues_func, get_input, set_input, get_issues, [...path, prop]);
			}

			if (prop === 'as') {
				/**
				 * @param {string} type
				 * @param {string} [input_value]
				 */
				const as_func = (type, input_value) => {
					const is_array =
						type === 'file multiple' ||
						type === 'select multiple' ||
						(type === 'checkbox' && typeof input_value === 'string');

					const prefix =
						type === 'number' || type === 'range'
							? 'n:'
							: type === 'checkbox' && !is_array
								? 'b:'
								: '';

					// Base properties for all input types
					/** @type {Record<string, any>} */
					const base_props = {
						name: prefix + key + (is_array ? '[]' : ''),
						get 'aria-invalid'() {
							const issues = get_issues();
							return key in issues ? 'true' : undefined;
						}
					};

					// Add type attribute only for non-text inputs and non-select elements
					if (type !== 'text' && type !== 'select' && type !== 'select multiple') {
						base_props.type = type === 'file multiple' ? 'file' : type;
					}

					// Handle submit and hidden inputs
					if (type === 'submit' || type === 'hidden') {
						if (DEV) {
							if (!input_value) {
								throw new Error(`\`${type}\` inputs must have a value`);
							}
						}

						return Object.defineProperties(base_props, {
							value: { value: input_value, enumerable: true }
						});
					}

					// Handle select inputs
					if (type === 'select' || type === 'select multiple') {
						return Object.defineProperties(base_props, {
							multiple: { value: is_array, enumerable: true },
							value: {
								enumerable: true,
								get() {
									return get_value();
								}
							}
						});
					}

					// Handle checkbox inputs
					if (type === 'checkbox' || type === 'radio') {
						if (DEV) {
							if (type === 'radio' && !input_value) {
								throw new Error('Radio inputs must have a value');
							}

							if (type === 'checkbox' && is_array && !input_value) {
								throw new Error('Checkbox array inputs must have a value');
							}
						}

						return Object.defineProperties(base_props, {
							value: { value: input_value ?? 'on', enumerable: true },
							checked: {
								enumerable: true,
								get() {
									const value = get_value();

									if (type === 'radio') {
										return value === input_value;
									}

									if (is_array) {
										return (value ?? []).includes(input_value);
									}

									return value;
								}
							}
						});
					}

					// Handle file inputs
					if (type === 'file' || type === 'file multiple') {
						return Object.defineProperties(base_props, {
							multiple: { value: is_array, enumerable: true },
							files: {
								enumerable: true,
								get() {
									const value = get_value();

									// Convert File/File[] to FileList-like object
									if (value instanceof File) {
										// In browsers, we can create a proper FileList using DataTransfer
										if (typeof DataTransfer !== 'undefined') {
											const fileList = new DataTransfer();
											fileList.items.add(value);
											return fileList.files;
										}
										// Fallback for environments without DataTransfer
										return { 0: value, length: 1 };
									}

									if (Array.isArray(value) && value.every((f) => f instanceof File)) {
										if (typeof DataTransfer !== 'undefined') {
											const fileList = new DataTransfer();
											value.forEach((file) => fileList.items.add(file));
											return fileList.files;
										}
										// Fallback for environments without DataTransfer
										/** @type {any} */
										const fileListLike = { length: value.length };
										value.forEach((file, index) => {
											fileListLike[index] = file;
										});
										return fileListLike;
									}

									return null;
								}
							}
						});
					}

					// Handle all other input types (text, number, etc.)
					return Object.defineProperties(base_props, {
						value: {
							enumerable: true,
							get() {
								const value = get_value();
								return value != null ? String(value) : '';
							}
						}
					});
				};

				return create_field_proxy(as_func, get_input, set_input, get_issues, [...path, 'as']);
			}

			// Handle property access (nested fields)
			return create_field_proxy({}, get_input, set_input, get_issues, [...path, prop]);
		}
	});
}

/**
 * Builds a path string from an array of path segments
 * @param {(string | number)[]} path
 * @returns {string}
 */
export function build_path_string(path) {
	let result = '';

	for (const segment of path) {
		if (typeof segment === 'number') {
			result += `[${segment}]`;
		} else {
			result += result === '' ? segment : '.' + segment;
		}
	}

	return result;
}

/**
 * @param {RemoteForm<any, any>} instance
 * @deprecated remove in 3.0
 */
export function throw_on_old_property_access(instance) {
	Object.defineProperty(instance, 'field', {
		value: (/** @type {string} */ name) => {
			const new_name = name.endsWith('[]') ? name.slice(0, -2) : name;
			throw new Error(
				`\`form.field\` has been removed: Instead of \`<input name={form.field('${name}')} />\` do \`<input {...form.fields.${new_name}.as(type)} />\``
			);
		}
	});

	for (const property of ['input', 'issues']) {
		Object.defineProperty(instance, property, {
			get() {
				const new_name = property === 'issues' ? 'issues' : 'value';
				return new Proxy(
					{},
					{
						get(_, prop) {
							const prop_string = typeof prop === 'string' ? prop : String(prop);
							const old =
								prop_string.includes('[') || prop_string.includes('.')
									? `['${prop_string}']`
									: `.${prop_string}`;
							const replacement = `.${prop_string}.${new_name}()`;
							throw new Error(
								`\`form.${property}\` has been removed: Instead of \`form.${property}${old}\` write \`form.fields${replacement}\``
							);
						}
					}
				);
			}
		});
	}
}
