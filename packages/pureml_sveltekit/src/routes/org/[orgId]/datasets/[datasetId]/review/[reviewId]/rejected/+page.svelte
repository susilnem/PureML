<script>
  import Avatar from "$lib/components/Avatar.svelte";
  import Breadcrumbs from "$lib/components/Breadcrumbs.svelte";
  import Tabbar from "$lib/components/Tabbar.svelte";
  import { GitFork, GitPullRequest } from "lucide-svelte";
</script>

<div
  class="flex justify-center sticky top-0 bg-slate-50 w-full border-b border-slate-200"
>
  <div class="flex justify-between px-12 2xl:pr-0 w-full max-w-screen-2xl">
    <Breadcrumbs />
    <Tabbar
      tabType="primary"
      tabContent="primaryDatasetTab"
      tab="review"
      fullWidth={false}
    />
  </div>
</div>
<div class="flex justify-center w-full">
  <div
    class="bg-slate-50 flex flex-col h-screen overflow-hidden w-full 2xl:max-w-screen-2xl"
  >
    <div class="flex justify-between h-full">
      <div class="w-4/5">
        <Tabbar
          tabType="secondary"
          tabContent="datasetReviewLineageTab"
          tab="datalineage"
        />
        <div class="px-12 pt-8 pb-8 h-[100vh] overflow-auto">
          <section>
            <!-- {!versionData && (
                    <div class="text-slate-600 pt-6">
                      No data lineage available
                    </div>
                  )} -->
            <!-- {versionData && ( -->
            <div class="flex pt-6 gap-x-8">
              <div class="w-full h-screen max-h-full">
                <!-- {node && <Pipeline pnode={node} pedge={edge} />} -->
                Pipe
              </div>
            </div>
            <!--   )} -->
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
            <option value="value" selected> dev </option>
          </select>
        </div>
        <!-- incoming branch for review will be displayed here -->
        <!-- {versionData ? ( -->
        <ul class="space-y-2 mt-8 overflow-auto">
          <li class="flex items-center justify-between py-2">
            <div class="flex items-center">
              <input
                class="checkbox checkbox-primary checkbox-sm"
                name={"version2"}
                type="checkbox"
                disabled
              />
              <!-- defaultChecked -->
              <div class="flex items-center justify-center pl-4 text-slate-600">
                <Avatar>R</Avatar>
                <div>
                  <p class="px-4 font-medium">Rejected review reviewVersion</p>
                  <p class="px-4">from branch reviewBranch</p>
                </div>
              </div>
            </div>
          </li>
        </ul>
        <!--   ) : (
                <div class="py-8 mt-8">No version created yet</div>
              )} -->
        <div class="h-3/5 mt-8">
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
                        Name
                      </span>
                    </div> -->
            </div>
          </div>
        </div>
      </aside>
    </div>
  </div>
</div>
