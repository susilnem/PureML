<script lang="ts">
  import { page } from "$app/stores";
  import Breadcrumbs from "$lib/components/Breadcrumbs.svelte";
  import Tabbar from "$lib/components/Tabbar.svelte";
  import Select from "$lib/components/Select.svelte";
  import { ListboxOption } from "@rgossiaux/svelte-headlessui";
  import Avatar from "$lib/components/Avatar.svelte";
  import { version1, type Version, version2 } from "./stores";
  import { onMount } from "svelte";

  export let data;

  function isJson(item: string | object) {
    let value = typeof item !== "string" ? JSON.stringify(item) : item;
    try {
      value = JSON.parse(value);
    } catch (e) {
      return false;
    }

    return typeof value === "object" && value !== null;
  }

  let versionData: Version[] = data.versions;

  onMount(() => {
    if (!versionData) return;
    if (!versionData[0]) return;

    version1.set(versionData[0].version);
    version2.set("");
  });

  function onVersionChangeHandler(event: Event, version: Version) {
    if ((event?.target as HTMLInputElement)?.checked) {
      version2.set(version.version);
    } else if ($version1 === version.version && $version2 === "") {
      new Error("You can't uncheck the present version");
    } else if ($version1 === version.version) {
      version1.set($version2);
      version2.set("");
    } else {
      version2.set("");
    }
  }
</script>

<div
  class="flex justify-center sticky top-0 bg-slate-50 w-full border-b border-slate-200"
>
  <div class="flex justify-between px-12 2xl:pr-0 w-full max-w-screen-2xl">
    <Breadcrumbs />
    <Tabbar
      tabType="primary"
      tabContent="primaryDataset"
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
              <a
                data-sveltekit-reload
                href={`/org/${$page.params.orgId}/datasets/${$page.params.datasetId}/versions/${branch.label}/datalineage`}
                class="w-full"
              >
                {branch.label}</a
              >
            </ListboxOption>
          {/each}</Select
        >
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
                    checked={version.version === $version1 ||
                      version.version === $version2}
                    on:change={(e) => onVersionChangeHandler(e, version)}
                  />
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
