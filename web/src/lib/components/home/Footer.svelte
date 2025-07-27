<script>
  const year = new Date().getFullYear();
  import BuyMeCoffee from "$lib/components/ui/buymeacoffee/BuyMeCoffee.svelte";
</script>

<footer class="border-t bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60">
  <div class="container flex h-14 items-center justify-between px-4">
    <p class="text-sm text-muted-foreground">
      Built in {year} by @Fabric <!-- Feel free to put your name here -->
    </p>
    
    <nav class="flex items-center gap-4 ">
      <BuyMeCoffee url="https://www.buymeacoffee.com/" /> <!-- And here -->
    </nav>
  </div>
</footer>