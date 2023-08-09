<script lang="ts">
  import logo from "$lib/logos/PureMLLogo.svg";
</script>

<svelte:head>
  <title>PureML</title>
  <meta name="description" content="Developer platform for Production ML." />
</svelte:head>

<div class="flex flex-col gap-y-16 justify-center items-center h-screen">
  <img src={logo} alt="Logo" class="w-24" />
  <div class="flex flex-col gap-y-2">
    <h1 class="font-medium text-brand-200 text-4xl">Welcome to PureML UI</h1>
    <h1 class="text-2xl">Developer platform for Production ML.</h1>
  </div>
</div>
