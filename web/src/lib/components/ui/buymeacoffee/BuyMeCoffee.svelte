<script lang="ts">
  export let url: string = 'https://www.buymeacoffee.com/johnconnor.sec';
  export let text: string = 'Buy me a coffee';
</script>

<a
	href={url}
	target="_blank"
	rel="noopener noreferrer"
	class="text-sm px-3 py-1.5 btn variant-filled-tertiary hover:variant-filled-secondary transition-all duration-200 flex items-center gap-2"
>
	<svg
		xmlns="http://www.w3.org/2000/svg"
		width="20"
		height="20"
		viewBox="0 0 24 24"
		fill="none"
		stroke="currentColor"
		stroke-width="2"
		stroke-linecap="round"
		stroke-linejoin="round"
		class="transition-transform duration-200 group-hover:rotate-12"
	>
		<path d="M17 8h1a4 4 0 1 1 0 8h-1" />
		<path d="M3 8h14v9a4 4 0 0 1-4 4H7a4 4 0 0 1-4-4Z" />
		<line x1="6" y1="2" x2="6" y2="4" />
		<line x1="10" y1="2" x2="10" y2="4" />
		<line x1="14" y1="2" x2="14" y2="4" />
	</svg>
	{text}
</a>
