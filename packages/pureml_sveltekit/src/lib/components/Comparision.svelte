<script lang="ts">
  import { ChevronDown, ChevronUp } from "lucide-svelte";

  let open = false;
  let commonMetrics = [];
  export let version1 = "";
  export let version2 = "";
</script>

<section>
  <button
    class="flex items-center justify-between w-full border-b-slate-300 border-b pb-4"
    on:click={() => {
      open = !open;
    }}
  >
    <h1 class="text-slate-800 font-medium text-sm">
      <!-- {metric} -->
      Metric
    </h1>
    {#if open}
      <ChevronUp class="text-slate-400" />
    {:else}
      <ChevronDown class="text-slate-400" />
    {/if}
  </button>

  {#if open}
    <div class="py-6">
      {#if commonMetrics.length !== 0}
        <table class="max-w-[1000px] w-full">
          {#if commonMetrics.length !== 0}
            <thead>
              <tr>
                <th class="text-slate-600 font-medium text-left border p-4">
                  {" "}
                </th>
                <th
                  class="text-slate-600 font-medium text-left border p-4 w-1/5"
                >
                  {version1}
                </th>
                {#if version2 !== ""}
                  <th
                    class="text-slate-600 font-medium text-left border p-4 w-1/5"
                  >
                    {version2}
                  </th>
                {:else}
                  {null}
                {/if}
              </tr>
            </thead>
            <!-- {commonMetrics.map((metric, i) => (
                      <tr key={i}>
                        <th class="text-slate-600 font-medium text-left border p-4">
                          {metric}
                        </th>
                        <td class="text-slate-600 font-medium text-left border p-4 w-1/5 truncate">
                          {Object.keys(dataVer1).length > 0
                            ? dataVer1[metric]
                              ? dataVer1[metric].slice(0, 5)
                              : "-"
                            : "No-data"}
                        </td>
                        {version2 !== "" && (
                          <td class="text-slate-600 font-medium text-left border p-4 w-1/5 truncate">
                            {dataVer2
                              ? dataVer2[metric]
                                ? dataVer2[metric].slice(0, 5)
                                : "-"
                              : "No-data"}
                          </td>
                        )}
                      </tr>
                    ))} -->
          {/if}
        </table>
      {:else}
        <div>No metric available</div>
      {/if}
    </div>
  {/if}
</section>
