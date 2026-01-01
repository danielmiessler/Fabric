import {
	page as _page,
	navigating as _navigating,
	updated as _updated
} from '../../client/state.svelte.js';
import { stores } from '../../client/client.js';

export const page = {
	get data() {
		return _page.data;
	},
	get error() {
		return _page.error;
	},
	get form() {
		return _page.form;
	},
	get params() {
		return _page.params;
	},
	get route() {
		return _page.route;
	},
	get state() {
		return _page.state;
	},
	get status() {
		return _page.status;
	},
	get url() {
		return _page.url;
	}
};

export const navigating = {
	get from() {
		return _navigating.current ? _navigating.current.from : null;
	},
	get to() {
		return _navigating.current ? _navigating.current.to : null;
	},
	get type() {
		return _navigating.current ? _navigating.current.type : null;
	},
	get willUnload() {
		return _navigating.current ? _navigating.current.willUnload : null;
	},
	get delta() {
		return _navigating.current ? _navigating.current.delta : null;
	},
	get complete() {
		return _navigating.current ? _navigating.current.complete : null;
	}
};

Object.defineProperty(navigating, 'current', {
	get() {
		// between 2.12.0 and 2.12.1 `navigating.current` existed
		throw new Error('Replace navigating.current.<prop> with navigating.<prop>');
	}
});

export const updated = {
	get current() {
		return _updated.current;
	},
	check: stores.updated.check
};
