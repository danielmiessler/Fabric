import { fileURLToPath } from 'node:url';
import { Worker, parentPort } from 'node:worker_threads';
import process from 'node:process';

/**
 * Runs a task in a subprocess so any dangling stuff gets killed upon completion.
 * The subprocess needs to be the file `forked` is called in, and `forked` needs to be called eagerly at the top level.
 * @template T
 * @template U
 * @param {string} module `import.meta.url` of the file
 * @param {(opts: T) => Promise<U>} callback The function that is invoked in the subprocess
 * @returns {(opts: T) => Promise<U>} A function that when called starts the subprocess
 */
export function forked(module, callback) {
	if (process.env.SVELTEKIT_FORK && parentPort) {
		parentPort.on(
			'message',
			/** @param {any} data */ async (data) => {
				if (data?.type === 'args' && data.module === module) {
					parentPort?.postMessage({
						type: 'result',
						module,
						payload: await callback(data.payload)
					});
				}
			}
		);

		parentPort.postMessage({ type: 'ready', module });
	}

	/**
	 * @param {T} opts
	 * @returns {Promise<U>}
	 */
	return function (opts) {
		return new Promise((fulfil, reject) => {
			const worker = new Worker(fileURLToPath(module), {
				env: {
					...process.env,
					SVELTEKIT_FORK: 'true'
				}
			});

			worker.on(
				'message',
				/** @param {any} data */ (data) => {
					if (data?.type === 'ready' && data.module === module) {
						worker.postMessage({
							type: 'args',
							module,
							payload: opts
						});
					}

					if (data?.type === 'result' && data.module === module) {
						worker.unref();
						fulfil(data.payload);
					}
				}
			);

			worker.on('exit', (code) => {
				if (code) {
					reject(new Error(`Failed with code ${code}`));
				}
			});
		});
	};
}
