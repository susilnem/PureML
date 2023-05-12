<script>
  import ComparisionTable from "$lib/components/Comparision.svelte";
  import Tabbar from "$lib/components/Tabbar.svelte";
  let version1 = "0.2";
  let version2 = "0.12";
</script>

<Tabbar tabType="secondary" tabContent="datasetTab" tab="datalineage" />
<div class="px-12 pt-2 pb-8 h-[100vh] overflow-auto">
  <h1 class="text-red-600">Adding datalineage soon...</h1>
  <section>
    <div class="flex pt-6 gap-x-8">
      <div class="w-1/2 h-[80vh] max-h-full">
        <div class="text-sm text-slate-600 font-medium pb-6">
          <!-- {version1} -->
          Ver 1
        </div>
      </div>
      <div class="w-1/2 h-[80vh] max-h-full">
        <div class="text-sm text-slate-600 font-medium pb-6">
          <!-- {version2} -->
          Ver 2
        </div>
      </div>
    </div>
  </section>
</div>
