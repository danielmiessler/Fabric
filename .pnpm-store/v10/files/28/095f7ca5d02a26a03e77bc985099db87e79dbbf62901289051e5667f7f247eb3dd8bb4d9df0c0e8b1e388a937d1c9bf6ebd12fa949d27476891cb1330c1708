/** @import { Tracer, Span, SpanContext } from '@opentelemetry/api' */

/**
 * Tracer implementation that does nothing (null object).
 * @type {Tracer}
 */
export const noop_tracer = {
	/**
	 * @returns {Span}
	 */
	startSpan() {
		return noop_span;
	},

	/**
	 * @param {unknown} _name
	 * @param {unknown} arg_1
	 * @param {unknown} [arg_2]
	 * @param {Function} [arg_3]
	 * @returns {unknown}
	 */
	startActiveSpan(_name, arg_1, arg_2, arg_3) {
		if (typeof arg_1 === 'function') {
			return arg_1(noop_span);
		}
		if (typeof arg_2 === 'function') {
			return arg_2(noop_span);
		}
		if (typeof arg_3 === 'function') {
			return arg_3(noop_span);
		}
	}
};

/**
 * @type {Span}
 */
export const noop_span = {
	spanContext() {
		return noop_span_context;
	},
	setAttribute() {
		return this;
	},
	setAttributes() {
		return this;
	},
	addEvent() {
		return this;
	},
	setStatus() {
		return this;
	},
	updateName() {
		return this;
	},
	end() {
		return this;
	},
	isRecording() {
		return false;
	},
	recordException() {
		return this;
	},
	addLink() {
		return this;
	},
	addLinks() {
		return this;
	}
};

/**
 * @type {SpanContext}
 */
const noop_span_context = {
	traceId: '',
	spanId: '',
	traceFlags: 0
};
