<script>import { createEventDispatcher } from "svelte";
import { tableA11y } from "./actions.js";
const dispatch = createEventDispatcher();
export let source;
export let interactive = false;
export let element = "table";
export let text = "";
export let color = "";
export let regionHead = "";
export let regionHeadCell = "";
export let regionBody = "";
export let regionCell = "";
export let regionFoot = "";
export let regionFootCell = "";
function onRowClick(event, rowIndex) {
  if (!interactive)
    return;
  event.preventDefault();
  event.stopPropagation();
  const rowMetaData = source.meta ? source.meta[rowIndex] : source.body[rowIndex];
  dispatch("selected", rowMetaData);
}
function onRowKeydown(event, rowIndex) {
  if (["Enter", "Space"].includes(event.code))
    onRowClick(event, rowIndex);
}
$:
  classesBase = `${$$props.class || ""}`;
$:
  classesTable = `${element} ${text} ${color}`;
</script>

<div class="table-container {classesBase}">
	<!-- Table -->
	<!-- prettier-ignore -->
	<table
		class="{classesTable}"
		class:table-interactive={interactive}
		role={interactive ? "grid" : "table"}
		use:tableA11y
	>
		<!-- on:keydown={(e) => onTableKeydown(elemTable, e)} -->
		<!-- Head -->
		<thead class="table-head {regionHead}">
			<tr>
				{#each source.head as heading }
					<th class="{regionHeadCell}" role="columnheader">{@html heading}</th>
				{/each}
			</tr>
		</thead>
		<!-- Body -->
		<tbody class="table-body {regionBody}">
			{#each source.body as row, rowIndex}
				<!-- Row -->
				<!-- prettier-ignore -->
				<tr
					on:click={(e) => { onRowClick(e, rowIndex); }}
					on:keydown={(e) => { onRowKeydown(e, rowIndex); }}
					aria-rowindex={rowIndex + 1}
				>
					{#each row as cell, cellIndex}
						<!-- Cell -->
						<td
							class="{regionCell}"
							role="gridcell"
							aria-colindex={cellIndex + 1}
							tabindex={cellIndex === 0 && interactive ? 0 : -1}
						>
							{@html Number(cell) === 0 ? cell : (cell ? cell : '-')}
						</td>
					{/each}
				</tr>
			{/each}
		</tbody>
		<!-- Foot -->
		{#if source.foot}
			<tfoot class="table-foot {regionFoot}">
				<tr>
					{#each source.foot as cell }
						<td class="{regionFootCell}">{@html cell}</td>
					{/each}
				</tr>
			</tfoot>
		{/if}
	</table>
</div>
