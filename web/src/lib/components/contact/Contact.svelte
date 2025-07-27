<script lang="ts">
  import { Contact } from 'lucide-svelte';
</script>

<div class="form-control w-full m-auto p-4 rounded-lg bg-gradient-to-br variant-gradient-success-warning shadow-lg text-current" title="contact form">
  <h2 class="font-bold pl-2">We'd love to hear from you</h2>
  <p class="font-bold pl-2">Email</p>
  <div class="input-group input-group-divider grid-cols-[1fr_auto]">
    <input type="text" placeholder="Enter an email address where you can be reached..." />
  </div>
  <p class="font-bold pl-2">Website</p>
  <div class="input-group input-group-divider grid-cols-[auto_1fr_auto]">
    <div class="input-group-shim">https://</div>
    <input type="text" placeholder="www.example.com" />
  </div>
  <p class="font-bold pl-2">Contact Information</p>
  <div class="input-group input-group-divider grid-cols-[1fr_auto]">
    <input type="text" placeholder="Enter a other contact information here..." />
  </div>
  <label class="label">
    <span class="font-bold pl-2">Message</span>
    <textarea class="textarea" rows="4" placeholder="Enter your message ..." />
  </label>
  <a href="/" title=""><button class="button variant-filled-secondary rounded-lg p-2"><Contact /></button></a>

</div>
