<!-- <ComparisionTable {version1} {version2} /> -->
<!-- dataVer1={version1Logs[key]}
        dataVer2={version2Logs[key]} -->
<script lang="ts">
  import { page } from "$app/stores";
  import Breadcrumbs from "$lib/components/Breadcrumbs.svelte";
  import Tabbar from "$lib/components/Tabbar.svelte";
  import Select from "$lib/components/Select.svelte";
  import { ListboxOption } from "@rgossiaux/svelte-headlessui";
  import Avatar from "$lib/components/Avatar.svelte";
  import comparisonIcon from "$lib/icons/comparison.svg";
  import { onMount } from "svelte";
  import {
    version1,
    version2,
    submitReviewVersion,
    branch,
    dataVersion,
    version1Logs,
    version2Logs,
    commonMetrics,
  } from "../stores";
  import type { Version } from "../stores";
  import VerticalLayout from "$lib/components/VerticalLayout.svelte";
  import Button from "$lib/components/Button.svelte";
  import { Info } from "lucide-svelte";
  import { enhance } from "$app/forms";
  import type { ActionData } from "./$types";

  export let data;
  export let form: ActionData;
  let compare = false;

  let latestBranch;
  let latestVersion: string;
  let _branch: string;
  let _version: string;
  let versionData: any;

  versionData = data.versions;

  if (versionData) {
    latestBranch = versionData[0].branch.name;
    latestVersion = versionData[0].version;
    _branch = latestBranch;
    _version = latestVersion;
  } else {
    _version = "-";
  }

  function isJson(item: string | object) {
    let value = typeof item !== "string" ? JSON.stringify(item) : item;
    try {
      value = JSON.parse(value);
    } catch (e) {
      return false;
    }

    return typeof value === "object" && value !== null;
  }

  // ##### checking version data #####
  // onMount(() => {
  //   if (!versionData) return;
  //   if (!versionData[0]) return;

  //   version1.set(versionData[0].version);
  //   version2.set("");
  // });

  // ##### fetching & displaying latest version data #####
  // onMount(() => {
  //   if (!versionData) return;

  //   const tempDict: Record<string, Version> = {};
  //   versionData.forEach((version) => {
  //     tempDict[version.version] = version;
  //   });
  //   dataVersion.set(tempDict);
  // });

  // dataVersion.subscribe((_dataVersion: any) => {
  //   console.log("subs");
  //   const tempVersion1Data = _dataVersion[$version1];
  //   if (tempVersion1Data) {
  //     if (tempVersion1Data.logs === null) {
  //       version1Logs.set({});
  //       commonMetrics.set([]);
  //       return;
  //     } else {
  //       const tempDictv1: Record<string, string> = {};
  //       tempVersion1Data.logs.forEach((log: any) => {
  //         if (isJson(log.data)) {
  //           tempDictv1[log.key] = JSON.parse(log.data);
  //           if (!$commonMetrics.includes(log.key)) {
  //             commonMetrics.set([...$commonMetrics, log.key]);
  //           }
  //         }
  //       });
  //       version1Logs.set(tempDictv1);
  //     }
  //   }

  //   if ($version2 === "") {
  //     version2Logs.set({});
  //     return;
  //   }
  //   const tempVersion2Data = _dataVersion[$version2];
  //   if (tempVersion2Data) {
  //     if (tempVersion2Data.logs === null) {
  //       version2Logs.set({});
  //       return;
  //     } else {
  //       const tempDictv2: Record<string, string> = {};
  //       tempVersion2Data.logs.forEach((log: any) => {
  //         tempDictv2[log.key] = JSON.parse(log.data);
  //         if (!$commonMetrics.includes(log.key)) {
  //           commonMetrics.set([...$commonMetrics, log.key]);
  //         }
  //       });
  //       version2Logs.set(tempDictv2);
  //     }
  //   }
  // });
</script>

<VerticalLayout>
  <div slot="tabbar" class="flex justify-between w-full">
    <Breadcrumbs />
    <Tabbar
      tabType="primary"
      tabContent="primaryModel"
      tab="analysis"
      fullWidth={false}
    />
  </div>
  <div slot="sidebar" class="flex flex-col gap-y-6">
    <button
      on:click={() => (compare = !compare)}
      class="flex items-center gap-x-2 text-slate-950 border border-slate-300 p-2.5 rounded w-fit cursor-pointer hover:bg-slate-100"
    >
      <img src={comparisonIcon} alt="CompareIcon" class="w-4" /> Compare
    </button>
    <div class="flex gap-x-4">
      <Select intent="primary" fullWidth={true} name="branch" title={_branch}>
        <form method="POST" action="?/changeBranch" use:enhance class="w-full">
          <input type="hidden" name="branchSelected" value={_branch} />
          {#each data.allBranches as branch}
            <button class="w-full"
              ><ListboxOption
                on:click={() => (_branch = branch.label)}
                value={branch.value}
                class="flex items-center justify-between px-4 py-2 rounded-md text-base text-slate-600 font-medium cursor-pointer hover:bg-slate-100 hover:border-none focus:outline-none"
              >
                {branch.label}
              </ListboxOption>
            </button>
          {/each}
        </form>
      </Select>
      <Select
        intent="primary"
        fullWidth={false}
        name="version"
        title={form?.branchVersions[0].version.version || _version}
      >
        {#each form?.branchVersions ?? data.versions as version}
          <ListboxOption
            on:click={() => (_version = version.version)}
            value={version.version}
            class="flex items-center justify-between px-4 py-2 rounded-md text-base text-slate-600 font-medium cursor-pointer hover:bg-slate-100 hover:border-none focus:outline-none"
          >
            <button>{version.version}</button>
          </ListboxOption>
        {/each}
      </Select>
    </div>
    <!-- <div class="flex flex-col gap-y-4">
      <div>
        <div class="text-sm text-slate-600">Dataset</div>
        <Select intent="primary" fullWidth={true} name="branch" title={_branch}>
          {#each data.allBranches as branch}
            <ListboxOption
              value={branch.value}
              class="flex items-center justify-between px-4 py-2 rounded-md text-base text-slate-600 font-medium cursor-pointer hover:bg-slate-100 hover:border-none focus:outline-none"
            >
              <a
                data-sveltekit-reload
                href={`/org/${$page.params.orgId}/models/${$page.params.modelId}/versions/${branch.label}/performance`}
                class="w-full"
              >
                {branch.label}
              </a>
            </ListboxOption>
          {/each}</Select
        >
      </div>
      <div class="flex gap-x-4">
        <Select intent="primary" fullWidth={true} name="branch" title={_branch}>
          {#each data.allBranches as branch}
            <ListboxOption
              value={branch.value}
              class="flex items-center justify-between px-4 py-2 rounded-md text-base text-slate-600 font-medium cursor-pointer hover:bg-slate-100 hover:border-none focus:outline-none"
            >
              <a
                data-sveltekit-reload
                href={`/org/${$page.params.orgId}/models/${$page.params.modelId}/versions/${branch.label}/performance`}
                class="w-full"
              >
                {branch.label}
              </a>
            </ListboxOption>
          {/each}</Select
        >
        <Select
          intent="primary"
          fullWidth={false}
          name="version"
          title="version"
        >
          {#each versionData as version}
            <ListboxOption
              value={version.version}
              class="flex items-center justify-between px-4 py-2 rounded-md text-base text-slate-600 font-medium cursor-pointer hover:bg-slate-100 hover:border-none focus:outline-none"
            >
              <a
                data-sveltekit-reload
                href={`/org/${$page.params.orgId}/models/${$page.params.modelId}/versions/performance`}
                class="w-full"
              >
                {version.version}
              </a>
            </ListboxOption>
          {/each}
        </Select>
      </div>
    </div> -->
    {#if compare}
      <div class="flex flex-col gap-y-4 mt-6">
        <div>
          <div class="text-sm text-slate-600">Model</div>
          <Select
            intent="primary"
            fullWidth={true}
            name="branch"
            title={_branch}
          >
            {#each data.allBranches as branch}
              <ListboxOption
                value={branch.value}
                class="flex items-center justify-between px-4 py-2 rounded-md text-base text-slate-600 font-medium cursor-pointer hover:bg-slate-100 hover:border-none focus:outline-none"
              >
                <a
                  data-sveltekit-reload
                  href={`/org/${$page.params.orgId}/models/${$page.params.modelId}/versions/${branch.label}/performance`}
                  class="w-full"
                >
                  {branch.label}
                </a>
              </ListboxOption>
            {/each}</Select
          >
        </div>
        <div class="flex gap-x-4">
          <Select
            intent="primary"
            fullWidth={true}
            name="branch"
            title={_branch}
          >
            {#each data.allBranches as branch}
              <ListboxOption
                value={branch.value}
                class="flex items-center justify-between px-4 py-2 rounded-md text-base text-slate-600 font-medium cursor-pointer hover:bg-slate-100 hover:border-none focus:outline-none"
              >
                <a
                  data-sveltekit-reload
                  href={`/org/${$page.params.orgId}/models/${$page.params.modelId}/versions/${branch.label}/performance`}
                  class="w-full"
                >
                  {branch.label}
                </a>
              </ListboxOption>
            {/each}</Select
          >
          <Select
            intent="primary"
            fullWidth={false}
            name="version"
            title="version"
          >
            {#each versionData as version}
              <ListboxOption
                value={version.version}
                class="flex items-center justify-between px-4 py-2 rounded-md text-base text-slate-600 font-medium cursor-pointer hover:bg-slate-100 hover:border-none focus:outline-none"
              >
                <a
                  data-sveltekit-reload
                  href={`/org/${$page.params.orgId}/models/${$page.params.modelId}/versions/performance`}
                  class="w-full"
                >
                  {version.version}
                </a>
              </ListboxOption>
            {/each}
          </Select>
        </div>
      </div>
      <div class="flex flex-col gap-y-4">
        <div>
          <div class="text-sm text-slate-600">Dataset</div>
          <Select
            intent="primary"
            fullWidth={true}
            name="branch"
            title={_branch}
          >
            {#each data.allBranches as branch}
              <ListboxOption
                value={branch.value}
                class="flex items-center justify-between px-4 py-2 rounded-md text-base text-slate-600 font-medium cursor-pointer hover:bg-slate-100 hover:border-none focus:outline-none"
              >
                <a
                  data-sveltekit-reload
                  href={`/org/${$page.params.orgId}/models/${$page.params.modelId}/versions/${branch.label}/performance`}
                  class="w-full"
                >
                  {branch.label}
                </a>
              </ListboxOption>
            {/each}</Select
          >
        </div>
        <div class="flex gap-x-4">
          <Select
            intent="primary"
            fullWidth={true}
            name="branch"
            title={_branch}
          >
            {#each data.allBranches as branch}
              <ListboxOption
                value={branch.value}
                class="flex items-center justify-between px-4 py-2 rounded-md text-base text-slate-600 font-medium cursor-pointer hover:bg-slate-100 hover:border-none focus:outline-none"
              >
                <a
                  data-sveltekit-reload
                  href={`/org/${$page.params.orgId}/models/${$page.params.modelId}/versions/${branch.label}/performance`}
                  class="w-full"
                >
                  {branch.label}
                </a>
              </ListboxOption>
            {/each}</Select
          >
          <Select
            intent="primary"
            fullWidth={false}
            name="version"
            title="version"
          >
            {#each versionData as version}
              <ListboxOption
                value={version.version}
                class="flex items-center justify-between px-4 py-2 rounded-md text-base text-slate-600 font-medium cursor-pointer hover:bg-slate-100 hover:border-none focus:outline-none"
              >
                <a
                  data-sveltekit-reload
                  href={`/org/${$page.params.orgId}/models/${$page.params.modelId}/versions/performance`}
                  class="w-full"
                >
                  {version.version}
                </a>
              </ListboxOption>
            {/each}
          </Select>
        </div>
      </div>
    {/if}
    <!-- {#if versionData}
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
              <div class="flex items-center justify-center pl-4 text-slate-600">
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
            {#if data.params.branchId !== "main"}
                  <form
                    method="post"
                    onChange={submitReview}
                    class="flex justify-end"
                  >
                    <input name="_action" value="submitReview" type="hidden" />
                    <input name="fromBranch" value={branch} type="hidden" />
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
                        <SelectPrimitive.ItemText
                          class="text-slate-600 text-base font-medium"
                        >
                          Submit For review
                        </SelectPrimitive.ItemText>
                      </SelectPrimitive.Item>
                    </Select>
                  </form>
                {/if}}
          </li>
        {/each}
      </ul>
    {:else}
      <div class="py-8">No version created yet</div>
    {/if} -->
    <div class="flex flex-col gap-y-4 py-6 border-t border-slate-200">
      <div class="flex gap-x-2 bg-slate-100 rounded-lg p-3 text-slate-600">
        <Info class="w-8 text-blue-700" /> Selected Model version X will be submitted
        for review.
      </div>
      <Button intent="primary">Submit for Review</Button>
    </div>
  </div>
  <div slot="content" class="w-full">
    {form?.branchVersions[0].version.version || _version}
    {_branch}
    {JSON.stringify(
      (form?.branchVersions || data.versions).filter(
        (ver) =>
          ver.version === (form?.branchVersions[0].version.version || _version)
      )
    )}
  </div>
</VerticalLayout>
