<script>
  import { Avatar } from '@skeletonlabs/skeleton';

  export let header = '';
  export let imageUrl = '';
  export let imageAlt = 'Post';
  export let title = '';
  export let content = '';
  export let authorName = '';
  export let authorAvatarUrl = '';
  export let link = '';
</script>

<div class="w-full text-token grid grid-cols-1 md:grid-cols-1 justify-end gap-4">
  <a class="card card-hover overflow-hidden" href={link}>
    <header>
      <img src={imageUrl} class="bg-black/50 w-full" alt={imageAlt} />
    </header>
    <div class="p-4 space-y-4">
      <h6 class="h6" data-toc-ignore>{header}</h6>
      <h3 class="h3" data-toc-ignore>{title}</h3>
      <article>
        <p>
          {content}
        </p>
      </article>
    </div>
    <hr class="opacity-50" />
    <footer class="p-4 flex justify-start items-center space-x-4">
      <Avatar src={authorAvatarUrl} width="w-8" />
      <div class="flex-auto flex justify-between items-center">
        <h6 class="font-bold" data-toc-ignore>By {authorName}</h6>
        <small>On {new Date().toLocaleDateString()}</small>
      </div>
    </footer>
  </a>
</div>
