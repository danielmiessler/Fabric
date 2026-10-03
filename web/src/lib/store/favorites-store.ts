import { writable } from 'svelte/store';
import { browser } from '$app/environment';

const storedFavorites = browser
  ? JSON.parse(localStorage.getItem('favoritePatterns') || '[]')
  : [];

const createFavoritesStore = () => {
  const { subscribe, set, update } = writable<string[]>(storedFavorites);
  
  return {
    subscribe,
    toggleFavorite: (patternName: string) => {
      update(favorites => {
        const newFavorites = favorites.includes(patternName)
          ? favorites.filter(name => name !== patternName)
          : [...favorites, patternName];
        
        if (browser) {
          localStorage.setItem('favoritePatterns', JSON.stringify(newFavorites));
        }
        
        return newFavorites;
      });
    },
    reset: () => {
      set([]);
      if (browser) {
        localStorage.removeItem('favoritePatterns');
      }
    }
  };
};

export const favorites = createFavoritesStore();
