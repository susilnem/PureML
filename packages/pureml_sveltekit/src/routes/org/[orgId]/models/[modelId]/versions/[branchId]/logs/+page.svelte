<script>
  import ComparisionTable from "$lib/components/Comparision.svelte";
  import Tabbar from "$lib/components/Tabbar.svelte";
  let ver1 = "0.2";
  let ver2 = "0.12";
</script>

<Tabbar tabType="secondary" tabContent="modelTab" tab="userlogs" />
<!-- <VersionContext.Provider value={{ ver1, ver2 }}> -->
<div class="px-12 py-6">
  <ComparisionTable {ver1} {ver2} />
  <!-- dataVer1={ver1Logs[key]}
            dataVer2={ver2Logs[key]} -->
</div>
<!-- </VersionContext.Provider> -->
