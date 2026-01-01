<script>import { fade } from "svelte/transition";
import { tocStore, tocActiveId } from "./stores.js";
export let inactive = "opacity-60 hover:opacity-100";
export let active = "text-primary-500";
export let activeId = "";
export let indentStyles = {
  h2: "",
  h3: "ml-4",
  h4: "ml-8",
  h5: "ml-12",
  h6: "ml-16"
};
export let regionLead = "font-bold";
export let regionList = "";
export let regionListItem = "";
export let regionAnchor = "";
const cBase = "space-y-4";
const cList = "space-y-2";
const cListItem = "block";
const cAnchor = "";
$:
  reactiveActiveId = $tocActiveId ? $tocActiveId : activeId.replace("#", "");
$:
  classesBase = `${cBase} ${$$props.class ?? ""}`;
$:
  classesList = `${cList} ${regionList}`;
$:
  classesListItem = `${cListItem} ${regionListItem}`;
$:
  classesAnchor = `${cAnchor} ${regionAnchor}`;
</script>

{#if $tocStore.length}
	<nav class="toc {classesBase}" data-testid="toc" transition:fade|local={{ duration: 100 }}>
		<!-- Slot: Default (title) -->
		<div class={regionLead}>
			<slot>Table of Contents</slot>
		</div>
		<!-- List -->
		<ul class="toc-list {classesList}">
			{#each $tocStore as tocHeading}
				<li class="toc-list-item {classesListItem} {indentStyles[tocHeading.element]}">
					<a
						href="#{tocHeading.id}"
						class="toc-anchor {classesAnchor} {tocHeading.id === reactiveActiveId ? active : inactive}"
						on:click={() => {
							reactiveActiveId = tocHeading.id;
						}}
					>
						{tocHeading.text}
					</a>
				</li>
			{/each}
		</ul>
	</nav>
{/if}
