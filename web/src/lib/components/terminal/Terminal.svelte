<script lang="ts">
  import { onMount } from 'svelte';
  // import { fade } from 'svelte/transition';
  import { goto } from '$app/navigation';

  let mounted = false;
  let currentCommand = '';
  let commandHistory: string[] = [];
  let showCursor = true;

  let terminalContent = '';
  let typing = false;
  
  const pages = {
    home: 'Welcome to Fabric\n\nType `help` to see available commands.',
    about: 'About Fabric',
    chat: 'Enter `chat` to start a chat session.',
    posts: 'Enter `posts` to view blog posts.',
    tags: 'Enter `tags` to view tags.',
    contact: 'Enter `contact` to view contact info.',
    help: `Available commands:
- help: Show this help message
- about: Navigate to About page
- chat: Start a chat session
- posts: View all blog posts
- tags: Browse content by tags
- contact: Get in touch
- clear: Clear the terminal
- ls: List available pages`,
  };

  // Simulate typing effect
  async function typeContent(content: string) {
    typing = true;
    terminalContent = '';
    for (const char of content) {
      terminalContent += char;
      await new Promise(resolve => setTimeout(resolve, 20));
    }
    typing = false;
  }

  function handleCommand(cmd: string) {
    commandHistory = [...commandHistory, cmd];
    
    switch (cmd) {
      case 'clear':
        terminalContent = '';
        break;
      case 'help':
        typeContent(pages.help);
        break;
      case 'about':
        goto('/about');
        break;
      case 'chat':
        goto('/chat');
        break;
      case 'posts': 
        goto('/posts');
        break;
      case 'tags':
        goto('/tags');
        break;
      case 'contact':
        goto('/contact');
        break;
      case 'ls':
        typeContent(Object.keys(pages).join('\n'));
        break;
      default:
          const page = cmd.slice(3);
          if (pages[page]) {
            typeContent(pages[page]);
          } else {
            typeContent(`Error: Page '${page}' not found`);
          }
        }
    }

  function handleKeydown(event: KeyboardEvent) {
    if (typing) return;

    if (event.key === 'Enter') {
      handleCommand(currentCommand.trim());
      currentCommand = '';
    }
  }

  onMount(() => {
    mounted = true;
    setInterval(() => {
      showCursor = !showCursor;
    }, 500);

    // Initial content
    typeContent(pages.home);
  });
</script>

<div class="pt-2 pb-8 px-4">
  <div class="container mx-auto max-w-4xl">
    <div class="terminal-window backdrop-blur-sm">
      <!-- Terminal header -->
      <div class="terminal-header flex items-center gap-2 px-4 py-2 border-b border-gray-700/50">
        <div class="flex gap-2">
          <div class="w-3 h-3 rounded-full bg-red-500/80"></div>
          <div class="w-3 h-3 rounded-full bg-yellow-500/80"></div>
          <div class="w-3 h-3 rounded-full bg-green-500/80"></div>
        </div>
        <span class="text-sm text-gray-400 ml-2">me@localhost</span>
      </div>

      <div class="p-6">
        <div class="mb-4 whitespace-pre-wrap terminal-text leading-relaxed">{terminalContent}</div>

        <!-- Command input -->
        {#if mounted}
          <div class="flex items-center">
            <span class="mr-2 terminal-prompt font-bold">$</span>
            {#if showCursor}
              <span class="animate-blink terminal-text">▋</span>
            {/if}
            <input
              type="text"
              bind:value={currentCommand}
              on:keydown={handleKeydown}
              class="flex-1 bg-transparent border-none outline-none terminal-text"
              placeholder="Type a command..."
            />
          </div>
        {/if}
      </div>
    </div>
  </div>
</div>

<style>
  .terminal-window {
    @apply rounded-lg border border-gray-700/50 bg-gray-900/95 shadow-2xl;
    box-shadow: 0 0 60px -15px rgba(0, 0, 0, 0.3);
  }

  .terminal-text {
    @apply font-mono text-green-400/90;
  }

  .terminal-prompt {
    @apply text-blue-400/90;
  }

  input::placeholder {
    @apply text-gray-600;
  }

  .animate-blink {
    animation: blink 1s step-end infinite;
    flex-col: 1; 

  }

  @keyframes blink {
    50% {
      opacity: 0;
    }
  }

  /*::-webkit-scrollbar {*/
  /*  @apply w-2;*/
  /*}*/
  /**/
  /*::-webkit-scrollbar-track {*/
  /*  @apply bg-gray-800/50 rounded-full;*/
  /*}*/
  /**/
  /*::-webkit-scrollbar-thumb {*/
  /*  @apply bg-gray-600/50 rounded-full hover:bg-gray-500/50 transition-colors;*/
  /*}*/
</style>
