import { createStorageAPI } from '$lib/api/base';
import type { Session } from '$lib/interfaces/session-interface';
import type { Message } from '$lib/interfaces/chat-interface';
import { get, writable } from 'svelte/store';
import { openFileDialog, readFileAsJson, saveToFile } from '../utils/file-utils';
import { toastService } from '../services/toast-service';

export const sessions = writable<Session[]>([]);

export const sessionAPI = {
  ...createStorageAPI<Session>('sessions'),

  async loadSessions() {
    try {
      const response = await fetch(`/api/sessions/names`);
      const sessionNames: string[] = await response.json();

      if (!sessionNames) {
        sessions.set([]);
        return [];
      }

      const sessionPromises = sessionNames.map(async (name: string) => {
        try {
          const response = await fetch(`/api/sessions/${name}`);
          const data = await response.json();
          return {
            Name: name,
            Message: Array.isArray(data.Message) ? data.Message : [],
            Session: data.Session
          };
        } catch (error) {
          console.error(`Error loading session ${name}:`, error);
          return {
            Name: name,
            Message: [],
            Session: ""
          };
        }
      });

      const sessionsData = await Promise.all(sessionPromises);
      sessions.set(sessionsData);
      return sessionsData;
    } catch (error) {
      console.error('Error loading sessions:', error);
      sessions.set([]);
      return [];
    }
  },


  selectSession(sessionName: string) {
    const allSessions = get(sessions);
    const selectedSession = allSessions.find(session => session.Name === sessionName);
    if (selectedSession) {
      sessions.set([selectedSession]);
    } else {
      sessions.set([]);
    }
  },

  async exportToFile(messages: Message[]) {
    try {
      await saveToFile(messages, 'session-history.json');
      toastService.success('Session exported successfully');
    } catch (error) {
      toastService.error('Failed to export session');
      throw error;
    }
  },

  async importFromFile(): Promise<Message[]> {
    try {
      const file = await openFileDialog('.json');
      if (!file) {
        throw new Error('No file selected');
      }

      const content = await readFileAsJson<Message[]>(file);
      if (!Array.isArray(content)) {
        throw new Error('Invalid session file format');
      }

      toastService.success('Session imported successfully');
      return content;
    } catch (error) {
      toastService.error(error instanceof Error ? error.message : 'Failed to import session');
      throw error;
    }
  },

  async loadSessionMessages(sessionName: string): Promise<Message[]> {
    try {
      const response = await fetch(`/api/sessions/${sessionName}`);
      if (!response.ok) {
        throw new Error(`Failed to load session: ${response.statusText}`);
      }
      const data = await response.json();
      const messages = Array.isArray(data.Message) ? data.Message : [];
      return messages;
    } catch (error) {
      console.error(`Error loading session messages for ${sessionName}:`, error);
      throw error;
    }
  }
};
