# Language Support Implementation

## Overview
The language support allows switching between languages using qualifiers (--fr, --en) in the chat input. The implementation is simple and effective, working at multiple layers of the application.

## Components

### 1. Language Store (language-store.ts)
```typescript
// Manages language state
export const languageStore = writable<string>('');
```

### 2. Chat Input (ChatInput.svelte)
- Detects language qualifiers in user input
- Updates language store
- Strips qualifier from message
```typescript
// Language qualifier handling
if (qualifier === 'fr') {
  languageStore.set('fr');
  userInput = userInput.replace(/--fr\s*/, '');
} else if (qualifier === 'en') {
  languageStore.set('en');
  userInput = userInput.replace(/--en\s*/, '');
}

// After sending message
try {
  await sendMessage(userInput);
  languageStore.set('en'); // Reset to default after send
} catch (error) {
  console.error('Failed to send message:', error);
}
```

### 3. Chat Service (ChatService.ts)
- Adds language instruction to prompts
- Defaults to English if no language specified
```typescript
const language = get(languageStore) || 'en';
const languageInstruction = language !== 'en' 
  ? `. Please use the language '${language}' for the output.` 
  : '';
const fullInput = userInput + languageInstruction;
```

### 4. Global Settings UI (Chat.svelte)
```typescript
// Language selector in Global Settings
<div class="flex flex-col gap-2">
  <Label>Language</Label>
  <Select bind:value={selectedLanguage}>
    <option value="">Default</option>
    <option value="en">English</option>
    <option value="fr">French</option>
  </Select>
</div>

// Script section
let selectedLanguage = $languageStore;
$: languageStore.set(selectedLanguage);
```

## How It Works

1. User Input:
   - User types message with language qualifier (e.g., "--fr Hello")
   - ChatInput detects qualifier and updates language store
   - Qualifier is stripped from message
   - OR user selects language from Global Settings dropdown

2. Request Processing:
   - ChatService gets language from store
   - Adds language instruction to prompt
   - Sends to backend

3. Response:
   - AI responds in requested language
   - Response is displayed without modification
   - Language store is reset to English after message is sent

## Usage Examples

1. English (Default):
```
User: What is the weather?
AI: The weather information...
```

2. French:
```
User: --fr What is the weather?
AI: Voici les informations météo...
```

3. Using Global Settings:
```
1. Select "French" from language dropdown
2. Type: What is the weather?
3. AI responds in French
4. Language resets to English after response
```

## Implementation Notes

1. Simple Design:
   - No complex language detection
   - No translation layer
   - Direct instruction to AI

2. Stateful:
   - Language persists until changed
   - Resets to English on page refresh
   - Resets to English after each message

3. Extensible:
   - Easy to add new languages
   - Just add new qualifiers and store values
   - Update Global Settings dropdown options

4. Error Handling:
   - Invalid qualifiers are ignored
   - Unknown languages default to English
   - Store reset on error to prevent state issues

## Best Practices

1. Always reset language after message:
```typescript
// Reset stores after successful send
languageStore.set('en');
```

2. Default to English:
```typescript
const language = get(languageStore) || 'en';
```

3. Clear language instruction:
```typescript
const languageInstruction = language !== 'en' 
  ? `. Please use the language '${language}' for the output.` 
  : '';
```

4. Handle UI State:
```typescript
// In Chat.svelte
let selectedLanguage = $languageStore;
$: {
  languageStore.set(selectedLanguage);
  // Update UI immediately when store changes
  selectedLanguage = $languageStore;
}