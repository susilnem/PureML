<script>
  import { goto } from "$app/navigation";
  import { onMount } from "svelte";

  export let data;
  const username = data?.userDetails[0].handle;
  onMount(() => {
    goto(`/${username}`);
  });
</script>
