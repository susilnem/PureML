<script lang="ts">
  import "../app.postcss";
  import { page } from "$app/stores";
  import Navbar from "$lib/components/Navbar.svelte";
  import { Toaster } from "svelte-french-toast";

  export let data;
</script>

<Toaster />
<div class="h-screen overflow-hidden">
  {#if $page.url.pathname === "/models" || $page.url.pathname === "/datasets" || $page.url.pathname === `/${$page.params.username}` || $page.url.pathname === `/org/${$page.params.orgId}` || $page.url.pathname.includes("/settings")}
    {#if !data.orgDetails}
      <Navbar intent="loggedOut" user="" orgAvatarName="" />
    {:else}
      <Navbar
        intent="loggedIn"
        user={data.orgDetails[0].handle.charAt(0).toUpperCase()}
        orgAvatarName={data.orgDetails[0].name.charAt(0)}
        ><a
          slot="orgName"
          href={`/org/${data.orgDetails[0].name}`}
          class="flex items-center justify-center"
        >
          {data.orgDetails[0].name}
        </a></Navbar
      >
    {/if}
    <slot />
  {:else}
    <slot />
  {/if}
</div>
