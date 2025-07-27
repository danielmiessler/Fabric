import { browser } from '$app/environment';
import { GlobalWorkerOptions } from 'pdfjs-dist';

// Set up the worker source location - point to static file in public directory
const workerSrc = '/pdf.worker.min.mjs';

// Configure the worker options only on the client side
if (browser) {
  GlobalWorkerOptions.workerSrc = workerSrc;
}

// Export the configuration
export default {
  initialize: () => {
    if (browser) {
      console.log('PDF.js worker initialized at', workerSrc);
    }
  }
};

