<script lang="ts">
  import type { PageData } from './$types';
  import PostContent from '$lib/components/posts/PostContent.svelte';
  import PostMeta from '$lib/components/posts/PostMeta.svelte';

  export let data: PageData;
  console.log('Page data:', data);
</script>

<PostContent post={{
    content: data.content,
    metadata: data.metadata,
    slug: data.slug
}} />
