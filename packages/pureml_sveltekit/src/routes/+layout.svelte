<script>
  import { page } from "$app/stores";
  import Navbar from "$lib/components/Navbar.svelte";
  import "../app.postcss";
</script>

<div class="h-screen overflow-hidden">
  {#if $page.url.pathname.includes("/auth") || $page.url.pathname === "/reset-password"}
    <slot />
  {:else}
    <Navbar intent="loggedIn" user="K" orgAvatarName="K" orgName="Private" />
    <slot />
  {/if}
</div>
