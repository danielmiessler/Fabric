<script lang="ts">
  import { Button } from "$lib/components/ui/button";
  import { Textarea } from "$lib/components/ui/textarea";
  import { messageStore } from '$lib/store/chat-store';
  import { systemPrompt, selectedPatternName } from '$lib/store/pattern-store';
  import { toastStore } from '$lib/store/toast-store';
  import { Paperclip, Send, FileCheck } from 'lucide-svelte';
  import { onMount } from 'svelte';
  import { get } from 'svelte/store';
  import { getTranscript } from '$lib/services/transcriptService';
  import { ChatService } from '$lib/services/ChatService';
  // import { obsidianSettings } from '$lib/store/obsidian-store';
  import { languageStore } from '$lib/store/language-store';
  import { obsidianSettings, updateObsidianSettings } from '$lib/store/obsidian-store';
  import { PdfConversionService } from '$lib/services/PdfConversionService';
  import { formatErrorMessage } from '$lib/utils/error-message';
  import { readFileContent } from './read-file-content';
  
  const pdfService = new PdfConversionService();
  



  const chatService = new ChatService();
  let userInput = "";
  let isYouTubeURL = false;
  let files: FileList | undefined = undefined;
  let uploadedFiles: string[] = [];
  let fileContents: string[] = [];
  let isProcessingFiles = false;
  let isFileIndicatorVisible = false; // Add new variable
  let fileButtonKey = false; // Add new key variable for FileButton
  function detectYouTubeURL(input: string): boolean {
    const youtubePattern = /(?:https?:\/\/)?(?:www\.)?(?:youtube\.com|youtu\.be)/i;
    const isYoutube = youtubePattern.test(input);
    if (isYoutube) {
      console.log('YouTube URL detected:', input);
      console.log('Current system prompt:', $systemPrompt?.length);
      console.log('Selected pattern:', $selectedPatternName);
    }
    return isYoutube;
  }

  function handleInput(event: Event) {
    console.log('\n=== Handle Input ===');
    const target = event.target as HTMLTextAreaElement;
    userInput = target.value;
    
    const currentLanguage = get(languageStore);
    
    const languageQualifiers = {
      '--en': 'en',
      '--fr': 'fr',
      '--es': 'es',
      '--de': 'de',
      '--zh': 'zh',
      '--ja': 'ja'
    };

    let detectedLang = '';
    for (const [qualifier, lang] of Object.entries(languageQualifiers)) {
      if (userInput.includes(qualifier)) {
        detectedLang = lang;
        languageStore.set(lang);
        userInput = userInput.replace(new RegExp(`${qualifier}\\s*`), '');
        break;
      }
    }

    console.log('2. Language state:', {
      previousLanguage: currentLanguage,
      currentLanguage: get(languageStore),
      detectedOverride: detectedLang,
      inputAfterLangRemoval: userInput
    });

    isYouTubeURL = detectYouTubeURL(userInput);
    console.log('3. URL detection:', {
      isYouTube: isYouTubeURL,
      pattern: $selectedPatternName,
      systemPromptLength: $systemPrompt?.length
    });
  }

  async function handleFileUpload(e: Event) {
  uploadedFiles = []; // Clear uploadedFiles at the beginning
  isFileIndicatorVisible = false;
  if (!files || files.length === 0) return;

  if (uploadedFiles.length >= 5 || (uploadedFiles.length + files.length) > 5) {
    toastStore.error('Maximum 5 files allowed');
    return;
  }

  isProcessingFiles = true;
  try {
    // Add processing indicator to message store
    messageStore.update(messages => [...messages, {
      role: 'system',
      content: 'Processing files...',
      format: 'loading'
    }]);

    for (let i = 0; i < files.length && uploadedFiles.length < 5; i++) {
      const file = files[i];
      const content = await readFileContent(file, pdfService, toastStore.warning);
      fileContents.push(content);
      uploadedFiles = [...uploadedFiles, file.name];
      isFileIndicatorVisible = true;
      
      // Update processing status per file
      messageStore.update(messages => {
        const newMessages = [...messages];
        const lastMessage = newMessages[newMessages.length - 1];
        if (lastMessage?.format === 'loading') {
          lastMessage.content = `Processing ${file.name} (${file.type})...`;
        }
        return newMessages;
      });
    }

    // Remove processing message on completion
    messageStore.update(messages => 
      messages.filter(m => m.format !== 'loading')
    );

  } catch (error) {
    toastStore.error('Error processing files: ' + (error as Error).message);
    
    // Clean up processing message on error
    messageStore.update(messages => 
      messages.filter(m => m.format !== 'loading')
    );
  } finally {
    isProcessingFiles = false;
  }
}


  async function saveToObsidian(content: string) {
    if (!$obsidianSettings.saveToObsidian) {
      console.log('Obsidian saving is disabled');
      return;
    }
    
    if (!$obsidianSettings.noteName) {
      toastStore.error('Please enter a note name in Obsidian settings');
      return;
    }

    if (!$selectedPatternName) {
      toastStore.error('No pattern selected');
      return;
    }

    if (!content) {
      toastStore.error('No content to save');
      return;
    }

    try {
      const response = await fetch('/obsidian', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          pattern: $selectedPatternName,
          noteName: $obsidianSettings.noteName,
          content
        })
      });

      const responseData = await response.json();
      
      if (!response.ok) {
        throw new Error(responseData.error || 'Failed to save to Obsidian');
      }
      // Add this after successful save
      updateObsidianSettings({ 
      saveToObsidian: false,  // Reset the save flag
      noteName: ''           // Clear the note name
      });
      toastStore.success(responseData.message || `Saved to Obsidian: ${responseData.fileName}`);
    } catch (error) {
      console.error('Failed to save to Obsidian:', error);
      toastStore.error(error instanceof Error ? error.message : 'Failed to save to Obsidian');
    }
  }

  // Centralized language instruction logic in ChatService.ts; YouTube flow now passes plain transcript and system prompt
  function extractYouTubeURLs(input: string): string[] {
      const youtubePattern = /(?:https?:\/\/)?(?:www\.)?(?:youtube\.com\/watch\?v=[\w-]+(?:&[^\s]*)?|youtu\.be\/[\w-]+(?:\?[^\s]*)?)/gi;
      return input.match(youtubePattern) || [];
  }

  async function replaceYouTubeURLsWithTranscripts(input: string): Promise<string> {
      const urls = extractYouTubeURLs(input);
      if (urls.length === 0) return input;

      let result = input;
      for (const url of urls) {
          const { transcript } = await getTranscript(url);
          result = result.replace(url, `[YouTube Transcript]:\n${transcript}`);
      }
      return result;
  }

  async function handleSubmit() {
  if (!userInput.trim()) return;

  try {
    console.log('\n=== Submit Handler Start ===');

    // Store the user input before any processing
    const inputText = userInput.trim();
    console.log('Captured user input:', inputText);

    // Add the user message to the UI first
    messageStore.update(messages => [...messages, {
      role: 'user',
      content: inputText
    }]);

    // Add loading indicator
    messageStore.update(messages => [...messages, {
      role: 'system',
      content: isYouTubeURL ? 'Processing YouTube video...' : 'Processing...',
      format: 'loading'
    }]);

    // Clear input fields
    userInput = "";
    const hadYouTubeURL = isYouTubeURL;
    isYouTubeURL = false;
    const filesForProcessing = [...uploadedFiles];
    const contentsForProcessing = [...fileContents];
    uploadedFiles = [];
    fileContents = [];
    isFileIndicatorVisible = false;
    fileButtonKey = !fileButtonKey;

    // If the message contains YouTube URLs, replace them with transcripts
    let processedText = inputText;
    if (hadYouTubeURL) {
      console.log('Replacing YouTube URLs with transcripts');
      processedText = await replaceYouTubeURLsWithTranscripts(inputText);
    }

    // Prepare content with file attachments if any
    const contentWithFiles = contentsForProcessing.length > 0
      ? `${processedText}\n\nFile Contents (${filesForProcessing.map(f => f.endsWith('.pdf') ? 'PDF' : 'Text').join(', ')}):\n${contentsForProcessing.join('\n\n---\n\n')}`
      : processedText;

    // Get the enhanced prompt
    const enhancedPrompt = contentsForProcessing.length > 0
      ? `${$systemPrompt}\nAnalyze and process the provided content according to these instructions.`
      : $systemPrompt;
    
    console.log('Content to send:', {
      text: contentWithFiles.substring(0, 100) + '...',
      length: contentWithFiles.length,
      hasFiles: contentsForProcessing.length > 0
    });
    
    try {
      // Get the chat stream
      const stream = await chatService.streamChat(contentWithFiles, enhancedPrompt);
      
      // Process the stream
      await chatService.processStream(
        stream,
        (content, response) => {
          messageStore.update(messages => {
            const newMessages = [...messages];
            // Remove the loading message
            const loadingIndex = newMessages.findIndex(m => m.format === 'loading');
            if (loadingIndex !== -1) {
              newMessages.splice(loadingIndex, 1);
            }

            const lastMessage = newMessages[newMessages.length - 1];
            if (lastMessage?.role === 'assistant') {
              lastMessage.content += content;
              lastMessage.format = response?.format;
            } else {
              newMessages.push({
                role: 'assistant',
                content,
                format: response?.format
              });
            }
            return newMessages;
          });
        },

        (error) => {
          // Make sure to remove loading message on error
          messageStore.update(messages =>
            messages.filter(m => m.format !== 'loading')
          );
          console.error('Stream processing error:', error);

          const message = formatErrorMessage(error);

          // Show the error in the chat, where it stays for the person to read.
          messageStore.update(messages => [...messages, {
            role: 'system',
            content: message,
            format: 'plain'
          }]);
          // And as a toast, so that a failure is visible even when the chat is
          // scrolled away from the end.
          toastStore.error(message);
        }
      );
    } catch (error) {
      // Make sure to remove loading message on error
      messageStore.update(messages => 
        messages.filter(m => m.format !== 'loading')
      );
      throw error; // Re-throw to be caught by the outer try/catch
    }
  } catch (error) {
    console.error('Chat submission error:', error);
    
    // Make sure to remove loading message on error (redundant but safe)
    messageStore.update(messages => 
      messages.filter(m => m.format !== 'loading')
    );
    
    const message = formatErrorMessage(error);

    // Show the error in the chat and as a toast, as the stream handler does.
    messageStore.update(messages => [...messages, {
      role: 'system',
      content: message,
      format: 'plain'
    }]);
    toastStore.error(message);
  } finally {
    // As a final safety measure, ensure loading message is removed
    messageStore.update(messages => 
      messages.filter(m => m.format !== 'loading')
    );
  }
}

  function handleKeydown(event: KeyboardEvent) {
    if (event.key === 'Enter' && !event.shiftKey) {
      event.preventDefault();
      handleSubmit();
    }
  }

  onMount(() => {
    console.log('ChatInput mounted, current system prompt:', $systemPrompt);
  });
</script>

<div class="h-full flex flex-col p-2">
  <div class="relative flex-1 min-h-0 bg-primary-800/30 rounded-lg">
    <Textarea
      bind:value={userInput}
      on:input={handleInput}
      on:keydown={handleKeydown}
      placeholder="Enter your message (YouTube URLs will be automatically processed)..."
      class="w-full h-full resize-none bg-transparent border-none text-sm focus:ring-0 transition-colors p-3 pb-[48px]"
    />
    <div class="absolute bottom-3 right-3 flex items-center gap-2">
      <div class="flex items-center gap-2">
        {#if isFileIndicatorVisible}
          <span class="text-xs text-white/70">
            {uploadedFiles.length} file{uploadedFiles.length > 1 ? 's' : ''} attached
          </span>
        {/if}
      {#key fileButtonKey}
        <!-- Skeleton 5 replaced FileButton with FileUpload, which draws a drop
          zone and reports files through a callback. A label with a hidden file
          input keeps both the appearance and the change handler of the button
          that was here before. -->
        <label
          class="btn-icon preset-tonal inline-flex h-10 w-10 cursor-pointer items-center justify-center rounded-full bg-primary-800/30 transition-colors hover:bg-primary-800/50"
          class:pointer-events-none={isProcessingFiles || uploadedFiles.length >= 5}
          class:opacity-50={isProcessingFiles || uploadedFiles.length >= 5}
        >
          <input
            type="file"
            name="file-upload"
            class="hidden"
            bind:files
            on:change={handleFileUpload}
            disabled={isProcessingFiles || uploadedFiles.length >= 5}
          />
          <Paperclip class="w-5 h-5" />
        </label>
      {/key}
        <Button
          type="button"
          variant="ghost"
          size="icon"
          name="send"
          on:click={handleSubmit}
          disabled={isProcessingFiles || !userInput.trim()}
          class="h-10 w-10 bg-primary-800/30 hover:bg-primary-800/50 rounded-full transition-colors disabled:opacity-30"
        >
          <Send class="w-5 h-5" />
        </Button>
      </div>
    </div>
  </div>
</div>

<style>
  :global(textarea) {
    scrollbar-width: thin;
    scrollbar-color: rgba(255, 255, 255, 0.2) transparent;
  }

  :global(textarea::-webkit-scrollbar) {
    width: 6px;
  }

  :global(textarea::-webkit-scrollbar-track) {
    background: transparent;
  }

  :global(textarea::-webkit-scrollbar-thumb) {
    background-color: rgba(255, 255, 255, 0.2);
    border-radius: 3px;
  }

  :global(textarea::-webkit-scrollbar-thumb:hover) {
    background-color: rgba(255, 255, 255, 0.3);
  }

  :global(textarea::selection) {
    background-color: rgba(255, 255, 255, 0.1);
  }
</style>
