<script lang="ts">
  import { page } from "$app/stores";

  const url = decodeURI($page.url.pathname.slice(1)).split("/");
  const urlitems = url.filter(function (val, idx) {
    if ((idx + 1) % 2 == 0) return val;
  });
</script>

<ul class="flex pt-6 pr-10 border-r border-slate-200">
  {#each urlitems as item, index}
    <li>
      <a
        href={`/${url.slice(0, index + 3).join("/")}`}
        class={item === urlitems[urlitems.length - 1]
          ? "font-semibold text-slate-950 w-full"
          : "text-slate-400 hover:text-slate-950 w-full"}
      >
        {item}
      </a>
      {#if index !== url.length - 1}
        <span class="text-slate-400 mx-1">/</span>
      {/if}
    </li>
  {/each}
</ul>
