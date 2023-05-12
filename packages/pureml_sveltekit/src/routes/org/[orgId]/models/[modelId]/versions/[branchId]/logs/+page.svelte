<script>
  import ComparisionTable from "$lib/components/Comparision.svelte";
  import Tabbar from "$lib/components/Tabbar.svelte";
  let version1 = "0.2";
  let version2 = "0.12";

  export let data;
</script>

<Tabbar tabType="secondary" tabContent="modelTab" tab="userlogs" />
<!-- <VersionContext.Provider value={{ version1, version2 }}> -->
<div class="px-12 py-6">
  <h1 class="text-red-600">Adding metrics soon...</h1>
  <ComparisionTable {version1} {version2} />
  <!-- dataVer1={version1Logs[key]}
            dataVer2={version2Logs[key]} -->
</div>
<!-- </VersionContext.Provider> -->
