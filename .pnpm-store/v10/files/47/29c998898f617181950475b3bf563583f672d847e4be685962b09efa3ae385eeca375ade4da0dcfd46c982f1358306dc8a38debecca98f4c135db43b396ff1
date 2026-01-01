import * as devalue from 'devalue';
import { DEV } from 'esm-env';
import { invalidateAll } from './navigation.js';
import { app, applyAction } from '../client/client.js';

export { applyAction };

/**
 * Use this function to deserialize the response from a form submission.
 * Usage:
 *
 * ```js
 * import { deserialize } from '$app/forms';
 *
 * async function handleSubmit(event) {
 *   const response = await fetch('/form?/action', {
 *     method: 'POST',
 *     body: new FormData(event.target)
 *   });
 *
 *   const result = deserialize(await response.text());
 *   // ...
 * }
 * ```
 * @template {Record<string, unknown> | undefined} Success
 * @template {Record<string, unknown> | undefined} Failure
 * @param {string} result
 * @returns {import('@sveltejs/kit').ActionResult<Success, Failure>}
 */
export function deserialize(result) {
	const parsed = JSON.parse(result);

	if (parsed.data) {
		parsed.data = devalue.parse(parsed.data, app.decoders);
	}

	return parsed;
}

/**
 * Shallow clone an element, so that we can access e.g. `form.action` without worrying
 * that someone has added an `<input name="action">` (https://github.com/sveltejs/kit/issues/7593)
 * @template {HTMLElement} T
 * @param {T} element
 * @returns {T}
 */
function clone(element) {
	return /** @type {T} */ (HTMLElement.prototype.cloneNode.call(element));
}

/**
 * This action enhances a `<form>` element that otherwise would work without JavaScript.
 *
 * The `submit` function is called upon submission with the given FormData and the `action` that should be triggered.
 * If `cancel` is called, the form will not be submitted.
 * You can use the abort `controller` to cancel the submission in case another one starts.
 * If a function is returned, that function is called with the response from the server.
 * If nothing is returned, the fallback will be used.
 *
 * If this function or its return value isn't set, it
 * - falls back to updating the `form` prop with the returned data if the action is on the same page as the form
 * - updates `page.status`
 * - resets the `<form>` element and invalidates all data in case of successful submission with no redirect response
 * - redirects in case of a redirect response
 * - redirects to the nearest error page in case of an unexpected error
 *
 * If you provide a custom function with a callback and want to use the default behavior, invoke `update` in your callback.
 * It accepts an options object
 * - `reset: false` if you don't want the `<form>` values to be reset after a successful submission
 * - `invalidateAll: false` if you don't want the action to call `invalidateAll` after submission
 * @template {Record<string, unknown> | undefined} Success
 * @template {Record<string, unknown> | undefined} Failure
 * @param {HTMLFormElement} form_element The form element
 * @param {import('@sveltejs/kit').SubmitFunction<Success, Failure>} submit Submit callback
 */
export function enhance(form_element, submit = () => {}) {
	if (DEV && clone(form_element).method !== 'post') {
		throw new Error('use:enhance can only be used on <form> fields with method="POST"');
	}

	/**
	 * @param {{
	 *   action: URL;
	 *   invalidateAll?: boolean;
	 *   result: import('@sveltejs/kit').ActionResult;
	 *   reset?: boolean
	 * }} opts
	 */
	const fallback_callback = async ({
		action,
		result,
		reset = true,
		invalidateAll: shouldInvalidateAll = true
	}) => {
		if (result.type === 'success') {
			if (reset) {
				// We call reset from the prototype to avoid DOM clobbering
				HTMLFormElement.prototype.reset.call(form_element);
			}
			if (shouldInvalidateAll) {
				await invalidateAll();
			}
		}

		// For success/failure results, only apply action if it belongs to the
		// current page, otherwise `form` will be updated erroneously
		if (
			location.origin + location.pathname === action.origin + action.pathname ||
			result.type === 'redirect' ||
			result.type === 'error'
		) {
			await applyAction(result);
		}
	};

	/** @param {SubmitEvent} event */
	async function handle_submit(event) {
		const method = event.submitter?.hasAttribute('formmethod')
			? /** @type {HTMLButtonElement | HTMLInputElement} */ (event.submitter).formMethod
			: clone(form_element).method;
		if (method !== 'post') return;

		event.preventDefault();

		const action = new URL(
			// We can't do submitter.formAction directly because that property is always set
			event.submitter?.hasAttribute('formaction')
				? /** @type {HTMLButtonElement | HTMLInputElement} */ (event.submitter).formAction
				: clone(form_element).action
		);

		const enctype = event.submitter?.hasAttribute('formenctype')
			? /** @type {HTMLButtonElement | HTMLInputElement} */ (event.submitter).formEnctype
			: clone(form_element).enctype;

		const form_data = new FormData(form_element);

		if (DEV && enctype !== 'multipart/form-data') {
			for (const value of form_data.values()) {
				if (value instanceof File) {
					throw new Error(
						'Your form contains <input type="file"> fields, but is missing the necessary `enctype="multipart/form-data"` attribute. This will lead to inconsistent behavior between enhanced and native forms. For more details, see https://github.com/sveltejs/kit/issues/9819.'
					);
				}
			}
		}

		const submitter_name = event.submitter?.getAttribute('name');
		if (submitter_name) {
			form_data.append(submitter_name, event.submitter?.getAttribute('value') ?? '');
		}

		const controller = new AbortController();

		let cancelled = false;
		const cancel = () => (cancelled = true);

		const callback =
			(await submit({
				action,
				cancel,
				controller,
				formData: form_data,
				formElement: form_element,
				submitter: event.submitter
			})) ?? fallback_callback;
		if (cancelled) return;

		/** @type {import('@sveltejs/kit').ActionResult} */
		let result;

		try {
			const headers = new Headers({
				accept: 'application/json',
				'x-sveltekit-action': 'true'
			});

			// do not explicitly set the `Content-Type` header when sending `FormData`
			// or else it will interfere with the browser's header setting
			// see https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest_API/Using_FormData_Objects#sect4
			if (enctype !== 'multipart/form-data') {
				headers.set(
					'Content-Type',
					/^(:?application\/x-www-form-urlencoded|text\/plain)$/.test(enctype)
						? enctype
						: 'application/x-www-form-urlencoded'
				);
			}

			// @ts-expect-error `URLSearchParams(form_data)` is kosher, but typescript doesn't know that
			const body = enctype === 'multipart/form-data' ? form_data : new URLSearchParams(form_data);

			const response = await fetch(action, {
				method: 'POST',
				headers,
				cache: 'no-store',
				body,
				signal: controller.signal
			});

			result = deserialize(await response.text());
			if (result.type === 'error') result.status = response.status;
		} catch (error) {
			if (/** @type {any} */ (error)?.name === 'AbortError') return;
			result = { type: 'error', error };
		}

		await callback({
			action,
			formData: form_data,
			formElement: form_element,
			update: (opts) =>
				fallback_callback({
					action,
					result,
					reset: opts?.reset,
					invalidateAll: opts?.invalidateAll
				}),
			// @ts-expect-error generic constraints stuff we don't care about
			result
		});
	}

	// @ts-expect-error
	HTMLFormElement.prototype.addEventListener.call(form_element, 'submit', handle_submit);

	return {
		destroy() {
			// @ts-expect-error
			HTMLFormElement.prototype.removeEventListener.call(form_element, 'submit', handle_submit);
		}
	};
}
