<script>
  import "../../app.postcss";
</script>

<svelte:head>
  <title>Join Us | PureML</title>
  <meta name="description" content="Start your journey with PureML." />
</svelte:head>

<div class="w-screen h-screen flex items-center justify-center">
  <div
    class="md:flex md:flex-col justify-center items-center md:pt-0 text-white"
  >
    <slot />
  </div>
</div>
