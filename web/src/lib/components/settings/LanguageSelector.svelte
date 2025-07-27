<script lang="ts">
  import { Label } from "$lib/components/ui/label";
  import { Select } from "$lib/components/ui/select";
  import { languageStore } from '$lib/store/language-store';

  let selectedLanguage = $languageStore;
  $: languageStore.set(selectedLanguage);

  const languages = [
    { code: '', name: 'Default' },
    { code: 'en', name: 'English' },
    { code: 'fr', name: 'French' },
    { code: 'es', name: 'Spanish' },
    { code: 'de', name: 'German' },
    { code: 'zh', name: 'Chinese' },
    { code: 'ja', name: 'Japanese' }
  ];
</script>

<div class="flex flex-col gap-2">
  <Label>Language</Label>
  <Select bind:value={selectedLanguage}>
    {#each languages as lang}
      <option value={lang.code}>{lang.name}</option>
    {/each}
  </Select>
</div>