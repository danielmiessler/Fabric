<script>export let files = void 0;
export let fileInput = void 0;
export let name;
export let border = "border-2 border-dashed";
export let padding = "p-4 py-8";
export let rounded = "rounded-container-token";
export let regionInterface = "";
export let regionInterfaceText = "";
export let slotLead = "mb-4";
export let slotMessage = "";
export let slotMeta = "opacity-75";
const cBase = "textarea relative flex justify-center items-center";
const cInput = "w-full absolute top-0 left-0 right-0 bottom-0 z-[1] opacity-0 disabled:!opacity-0 cursor-pointer";
const cInterface = "flex justify-center items-center text-center";
$:
  classesBase = `${cBase} ${border} ${padding} ${rounded} ${$$props.class ?? ""}`;
$:
  classesInput = `${cInput}`;
$:
  classesInterface = `${cInterface}`;
function prunedRestProps() {
  delete $$restProps.class;
  return $$restProps;
}
</script>

<div class="dropzone {classesBase}" class:opacity-50={$$restProps.disabled} data-testid="file-dropzone">
	<!-- Input: File (hidden) -->
	<!-- NOTE: keep `bind:files` here, unlike FileButton -->
	<input
		bind:files
		bind:this={fileInput}
		type="file"
		{name}
		class="dropzone-input {classesInput}"
		{...prunedRestProps()}
		on:change
		on:dragenter
		on:dragover
		on:dragleave
		on:drop
		on:click
		on:keydown
		on:keyup
		on:keypress
		on:focus
		on:focusin
		on:focusout
	/>
	<!-- Interface -->
	<div class="dropzone-interface {classesInterface} {regionInterface}">
		<div class="dropzone-interface-text {regionInterfaceText}">
			<!-- Lead -->
			{#if $$slots.lead}<div class="dropzone-lead {slotLead}"><slot name="lead" /></div>{/if}
			<!-- Message -->
			<div class="dropzone-message {slotMessage}">
				<slot name="message"><strong>Upload a file</strong> or drag and drop</slot>
			</div>
			<!-- Meta Text -->
			{#if $$slots.meta}<small class="dropzone-meta {slotMeta}"><slot name="meta" /></small>{/if}
		</div>
	</div>
</div>
