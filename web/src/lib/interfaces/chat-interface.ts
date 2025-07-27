export type MessageRole = 'system' | 'user' | 'assistant';
export type ResponseFormat = 'markdown' | 'mermaid' | 'plain' | 'loading';
export type ResponseType = 'content' | 'error' | 'complete';

export interface ChatPrompt {
  userInput: string;
  systemPrompt: string;
  model: string;
  patternName?: string;
  strategyName?: string; // Optional strategy name to prepend strategy prompt
  variables?: { [key: string]: string }; // Pattern variables
}

export interface ChatConfig {
  temperature: number;
  top_p: number;
  frequency_penalty: number;
  presence_penalty: number;
}

export interface ChatRequest {
  prompts: ChatPrompt[];
  messages: Message[];
  temperature: number;
  top_p: number;
  frequency_penalty: number;
  presence_penalty: number;
  language?: string;
}

export interface Message {
  role: MessageRole;
  content: string;
  format?: ResponseFormat;
}

export interface ChatState {
  messages: Message[];
  isStreaming: boolean;
}

export interface StreamResponse {
  type: ResponseType;
  format: ResponseFormat;
  content: string;
}

export interface ChatError {
  code: string;
  message: string;
  details?: unknown;
}
