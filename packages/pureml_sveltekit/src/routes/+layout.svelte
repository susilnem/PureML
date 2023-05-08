<script>
  import { page } from "$app/stores";
  import Navbar from "$lib/components/Navbar.svelte";
  import "../app.postcss";
  import { Toaster } from "svelte-french-toast";
</script>

<Toaster />
<div class="h-screen overflow-hidden">
  {#if $page.url.pathname.includes("/auth") || $page.url.pathname === "/reset-password"}
    <slot />
  {:else}
    <Navbar intent="loggedIn" user="U" orgAvatarName="U" orgName="PureML" />
    <slot />
  {/if}
</div>
