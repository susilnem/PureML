<script>
  import Avatar from "$lib/components/Avatar.svelte";
  import Breadcrumbs from "$lib/components/Breadcrumbs.svelte";
  import Button from "$lib/components/Button.svelte";
  import Tabbar from "$lib/components/Tabbar.svelte";
  import {
    ChevronDown,
    ChevronUp,
    GitFork,
    GitPullRequest,
  } from "lucide-svelte";

  let params = false;
</script>

<div class="flex justify-center w-full">
  <div
    class="bg-slate-50 flex flex-col h-screen overflow-hidden w-full 2xl:max-w-screen-2xl"
  >
    <div class="flex justify-between h-full">
      <div class="w-4/5">
        <Tabbar
          tabType="secondary"
          tabContent="modelReviewMetricsTab"
          tab="metrics"
        />
        <div class="px-12 pt-8 pb-8 h-[70vh] overflow-auto">
          <!-- {/* ##### parameters section ##### */} -->
          <section class="pt-16">
            <button
              class="flex items-center justify-between w-full border-b-slate-300 border-b pb-4"
              on:click={() => (params = !params)}
            >
              <h1 class="text-slate-800 font-medium text-sm">Parameters</h1>
              {#if params}
                <ChevronUp class="text-slate-400" />
              {:else}
                <ChevronDown class="text-slate-400" />
              {/if}
            </button>
            <div class="py-6">
              <table class=" max-w-[1000px] w-full">
                <thead>
                  <tr>
                    <th class="text-slate-600 font-medium text-left border p-4">
                      {" "}
                    </th>
                    <th
                      class="text-slate-600 font-medium text-left border p-4 w-1/5"
                    >
                      Submitted commit ver1
                    </th>
                    <!-- {ver2 !== "" ? ( -->
                    <th
                      class="text-slate-600 font-medium text-left border p-4 w-1/5"
                    >
                      ver2
                    </th>
                    <!--   ) : null} -->
                  </tr>
                </thead>
                <!-- {paramsKeys.map(
                                    (param: any, index: number) => (
                                      <tr key={index}>
                                        <th class="text-slate-600 font-medium text-left border p-4">
                                          {param}
                                        </th>
                                        <td class="text-slate-600 font-medium text-left border p-4">
                                          {reviewParams[param]
                                            ? reviewParams[param].slice(0, 5)
                                            : "-"}
                                        </td>
                                        {ver2 !== "" && (
                                          <td class="text-slate-600 font-medium text-left border p-4">
                                            {ver2Params[param]
                                              ? ver2Params[param].slice(0, 5)
                                              : "-"}
                                          </td>
                                        )}
                                      </tr>
                                    )
                                  )} -->
              </table>
            </div>
          </section>
        </div>
      </div>
      <!-- ##### versions list right sidebar ##### -->
      <aside
        class="bg-slate-50 border-l-2 border-slate-100 h-full w-1/4 max-w-[400px] py-8 px-12 z-10"
      >
        <div class="flex justify-end">
          <select
            name="branch"
            class="text-slate-500 border-slate-400 bg-transparent rounded"
            disabled
          >
            <option value="value" selected>
              <!-- {toBranch} -->
              Branch
            </option>
          </select>
        </div>
        <!-- {/* incoming branch for review will be displayed here */} -->
        <ul class="h-2/5 space-y-2 mt-8 overflow-auto">
          <li class="flex items-center justify-between py-2">
            <div class="flex items-center">
              <input
                class="checkbox checkbox-primary checkbox-sm"
                name={"version2"}
                type="checkbox"
                value="version.version"
                disabled
              />
              <!-- defaultChecked -->
              <div class="flex items-center justify-center pl-4 text-slate-600">
                <Avatar>S</Avatar>
                <div>
                  <p class="px-4 font-medium">
                    Submitted review
                    <!-- {reviewVersion} -->
                  </p>
                  <p class="px-4">from branch reviewBranch</p>
                </div>
              </div>
            </div>
          </li>
          <li class="flex items-center justify-between py-2">
            <!-- key={version.version} -->
            <div class="flex items-center">
              <input
                class="checkbox checkbox-primary checkbox-sm"
                name={"version2"}
                value="version.version"
                type="checkbox"
              />
              <!-- checked={version.version === ver2}
                  onChange={(e) => {
                    setVer2(version.version);
                  }} -->
              <div class="flex items-center justify-center pl-4 text-slate-600">
                <Avatar intent="primary">
                  <!-- {version.created_by.name.charAt(0).toUpperCase()} -->
                  K
                </Avatar>
                <div>
                  <p class="px-4 font-medium">
                    <!-- {version.version} -->
                    Vx
                  </p>
                  <p class="px-4">
                    Created by
                    <!-- {version.created_by.name} -->
                  </p>
                </div>
              </div>
            </div>
          </li>
        </ul>
        <div class="h-2/5 mt-8">
          <div class="text-base text-slate-800 font-medium pb-4">Details</div>
          <div class="h-4/5 overflow-auto">
            <div class="text-slate-400">
              <div class="flex justify-between items-center py-1">
                <span class="w-1/7 flex items-center">
                  <GitFork class="w-4" />
                </span>
                <span class="w-1/2 pl-2">Commit version</span>
                <span class="w-1/2 pl-2 text-slate-600 font-medium">
                  <!-- {reviewVersion || "Version"} -->
                  Version
                </span>
              </div>
              <div class="flex justify-between items-center py-1">
                <span class="w-1/7 flex items-center">
                  <GitPullRequest class="w-4" />
                </span>
                <span class="w-1/2 pl-2">Commit branch</span>
                <span class="w-1/2 pl-2 text-slate-600 font-medium">
                  <!-- {reviewBranch || "Branch"} -->
                  Branch
                </span>
              </div>
              <!-- <div class="flex justify-between items-center py-1">
                        <span class="w-1/7 flex items-center">
                          <User class="w-4" />
                        </span>
                        <span class="w-1/2 pl-2">Pushed by</span>
                        <span class="w-1/2 pl-2 text-slate-600 font-medium">
                          {`${datasetDetails[0].is_public ? "Yes" : "No"}`}
                        </span>
                      </div> -->
            </div>
            <div class="flex gap-x-4 mt-6">
              <form method="post">
                <input name="_action" value="rejected" type="hidden" />
                <input
                  name="fromBranch"
                  value="data.from_branch"
                  type="hidden"
                />
                <input
                  name="fromBranchVersion"
                  value="data.from_branch_version"
                  type="hidden"
                />
                <input name="isAccepted" value="false" type="hidden" />
                <Button intent="secondary" fullWidth={false}>Reject</Button>
              </form>
              <form method="post">
                <input name="_action" value="accepted" type="hidden" />
                <input
                  name="fromBranch"
                  value="data.from_branch"
                  type="hidden"
                />
                <input
                  name="fromBranchVersion"
                  value="data.from_branch_version"
                  type="hidden"
                />
                <input name="isAccepted" value="true" type="hidden" />
                <Button intent="primary" fullWidth={false}>Approve</Button>
              </form>
            </div>
          </div>
        </div>
      </aside>
    </div>
  </div>
</div>
