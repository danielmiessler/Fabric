<script>import TreeViewItem from "./TreeViewItem.svelte";
import RecursiveTreeViewItem from "./RecursiveTreeViewItem.svelte";
import { createEventDispatcher, getContext, onMount } from "svelte";
export let nodes = [];
export let expandedNodes = [];
export let disabledNodes = [];
export let checkedNodes = [];
export let indeterminateNodes = [];
let selection = getContext("selection");
let multiple = getContext("multiple");
let relational = getContext("relational");
let group = multiple ? [] : "";
let name = "";
const dispatch = createEventDispatcher();
function toggleNode(node, open) {
  if (!node.children?.length)
    return;
  if (open) {
    if (!expandedNodes.includes(node.id)) {
      expandedNodes.push(node.id);
      expandedNodes = expandedNodes;
    }
  } else {
    if (expandedNodes.includes(node.id)) {
      expandedNodes.splice(expandedNodes.indexOf(node.id), 1);
      expandedNodes = expandedNodes;
    }
  }
}
function checkNode(node, checked, indeterminate) {
  if (checked) {
    if (!checkedNodes.includes(node.id)) {
      checkedNodes.push(node.id);
      checkedNodes = checkedNodes;
    }
    if (!indeterminate && indeterminateNodes.includes(node.id)) {
      indeterminateNodes.splice(indeterminateNodes.indexOf(node.id), 1);
      indeterminateNodes = indeterminateNodes;
    }
  } else {
    if (checkedNodes.includes(node.id)) {
      checkedNodes.splice(checkedNodes.indexOf(node.id), 1);
      checkedNodes = checkedNodes;
    }
    if (indeterminate && !indeterminateNodes.includes(node.id)) {
      indeterminateNodes.push(node.id);
      indeterminateNodes = indeterminateNodes;
    } else if (!indeterminate && indeterminateNodes.includes(node.id)) {
      indeterminateNodes.splice(indeterminateNodes.indexOf(node.id), 1);
      indeterminateNodes = indeterminateNodes;
    }
  }
}
if (selection) {
  if (multiple) {
    nodes.forEach((node) => {
      if (!Array.isArray(group))
        return;
      if (checkedNodes.includes(node.id) && !group.includes(node.id)) {
        group.push(node.id);
      }
    });
    group = group;
  } else {
    nodes.forEach((node) => {
      if (checkedNodes.includes(node.id) && group !== node.id) {
        group = node.id;
      }
    });
  }
}
onMount(async () => {
  if (selection) {
    name = String(Math.random());
    if (!relational)
      treeItems = [];
  }
});
export let treeItems = [];
let children = [];
</script>

{#if nodes && nodes.length > 0}
	{#each nodes as node, i}
		<TreeViewItem
			bind:this={treeItems[i]}
			bind:children={children[i]}
			bind:group
			bind:name
			bind:value={node.id}
			hideLead={!node.lead}
			hideChildren={!node.children || node.children.length === 0}
			open={expandedNodes.includes(node.id)}
			disabled={disabledNodes.includes(node.id)}
			checked={checkedNodes.includes(node.id)}
			indeterminate={indeterminateNodes.includes(node.id)}
			on:toggle={(e) => toggleNode(node, e.detail.open)}
			on:groupChange={(e) => checkNode(node, e.detail.checked, e.detail.indeterminate)}
			on:click={() =>
				dispatch('click', {
					id: node.id
				})}
			on:toggle={() => {
				dispatch('toggle', {
					id: node.id
				});
			}}
		>
			{#if typeof node.content === 'string'}
				{@html node.content}
			{:else}
				<svelte:component this={node.content} {...node.contentProps} />
			{/if}
			<svelte:fragment slot="lead">
				{#if typeof node.lead === 'string'}
					{@html node.lead}
				{:else}
					<svelte:component this={node.lead} {...node.leadProps} />
				{/if}
			</svelte:fragment>
			<svelte:fragment slot="children">
				<RecursiveTreeViewItem
					nodes={node.children}
					bind:expandedNodes
					bind:disabledNodes
					bind:checkedNodes
					bind:indeterminateNodes
					bind:treeItems={children[i]}
					on:click={(e) =>
						dispatch('click', {
							id: e.detail.id
						})}
					on:toggle={(e) =>
						dispatch('toggle', {
							id: e.detail.id
						})}
				/>
			</svelte:fragment>
		</TreeViewItem>
	{/each}
{/if}
