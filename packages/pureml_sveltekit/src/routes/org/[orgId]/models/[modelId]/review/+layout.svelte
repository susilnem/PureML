<script>
  import Avatar from "$lib/components/Avatar.svelte";
  import Breadcrumbs from "$lib/components/Breadcrumbs.svelte";
  import Tabbar from "$lib/components/Tabbar.svelte";
</script>

<div id="modelsReview">
  <div
    class="flex justify-center sticky top-0 bg-slate-50 w-full border-b border-slate-200"
  >
    <div class="flex px-12 2xl:pr-0 w-full max-w-screen-2xl">
      <Breadcrumbs />
      <Tabbar
        tabType="primary"
        tabContent="primaryModel"
        tab="review"
        fullWidth={false}
      />
    </div>
  </div>
  <div class="flex justify-center w-full">
    <div
      class="bg-slate-50 flex flex-col h-screen overflow-hidden w-full 2xl:max-w-screen-2xl"
    >
      <slot />
    </div>
  </div>
</div>
