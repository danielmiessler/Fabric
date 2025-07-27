import { writable } from 'svelte/store';
import { browser } from '$app/environment';

const storedLanguage = browser ? localStorage.getItem('selectedLanguage') || 'en' : 'en';
const languageStore = writable<string>(storedLanguage);

if (browser) {
    languageStore.subscribe(value => {
        localStorage.setItem('selectedLanguage', value);
    });
}

export { languageStore };