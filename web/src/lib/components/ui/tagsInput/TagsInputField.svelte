<script lang="ts">
  // Skeleton 5 renamed InputChip to TagsInput and made it a set of parts that
  // each caller assembles. This component holds that assembly in one place and
  // keeps the small API that the InputChip call sites used.
  import { TagsInput } from '@skeletonlabs/skeleton-svelte';

  interface Props {
    value?: string[];
    name?: string;
    placeholder?: string;
    /** Return false to refuse the tag that the person typed. */
    // eslint.config.js uses the core no-unused-vars rule for Svelte components,
    // and that rule reads the parameter name of a function type as a variable.
    // eslint-disable-next-line no-unused-vars
    validation?: (value: string) => boolean;
    class?: string;
  }

  let {
    value = $bindable([]),
    name = 'tags',
    placeholder = '',
    validation,
    class: classes = ''
  }: Props = $props();
</script>

<TagsInput
  {name}
  {value}
  onValueChange={(details) => (value = details.value)}
  validate={validation ? (details) => validation(details.inputValue) : undefined}
  class={classes}
>
  <TagsInput.Context>
    {#snippet children(api)}
      <TagsInput.Control class="input flex flex-wrap items-center gap-2">
        {#each api().value as tag, index (tag)}
          <TagsInput.Item {index} value={tag} class="chip preset-filled-primary-500">
            <TagsInput.ItemPreview class="flex items-center gap-1">
              <TagsInput.ItemText />
              <TagsInput.ItemDeleteTrigger class="cursor-pointer opacity-70 hover:opacity-100">
                &times;
              </TagsInput.ItemDeleteTrigger>
            </TagsInput.ItemPreview>
            <TagsInput.ItemInput />
          </TagsInput.Item>
        {/each}
        <TagsInput.Input {placeholder} class="flex-auto border-none bg-transparent outline-none" />
      </TagsInput.Control>
      <TagsInput.HiddenInput />
    {/snippet}
  </TagsInput.Context>
</TagsInput>
