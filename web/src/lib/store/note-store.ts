import { writable, get } from 'svelte/store';
import { browser } from '$app/environment';
import { buildNoteFrontmatter, serializeMarkdownNote } from '$lib/utils/frontmatter';

interface NoteState {
  content: string;
  lastSaved: Date | null;
  isDirty: boolean;
}

function createNoteStore() {
  const { subscribe, set, update } = writable<NoteState>({
      content: '',
      lastSaved: null,
      isDirty: false
  });

  const generateUniqueFilename = () => {
      const now = new Date();
      const date = now.toISOString().split('T')[0];
      const time = now.toISOString().split('T')[1]
          .replace(/:/g, '-')
          .split('.')[0];
      return `${date}-${time}.md`;
  };

  const saveToFile = async (content: string) => {
      if (!browser) return;

      const filename = generateUniqueFilename();
      const frontmatter = buildNoteFrontmatter(content);
      const fileContent = serializeMarkdownNote(frontmatter, content);

      const response = await fetch('/notes', {
          method: 'POST',
          headers: {
              'Content-Type': 'application/json',
          },
          body: JSON.stringify({
              filename,
              content: fileContent
          })
      });

      if (!response.ok) {
          throw new Error(await response.text());
      }

      return filename;
  };

  return {
      subscribe,
      updateContent: (content: string) => update(state => ({
          ...state,
          content,
          isDirty: true
      })),
      save: async () => {
          const state = get({ subscribe });
          const filename = await saveToFile(state.content);

          update(state => ({
              ...state,
              lastSaved: new Date(),
              isDirty: false
          }));

          return filename;
      },
      reset: () => set({
          content: '',
          lastSaved: null,
          isDirty: false
      })
  };
}

export const noteStore = createNoteStore();
