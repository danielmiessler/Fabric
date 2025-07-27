<script lang="ts">
  import { page } from '$app/stores';
</script>

<div class="container flex min-h-[400px] flex-col items-center justify-center">
	<div class="text-center">
		<h1 class="text-4xl font-bold tracking-tighter sm:text-5xl">
			{$page.status}: {$page.error?.message || 'Something went wrong'}
		</h1>
		<p class="mt-4 text-muted-foreground">
			{#if $page.status === 404}
				Sorry, we couldn't find the post you're looking for.
			{:else}
				An error occurred while loading the post. Please try again later.
			{/if}
		</p>
		<div class="mt-8">
			<a
				href="/posts"
				class="inline-flex items-center justify-center rounded-md bg-primary px-8 py-2 text-sm font-medium text-primary-foreground ring-offset-background transition-colors hover:bg-primary/90 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2"
			>
				Back to Posts
			</a>
		</div>
	</div>
</div>
