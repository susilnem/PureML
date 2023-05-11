<script lang="ts">
  import { page } from "$app/stores";
  import Breadcrumbs from "$lib/components/Breadcrumbs.svelte";
  import Tabbar from "$lib/components/Tabbar.svelte";
  import Select from "$lib/components/Select.svelte";
  import { ListboxOption } from "@rgossiaux/svelte-headlessui";
  import Avatar from "$lib/components/Avatar.svelte";

  export let data;

  // let ver1 = "";
  // let setVer1 = "";
  // let ver2 = "";
  // let setVer2 = "";
  // // const [submitReviewVersion, setSubmitReviewVersion] = useState("");
  // let branch = "main";
  // let setBranch = "";
  // let dataVersion = {};
  // let setDataVersion = {};
  // // const [ver1Logs, setVer1Logs] = useState<{ [key: string]: string }>({});
  // // const [ver2Logs, setVer2Logs] = useState<{ [key: string]: string }>({});
  // // const [commonMetrics, setCommonMetrics] = useState<string[]>([]);
  let versionData = data.versions;
  // let setVersionData = data.versions;

  // setVer1(versionData.at(0).version);
  // setVer2("");

  function branchChange(event: any) {
    // setBranch(event.target.value);
    // submit(event.currentTarget, { replace: true });
  }
</script>

<div
  class="flex justify-center sticky top-0 bg-slate-50 w-full border-b border-slate-200"
>
  <div class="flex justify-between px-12 2xl:pr-0 w-full max-w-screen-2xl">
    <Breadcrumbs />
    <Tabbar
      tabType="primary"
      tabContent="primaryModel"
      tab="versions"
      fullWidth={false}
    />
  </div>
</div>
<div class="flex justify-center w-full">
  <div
    class="bg-slate-50 flex flex-col h-screen overflow-hidden w-full 2xl:max-w-screen-2xl"
  >
    <div class="flex justify-between h-full">
      <div class="w-full"><slot /></div>

      <!-- ##### versions list right sidebar ##### -->
      <aside
        class="bg-slate-50 border-l-2 border-slate-100 h-full w-1/4 max-w-[400px] py-8 px-12 z-10"
      >
        <form method="post" class="flex justify-end" on:change={branchChange}>
          <input name="_action" value="changeBranch" type="hidden" />
          <!-- <Select
                  intent="primary"
                  name="branch"
                  title={data.params.branchId}
                >
                  {branchData.map((branch: any, index: number) => (
                    <SelectPrimitive.Item
                      key={`${branch}-${index}`}
                      value={branch.value}
                      class="flex items-center justify-between px-4 py-2 rounded-md text-base text-slate-600 font-medium cursor-pointer  hover:bg-slate-100 hover:border-none focus:outline-none"
                    >
                      <SelectPrimitive.ItemText class="text-slate-600 text-base font-medium">
                        {branch.label}
                      </SelectPrimitive.ItemText>
                      <SelectPrimitive.ItemIndicator>
                        <Check class="text-slate-400 w-4" />
                      </SelectPrimitive.ItemIndicator>
                    </SelectPrimitive.Item>
                  ))}
                </Select> -->
          <Select
            intent="primary"
            fullWidth={false}
            name="branch"
            title={$page.params.branchId}
            >{#each data.allBranches as branch}
              <ListboxOption
                value={branch.value}
                class="flex items-center justify-between px-4 py-2 rounded-md text-base text-slate-600 font-medium cursor-pointer hover:bg-slate-100 hover:border-none focus:outline-none"
              >
                {branch.label}
              </ListboxOption>
            {/each}</Select
          >
        </form>
        {#if versionData}
          <ul class="h-3/4 space-y-2 mt-8 overflow-auto">
            {#each versionData as version}
              <li class="flex items-center justify-between py-2">
                <div class="flex items-center">
                  <input
                    class="checkbox checkbox-primary checkbox-sm"
                    name={"version2"}
                    value={version.version}
                    type="checkbox"
                  />
                  <!-- checked={version.version === ver1 ||
                      version.version === ver2} -->
                  <!-- onChange={(e) => {
                      if (e.target.checked) {
                        setVer2(version.version);
                      } else if (ver1 === version.version && ver2 === "") {
                        new Error("You can't uncheck the present version");
                      } else if (ver1 === version.version) {
                        setVer1(ver2);
                        setVer2("");
                      } else {
                        setVer2("");
                      }
                    }} -->
                  <div
                    class="flex items-center justify-center pl-4 text-slate-600"
                  >
                    <Avatar intent="primary">
                      {version.created_by.name.charAt(0).toUpperCase()}
                    </Avatar>
                    <div>
                      <p class="px-4 font-medium">
                        {version.version}
                      </p>
                      <p class="px-4">
                        Created by {version.created_by.name}
                      </p>
                    </div>
                  </div>
                </div>
                <!-- {data.params.branchId !== "main" && (
                        <form
                          method="post"
                          onChange={submitReview}
                          class="flex justify-end"
                        >
                          <input
                            name="_action"
                            value="submitReview"
                            type="hidden"
                          />
                          <input
                            name="fromBranch"
                            value={branch}
                            type="hidden"
                          />
                          <input name="toBranch" value="main" type="hidden" />
                          <Select
                            intent="more"
                            name="version"
                            title={submitReviewVersion}
                          >
                            <SelectPrimitive.Item
                              value={version.version}
                              class="flex items-center justify-between px-4 py-2 rounded-md text-base text-slate-600 font-medium cursor-pointer  hover:bg-slate-100 hover:border-none focus:outline-none"
                            >
                              <SelectPrimitive.ItemText class="text-slate-600 text-base font-medium">
                                Submit For review
                              </SelectPrimitive.ItemText>
                            </SelectPrimitive.Item>
                          </Select>
                        </form>
                      )} -->
              </li>
            {/each}
          </ul>
        {:else}
          <div class="py-8">No version created yet</div>
        {/if}
      </aside>
    </div>
  </div>
</div>
