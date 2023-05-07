<script>
  import Navbar from "$lib/components/Navbar.svelte";
  import "../app.postcss";
</script>

<div class="h-screen overflow-hidden">
  <Navbar intent="loggedIn" user="K" orgAvatarName="K" orgName="Private" />
  <slot />
</div>
