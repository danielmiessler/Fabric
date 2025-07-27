import { writable } from 'svelte/store';

// Load favorites from localStorage if available
const storedFavorites = typeof localStorage !== 'undefined' 
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
        
        // Save to localStorage
        if (typeof localStorage !== 'undefined') {
          localStorage.setItem('favoritePatterns', JSON.stringify(newFavorites));
        }
        
        return newFavorites;
      });
    },
    reset: () => {
      set([]);
      if (typeof localStorage !== 'undefined') {
        localStorage.removeItem('favoritePatterns');
      }
    }
  };
};

export const favorites = createFavoritesStore();