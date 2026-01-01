import { DEV } from 'esm-env';
import { getContext } from 'svelte';

function context() {
	return getContext('__request__');
}

/** @param {string} name */
function context_dev(name) {
	try {
		return context();
	} catch {
		throw new Error(
			`Can only read '${name}' on the server during rendering (not in e.g. \`load\` functions), as it is bound to the current request via component context. This prevents state from leaking between users.` +
				'For more information, see https://svelte.dev/docs/kit/state-management#avoid-shared-state-on-the-server'
		);
	}
}

export const page = {
	get data() {
		return (DEV ? context_dev('page.data') : context()).page.data;
	},
	get error() {
		return (DEV ? context_dev('page.error') : context()).page.error;
	},
	get form() {
		return (DEV ? context_dev('page.form') : context()).page.form;
	},
	get params() {
		return (DEV ? context_dev('page.params') : context()).page.params;
	},
	get route() {
		return (DEV ? context_dev('page.route') : context()).page.route;
	},
	get state() {
		return (DEV ? context_dev('page.state') : context()).page.state;
	},
	get status() {
		return (DEV ? context_dev('page.status') : context()).page.status;
	},
	get url() {
		return (DEV ? context_dev('page.url') : context()).page.url;
	}
};

export const navigating = {
	from: null,
	to: null,
	type: null,
	willUnload: null,
	delta: null,
	complete: null
};

export const updated = {
	get current() {
		return false;
	},
	check: () => {
		throw new Error('Can only call updated.check() in the browser');
	}
};
