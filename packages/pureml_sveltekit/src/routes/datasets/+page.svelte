<script>
  import Card from "$lib/components/Card.svelte";
  import Tag from "$lib/components/Tag.svelte";
  import { Database, Fingerprint, Globe } from "lucide-svelte";

  export let data;
</script>

<div class="flex justify-center w-full">
  <div
    class="bg-slate-50 flex flex-col h-screen overflow-hidden w-full 2xl:max-w-screen-2xl"
  >
    <div
      class="flex justify-between px-12 pt-6 pb-4 font-medium text-slate-800"
    >
      Datasets
    </div>
    <div class="h-[83%] overflow-auto">
      <div
        class="px-12 pt-2 grid justify-center grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-6 4xl:grid-cols-10 gap-8 min-w-72"
      >
        {#if data && data.datasets?.length > 0}
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
                      <p slot="text" class="text-sm">Public</p></Tag
                    >
                  {:else}
                    <Tag
                      ><p slot="icon">
                        <Fingerprint class="text-slate-400 w-2.5" />
                      </p>
                      <p slot="text" class="text-sm">Private</p></Tag
                    >
                  {/if}
                </p></Card
              >
            </a>
          {/each}
        {:else}
          <div class="rounded-sm border-2 border-slate-200 px-6 py-4">
            <div class="font-medium text-sm pb-6">
              There are no datasets yet
            </div>
            <div class="rounded-lg h-2 bg-slate-200 w-1/3" />
            <div class="pt-2" />
            <div class="rounded-lg h-2 bg-slate-200 w-2/3" />
          </div>
        {/if}
      </div>
    </div>
  </div>
</div>
