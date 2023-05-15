<script lang="ts">
  import { page } from "$app/stores";
  import clsx from "clsx";
  import { tv } from "tailwind-variants";

  const tabStyles = tv({
    base: "sticky text-zinc-400 pt-6",
    variants: {
      intent: {
        primary: "flex z-10",
        secondary: "flex z-10 pl-6",
        tertiary: "font-medium flex bg-slate-50 px-6",
      },
      content: {
        primaryModelTab: "bg-slate-50",
        primaryDatasetTab: "bg-slate-50",
        primarySettingTab: "top-[3.3rem]",
        modelTab: "pb-6",
        datasetTab: "pb-6",
        modelReviewTab: "pt-8",
        datasetReviewTab: "pt-8",
        modelReviewMetricsTab: "pt-8",
        datasetReviewLineageTab: "pt-8",
      },
      fullWidth: {
        true: "w-full",
        false: "w-fit",
      },
    },
    defaultVariants: {
      intent: "primary",
      fullWidth: true,
    },
  });

  const orgId = $page.params.orgId;
  const modelId = $page.params.modelId;
  const datasetId = $page.params.datasetId;
  const reviewId = $page.params.reviewId;

  export let tabType: string;
  export let tabContent: string;
  export let tab: string;
  let tabItems: any = [];

  $$restProps.intent = $$props.tabType;
  tabContent = $$props.tabContent;

  if (tabContent === "primaryModel") {
    tabItems = [
      {
        id: "modelCard",
        name: "Model Card",
        hyperlink: `/org/${orgId}/models/${modelId}`,
      },
      {
        id: "analysis",
        name: "Analysis",
        hyperlink: `/org/${orgId}/models/${modelId}/versions/main/performance`,
      },
      {
        id: "review",
        name: "Review",
        hyperlink: `/org/${orgId}/models/${modelId}/review`,
      },
    ];
  }
  if (tabContent === "primaryDataset") {
    tabItems = [
      {
        id: "datasetCard",
        name: "Dataset Card",
        hyperlink: `/org/${orgId}/datasets/${datasetId}`,
      },
      {
        id: "analysis",
        name: "Analysis",
        hyperlink: `/org/${orgId}/datasets/${datasetId}/versions/main/datalineage`,
      },
      {
        id: "review",
        name: "Review",
        hyperlink: `/org/${orgId}/datasets/${datasetId}/review`,
      },
    ];
  }
  if (tabContent === "primarySettings") {
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
  if (tabContent === "modelTab") {
    tabItems = [
      {
        id: "performance",
        name: "Performance",
        hyperlink: `/org/${orgId}/models/${modelId}/versions/main/performance`,
      },
      // {
      //   id: "integrity",
      //   name: "Integrity",
      //   hyperlink: `/org/versions/integrity`,
      // },
    ];
  }
  if (tabContent === "datasetTab") {
    tabItems = [
      {
        id: "datalineage",
        name: "Data Lineage",
        hyperlink: `/org/${orgId}/datasets/${datasetId}/versions/main/datalineage`,
      },
      // {
      //   id: "code",
      //   name: "Code",
      //   hyperlink: `/org/versions/code`,
      // },
    ];
  }
  if (tabContent === "modelReviewTab") {
    tabItems = [
      {
        id: "newcommits",
        name: "New Commits",
        hyperlink: `/org/${orgId}/models/${modelId}/review`,
      },
      {
        id: "approved",
        name: "Approved",
        hyperlink: `/org/${orgId}/models/${modelId}/review/approved`,
      },
      {
        id: "rejected",
        name: "Rejected",
        hyperlink: `/org/${orgId}/models/${modelId}/review/rejected`,
      },
    ];
  }
  if (tabContent === "datasetReviewTab") {
    tabItems = [
      {
        id: "newcommits",
        name: "New Commits",
        hyperlink: `/org/${orgId}/datasets/${datasetId}/review`,
      },
      {
        id: "approved",
        name: "Approved",
        hyperlink: `/org/${orgId}/datasets/${datasetId}/review/approved`,
      },
      {
        id: "rejected",
        name: "Rejected",
        hyperlink: `/org/${orgId}/datasets/${datasetId}/review/rejected`,
      },
    ];
  }
  if (tabContent === "modelReviewMetricsTab") {
    tabItems = [
      {
        id: "metrics",
        name: "Metrics",
        hyperlink: `/org/${orgId}/models/${modelId}/review/${reviewId}/logs`,
      },
      // {
      //   id: "graphs",
      //   name: "Graphs",
      //   hyperlink: `/org/models/review/${reviewId}/graphs`,
      // },
    ];
  }
  if (tabContent === "datasetReviewLineageTabs") {
    tabItems = [
      {
        id: "datalineage",
        name: "Data Lineage",
        hyperlink: `/org/${orgId}/datasets/${datasetId}/review/${reviewId}/datalineage`,
      },
      // {
      //   id: "graphs",
      //   name: "Graphs",
      //   hyperlink: `/org/datasets/review/${reviewId}/graphs`,
      // },
    ];
  }
</script>

{#if tabType === "primary"}
  <div class={tabStyles({ ...$$restProps })}>
    <div class="flex">
      {#each tabItems as item}
        <div class="pl-6">
          <div
            class={`${
              tab === item.id
                ? "text-brand-200 border-b-2 border-brand-200 font-medium"
                : "text-slate-500"
            } pb-4 hover:text-slate-850`}
          >
            <a href={item.hyperlink}>
              <span>{item.name}</span>
            </a>
          </div>
        </div>
      {/each}
    </div>
  </div>
{:else if tabType === "secondary"}
  <div class={tabStyles({ ...$$restProps })}>
    <div class="flex">
      {#each tabItems as item}
        <div class="pl-6">
          <div
            class={`${
              tab === item.id ? "bg-slate-200 rounded text-slate-600" : ""
            } px-4 py-2`}
          >
            <a href={item.hyperlink}>
              <span>{item.name}</span>
            </a>
          </div>
        </div>
      {/each}
    </div>
  </div>
{:else}
  <div class={tabStyles({ ...$$restProps })}>
    <div class="flex">
      {#each tabItems as item}
        <div class="pl-6">
          <div
            class={`${
              tab === item.id
                ? "text-slate-600 rounded-lg border border-brand-200"
                : "text-slate-600"
            } px-4 py-2`}
          >
            <a href={item.hyperlink}>
              <span>{item.name}</span>
            </a>
          </div>
        </div>
      {/each}
    </div>
  </div>
{/if}
