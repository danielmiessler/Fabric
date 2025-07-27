import type {
  ChatRequest,
  StreamResponse,
  ChatError as IChatError,
  ChatPrompt
} from '$lib/interfaces/chat-interface';
import { get } from 'svelte/store';
import { modelConfig } from '$lib/store/model-store';
import { systemPrompt, selectedPatternName, patternVariables } from '$lib/store/pattern-store';
import { chatConfig } from '$lib/store/chat-config';
import { messageStore } from '$lib/store/chat-store';
import { languageStore } from '$lib/store/language-store';
import { selectedStrategy } from '$lib/store/strategy-store';

class LanguageValidator {
  constructor(private targetLanguage: string) {}

  enforceLanguage(content: string): string {
    if (this.targetLanguage === 'en') return content;
    return `[Language: ${this.targetLanguage}]\n${content}`;
  }
}

export class ChatError extends Error implements IChatError {
  constructor(
    message: string,
    public readonly code: string = 'CHAT_ERROR',
    public readonly details?: unknown
  ) {
    super(message);
    this.name = 'ChatError';
  }
}

export class ChatService {
  private validator: LanguageValidator;

  constructor() {
    this.validator = new LanguageValidator(get(languageStore));
  }

  private async fetchStream(request: ChatRequest): Promise<ReadableStream<StreamResponse>> {
    try {
      console.log('\n=== ChatService Request Start ===');
      console.log('1. Request details:', {
        language: get(languageStore),
        pattern: get(selectedPatternName),
        promptCount: request.prompts?.length,
        messageCount: request.messages?.length
      });
      // NEW: Log the full payload before sending to backend
      console.log('Final ChatRequest payload:', JSON.stringify(request, null, 2));

      const response = await fetch('/api/chat', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(request),
      });

      if (!response.ok) {
        throw new ChatError(`HTTP error! status: ${response.status}`, 'HTTP_ERROR', { status: response.status });
      }

      const reader = response.body?.getReader();
      if (!reader) {
        throw new ChatError('Response body is null', 'NULL_RESPONSE');
      }

      return this.createMessageStream(reader);
    } catch (error) {
      if (error instanceof ChatError) throw error;
      throw new ChatError('Failed to fetch chat stream', 'FETCH_ERROR', error);
    }
  }

  private cleanPatternOutput(content: string): string {
    // Remove markdown fence if present
    let cleaned = content.replace(/^```markdown\n/, '');
    cleaned = cleaned.replace(/\n```$/, '');

    // Existing cleaning
    cleaned = cleaned.replace(/^# OUTPUT\s*\n/, '');
    cleaned = cleaned.replace(/^\s*\n/, '');
    cleaned = cleaned.replace(/\n\s*$/, '');
    cleaned = cleaned.replace(/^#\s+([A-Z]+):/gm, '$1:');
    cleaned = cleaned.replace(/^#\s+([A-Z]+)\s*$/gm, '$1');
    cleaned = cleaned.trim();
    cleaned = cleaned.replace(/\n{3,}/g, '\n\n');
    return cleaned;
  }

  private createMessageStream(reader: ReadableStreamDefaultReader<Uint8Array>): ReadableStream<StreamResponse> {
      let buffer = '';
      const cleanPatternOutput = this.cleanPatternOutput.bind(this);
      const language = get(languageStore);
      const validator = new LanguageValidator(language);

      const processResponse = (response: StreamResponse) => {
          const pattern = get(selectedPatternName);

          if (pattern) {
              response.content = cleanPatternOutput(response.content);
              // Simplified format determination - always markdown unless mermaid
              const isMermaid = [
                  'graph TD', 'gantt', 'flowchart',
                  'sequenceDiagram', 'classDiagram', 'stateDiagram'
              ].some(starter => response.content.trim().startsWith(starter));

              response.format = isMermaid ? 'mermaid' : 'markdown';
          }

          if (response.type === 'content') {
              response.content = validator.enforceLanguage(response.content);
          }

          return response;
      };
      return new ReadableStream({
          async start(controller) {
              try {
                  while (true) {
                      const { done, value } = await reader.read();
                      if (done) break;

                      buffer += new TextDecoder().decode(value);
                      const messages = buffer.split('\n\n').filter(msg => msg.startsWith('data: '));

                      if (messages.length > 1) {
                          buffer = messages.pop() || '';
                          for (const msg of messages) {
                              try {
                                  let response = JSON.parse(msg.slice(6)) as StreamResponse;
                                  response = processResponse(response);
                                  controller.enqueue(response);
                              } catch (parseError) {
                                  console.error('Error parsing stream message:', parseError);
                              }
                          }
                      }
                  }

                  if (buffer.startsWith('data: ')) {
                      try {
                          let response = JSON.parse(buffer.slice(6)) as StreamResponse;
                          response = processResponse(response);
                          controller.enqueue(response);
                      } catch (parseError) {
                          console.error('Error parsing final message:', parseError);
                      }
                  }
              } catch (error) {
                  controller.error(new ChatError('Stream processing error', 'STREAM_ERROR', error));
              } finally {
                  reader.releaseLock();
                  controller.close();
              }
          },
          cancel() {
              reader.cancel();
          }
      });
  }

  private createChatPrompt(userInput: string, systemPromptText?: string): ChatPrompt {
    const config = get(modelConfig);
    const language = get(languageStore);

    const languageInstruction = language !== 'en'
        ? `You MUST respond in ${language} language. All output must be in ${language}. `
        // ? `You MUST respond in ${language} language. ALL output, including section headers, titles, and formatting, MUST be translated into ${language}.  It is CRITICAL that you translate ALL headers, such as SUMMARY, IDEAS, QUOTES, TAKEAWAYS, MAIN POINTS, etc., into ${language}. Maintain markdown formatting in the response. Do not output any English headers.`
        : '';

    const finalSystemPrompt = languageInstruction + (systemPromptText ?? get(systemPrompt));

    const finalUserInput = language !== 'en'
        ? `${userInput}\n\nIMPORTANT: Respond in ${language} language only.`
        : userInput;

    return {
        userInput: finalUserInput,
        systemPrompt: finalSystemPrompt,
        model: config.model,
        patternName: get(selectedPatternName),
        strategyName: get(selectedStrategy), // Add selected strategy to prompt
        variables: get(patternVariables) // Add pattern variables
    };
}

  public async createChatRequest(userInput: string, systemPromptText?: string, isPattern: boolean = false): Promise<ChatRequest> {
    const prompt = this.createChatPrompt(userInput, systemPromptText);
    const config = get(chatConfig);
    const language = get(languageStore);

    return {
      prompts: [prompt],
      messages: [],
      language: language, // Add language at the top level for backend compatibility
      ...config
    };
  }

  public async streamPattern(userInput: string, systemPromptText?: string): Promise<ReadableStream<StreamResponse>> {
    const request = await this.createChatRequest(userInput, systemPromptText, true);
    return this.fetchStream(request);
  }

  public async streamChat(userInput: string, systemPromptText?: string): Promise<ReadableStream<StreamResponse>> {
    const request = await this.createChatRequest(userInput, systemPromptText);
    return this.fetchStream(request);
  }

  public async processStream(
    stream: ReadableStream<StreamResponse>,
    onContent: (content: string, response?: StreamResponse) => void,
    onError: (error: Error) => void
  ): Promise<void> {
    const reader = stream.getReader();

    try {
      while (true) {
        const { done, value } = await reader.read();
        if (done) break;

        if (value.type === 'error') {
          throw new ChatError(value.content, 'STREAM_CONTENT_ERROR');
        }

        if (value.type === 'content') {
          onContent(value.content, value);
        }
      }
    } catch (error) {
      onError(error instanceof ChatError ? error : new ChatError('Stream processing error', 'STREAM_ERROR', error));
    } finally {
      reader.releaseLock();

    }
  }
}
