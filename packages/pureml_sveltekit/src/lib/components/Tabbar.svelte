<script lang="ts">
  import clsx from "clsx";
  import { tv } from "tailwind-variants";

  const tabStyles = tv({
    base: "pt-6 text-zinc-400 flex sticky z-10 pl-12",
    variants: {
      intent: {
        primaryModelTab: "bg-slate-50",
        primaryDatasetTab: "bg-slate-50",
        primarySettingTab: "top-[3.3rem]",
        modelTab: "pb-6",
        datasetTab: "pb-6",
      },
      fullWidth: {
        true: "w-full",
        false: "w-fit",
      },
    },
    defaultVariants: {
      intent: "primaryModelTab",
      fullWidth: true,
    },
  });

  function primaryLinkCss(currentPage: boolean) {
    return clsx(
      currentPage ? "text-slate-600 font-medium" : "text-slate-500",
      "flex justify-center items-center hover:text-slate-600"
    );
  }

  export let tabPage: string;
  export let tab: string;
  let tabItems: any = [];

  tabPage = $$props.tabPage;

  if (tabPage === "primaryModel") {
    tabItems = [
      {
        id: "modelCard",
        name: "Model Card",
        hyperlink: `/org`,
      },
      {
        id: "versions",
        name: "Versions",
        hyperlink: `/org/versions/main/logs`,
      },
      {
        id: "review",
        name: "Review",
        hyperlink: `/org/review`,
      },
    ];
  }
  if (tabPage === "primaryDataset") {
    tabItems = [
      {
        id: "datasetCard",
        name: "Dataset Card",
        hyperlink: `/org`,
      },
      {
        id: "versions",
        name: "Versions",
        hyperlink: `/org/versions/main/datalineage`,
      },
      {
        id: "review",
        name: "Review",
        hyperlink: `/org/review`,
      },
    ];
  }
  if (tabPage === "primarySettings") {
    tabItems = [
      {
        id: "profile",
        name: "Profile",
        hyperlink: `/settings`,
      },
      {
        id: "organization",
        name: "Organization",
        hyperlink: `/settings/organization`,
      },
      {
        id: "members",
        name: "Members",
        hyperlink: `/settings/members`,
      },
      {
        id: "apiToken",
        name: "API Tokens",
        hyperlink: `/settings/apitokens`,
      },
    ];
  }
  if (tabPage === "modelTabs") {
    tabItems = [
      {
        id: "userlogs",
        name: "User Logs",
        hyperlink: `/org/versions/main/logs`,
      },
      // {
      //   id: "graphs",
      //   name: "Graphs",
      //   hyperlink: `/org/versions/logs`,
      // },
    ];
  }
  if (tabPage === "datasetTabs") {
    tabItems = [
      {
        id: "datalineage",
        name: "Data Lineage",
        hyperlink: `/org/versions/main/datalineage`,
      },
      // {
      //   id: "code",
      //   name: "Code",
      //   hyperlink: `/org/versions/code`,
      // },
    ];
  }
</script>

<div class={tabStyles({ ...$$restProps })}>
  <div class="flex">
    {#each tabItems as item}
      <div class="pl-6">
        {#if tab === item.id}
          <div
            class="pb-4 hover:text-slate-850 text-brand-200 border-b-2 border-brand-200 font-medium"
          >
            <a href={item.hyperlink}>
              <span>{item.name}</span>
            </a>
          </div>
        {:else}
          <div class="pb-4 hover:text-slate-850 text-slate-500">
            <a href={item.hyperlink}>
              <span>{item.name}</span>
            </a>
          </div>
        {/if}
      </div>
    {/each}
  </div>
</div>
