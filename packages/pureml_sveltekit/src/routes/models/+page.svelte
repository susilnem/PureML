<script>
  import Card from "$lib/components/Card.svelte";
  import Tag from "$lib/components/Tag.svelte";
  import { Box, Globe } from "lucide-svelte";
</script>

<div class="flex justify-center w-full">
  <div
    class="bg-slate-50 flex flex-col h-screen overflow-hidden w-full 2xl:max-w-screen-2xl"
  >
    <div
      class="flex justify-between px-12 pt-6 pb-4 font-medium text-slate-800"
    >
      Models
    </div>
    <div class="h-[84%] overflow-auto pb-8">
      <div
        class="px-12 pt-2 grid justify-center grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-6 4xl:grid-cols-10 gap-8 min-w-72"
      >
        <a href="/org/pureml/models/churn">
          <Card
            intent="badged"
            title="Model Name"
            description="My first model"
            count=""
            ><p slot="icon"><Box class="text-slate-400 w-4" /></p>
            <p slot="tag">
              <Tag
                ><p slot="icon"><Globe class="text-slate-400 w-4" /></p>
                <p slot="text">Public</p></Tag
              >
            </p></Card
          >
        </a>
      </div>
    </div>
  </div>
</div>
