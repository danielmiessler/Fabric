<script>
  import { languageStore } from '$lib/store/language-store';
</script>

<div class="language-indicator">
  Current Language: {$languageStore}
</div>

<style>
  .language-indicator {
    padding: 0.5rem;
    border-radius: 0.25rem;
    background-color: rgba(0, 0, 0, 0.1);
    display: inline-block;
  }
</style>
