import type { Message } from "./chat-interface";

export interface Session {
  Name: string;
  Message: string | Message[];
  Session: string | object;
}
