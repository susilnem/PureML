<script>
  import Avatar from "$lib/components/Avatar.svelte";
  import Breadcrumbs from "$lib/components/Breadcrumbs.svelte";
  import Button from "$lib/components/Button.svelte";
  import Tabbar from "$lib/components/Tabbar.svelte";
  import { GitFork, GitPullRequest } from "lucide-svelte";
</script>

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
            <div class="flex pt-6 gap-x-8">
              <div class="w-1/2 h-screen max-h-full">
                <div class="text-sm text-slate-600 font-medium pb-6">
                  Submitted commit
                  <!-- {reviewVersion} -->
                </div>
                <!-- {node && <Pipeline pnode={node} pedge={edge} />} -->
              </div>
              <div class="w-1/2 h-screen max-h-full">
                <div class="text-sm text-slate-600 font-medium pb-6">ver2</div>
                <!-- {node2 && <Pipeline pnode={node2} pedge={edge2} />} -->
              </div>
            </div>
          </section>
        </div>
      </div>
      <!-- {/* ##### versions list right sidebar ##### */} -->
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
                disabled
              />
              <!-- defaultChecked -->
              <div class="flex items-center justify-center pl-4 text-slate-600">
                <Avatar>S</Avatar>
                <div>
                  <p class="px-4 font-medium">Submitted review reviewVersion</p>
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
              <!--    checked={version.version === ver2}
                           onChange={(e) => {
                             setVer2(version.version);
                           }} -->
              <div class="flex items-center justify-center pl-4 text-slate-600">
                <Avatar intent="primary">
                  <!-- {version.created_by.name.charAt(0).toUpperCase()} -->
                  K
                </Avatar>
                <div>
                  <p class="px-4 font-medium">version.version</p>
                  <p class="px-4">
                    Created by
                    <!--  {version.created_by.name} -->
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
              <!-- {/* <div class="flex justify-between items-center py-1">
                      <span class="w-1/7 flex items-center">
                        <User class="w-4" />
                      </span>
                      <span class="w-1/2 pl-2">Pushed by</span>
                      <span class="w-1/2 pl-2 text-slate-600 font-medium">
                        {`${datasetDetails[0].is_public ? "Yes" : "No"}`}
                      </span>
                    </div> */} -->
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
