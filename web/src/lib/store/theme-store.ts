import { writable } from 'svelte/store';

// Skeleton 5 renamed two of these themes: "skeleton" is now "legacy" and
// "gold-nouveau" is now "nouveau". Each name needs an @import in app.css.
const themes = [
  'my-custom-theme',
  'legacy',
  'modern',
  'crimson',
  'nouveau',
  'hamlindigo',
  'vintage',
  'seafoam',
  'sahara',
  'rocket'
];

type ThemeType = typeof themes[number];

function createThemeStore() {
  const { subscribe, set, update } = writable<ThemeType>('legacy');

  return {
    subscribe,
    cycleTheme: () => update(currentTheme => {
      const currentIndex = themes.indexOf(currentTheme);
      const nextIndex = (currentIndex + 1) % themes.length;
      const newTheme = themes[nextIndex];
      
      if (typeof document !== 'undefined') {
        document.body.setAttribute('data-theme', newTheme);
        localStorage.setItem('theme', newTheme);
      }
      return newTheme;
    }),
    setTheme: (theme: ThemeType) => {
      set(theme);
      if (typeof document !== 'undefined') {
        document.body.setAttribute('data-theme', theme);
        localStorage.setItem('theme', theme);
      }
    },
    initTheme: () => {
      if (typeof document !== 'undefined') {
        const savedTheme = localStorage.getItem('theme') as ThemeType;
        if (savedTheme && themes.includes(savedTheme)) {
          set(savedTheme);
          document.body.setAttribute('data-theme', savedTheme);
        }
      }
    }
  };
}

export const theme = createThemeStore();
export const cycleTheme = theme.cycleTheme;
export const initTheme = theme.initTheme;
