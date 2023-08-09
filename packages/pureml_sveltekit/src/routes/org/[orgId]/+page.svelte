<script>
  import Card from "$lib/components/Card.svelte";
  import { Box, Database, Fingerprint, Globe } from "lucide-svelte";
  import Avatar from "$lib/components/Avatar.svelte";
  import Tag from "$lib/components/Tag.svelte";

  export let data;
</script>

<svelte:head>
  <title>{data.orgDetails[0].name} | PureML</title>
</svelte:head>

<div class="flex justify-center h-[90%] overflow-auto">
  <div class="w-full max-w-screen-2xl pt-8 px-12">
    <!-- ###### org details ##### -->
    <div class="p-6 bg-slate-100 flex items-center rounded-lg">
      <Avatar intent="large">
        {data.orgDetails[0].name.charAt(0)}
      </Avatar>
      <div class="pl-8">
        <div class="font-medium text-slate-600">
          {data.orgDetails[0].name}
        </div>
        <div class="text-sm text-slate-600">
          {data.orgDetails[0].description}
        </div>
      </div>
    </div>

    <!-- ###### org models displayed here ##### -->
    <div>
      <div
        class="flex justify-between font-medium text-slate-800 text-base pt-6"
      >
        Models
      </div>
      <div
        class="pt-2 grid justify-center grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-5 4xl:grid-cols-6 gap-8 min-w-72"
      >
        {#if data && data.models != ""}
          {#each data.models as model}
            <a href={`/org/${data.orgDetails[0].name}/models/${model.name}`}>
              <Card
                intent="badged"
                title={model.name}
                description={`Updated by ${model.updated_by.name}`}
                count=""
                ><p slot="icon">
                  <Box class="text-slate-400 w-4" />
                </p>
                <p slot="tag">
                  {#if model.is_public === true}
                    <Tag
                      ><p slot="icon"><Globe class="text-slate-400 w-2.5" /></p>
                      <p slot="text">Public</p></Tag
                    >
                  {:else}
                    <Tag
                      ><p slot="icon">
                        <Fingerprint class="text-slate-400 w-2.5" />
                      </p>
                      <p slot="text">Private</p></Tag
                    >
                  {/if}
                </p></Card
              >
            </a>
          {/each}
        {:else}
          <div class="rounded-sm border-2 border-slate-200 px-6 py-4">
            <div class="font-medium text-sm pb-6">There are no models yet</div>
            <div class="rounded-lg h-2 bg-slate-200 w-1/3" />
            <div class="pt-2" />
            <div class="rounded-lg h-2 bg-slate-200 w-2/3" />
          </div>
        {/if}
      </div>
    </div>

    <!-- ###### org datasets displayed here ##### -->
    <div>
      <div
        class="flex justify-between font-medium text-slate-800 text-base pt-6"
      >
        Datasets
      </div>
      <div
        class="pt-2 grid justify-center grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-5 4xl:grid-cols-6 gap-8 min-w-72"
      >
        {#if data && data.datasets != ""}
          {#each data.datasets as dataset}
            <a
              href={`/org/${data.orgDetails[0].name}/datasets/${dataset.name}`}
            >
              <Card
                intent="badged"
                title={dataset.name}
                description={`Updated by ${dataset.updated_by.name}`}
                count=""
                ><p slot="icon">
                  <Database class="text-slate-400 w-4" />
                </p>
                <p slot="tag">
                  {#if dataset.is_public === true}
                    <Tag
                      ><p slot="icon"><Globe class="text-slate-400 w-2.5" /></p>
                      <p slot="text">Public</p></Tag
                    >
                  {:else}
                    <Tag
                      ><p slot="icon">
                        <Fingerprint class="text-slate-400 w-2.5" />
                      </p>
                      <p slot="text">Private</p></Tag
                    >
                  {/if}
                </p></Card
              >
            </a>
          {/each}
        {:else}
          <div class="rounded-sm border-2 border-slate-200 px-6 py-4">
            <div class="font-medium text-sm pb-6">There are no models yet</div>
            <div class="rounded-lg h-2 bg-slate-200 w-1/3" />
            <div class="pt-2" />
            <div class="rounded-lg h-2 bg-slate-200 w-2/3" />
          </div>
        {/if}
      </div>
    </div>
  </div>

  <!-- ###### org activity section #####
        <div class="hidden md:block fixed right-0 h-full w-1/5 border-l-2 border-slate-200">
          <div class="text-slate-900 font-medium text-base px-8 py-6">
            Activity
          </div>
        </div> -->
</div>
