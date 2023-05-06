<script>
  import Card from "$lib/components/Card.svelte";
  import { Database, Fingerprint } from "lucide-svelte";
  import Avatar from "$lib/components/Avatar.svelte";
  import Tag from "$lib/components/Tag.svelte";
</script>

<svelte:head>
  <title>Organization Name | PureML</title>
  <meta name="description" content="Start your journey with PureML." />
</svelte:head>

<div class="flex h-[93%] overflow-auto">
  <div class="w-full pt-8 px-12">
    <!-- ###### org details ##### -->

    <div class="p-6 bg-slate-100 flex items-center rounded-lg">
      <Avatar intent="large">
        <!-- {orgData.orgDetails[0].name.charAt(0)} -->
        P
      </Avatar>
      <div class="pl-8">
        <div class="font-medium text-slate-600">
          <!-- {orgData.orgDetails[0].name} -->
          PureML
        </div>
        <div class="text-sm text-slate-600">
          <!-- {orgData.orgDetails[0].description} -->
          This is a dummy organization
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
        class="pt-2 grid justify-center grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-6 4xl:grid-cols-10 gap-8 min-w-72"
      >
        <a href="https://pureml.com">
          <Card
            intent="badged"
            title="Dataset Name"
            description="My first dataset"
            count=""
            ><p slot="icon"><Database class="text-slate-400 w-4" /></p>
            <p slot="tag">
              <Tag
                ><p slot="icon"><Fingerprint class="text-slate-400 w-4" /></p>
                <p slot="text">Public</p></Tag
              >
            </p></Card
          >
        </a>
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
        class="pt-2 grid justify-center grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-6 4xl:grid-cols-10 gap-8 min-w-72"
      >
        <a href="https://pureml.com">
          <Card
            intent="badged"
            title="Dataset Name"
            description="My first dataset"
            count=""
            ><p slot="icon"><Database class="text-slate-400 w-4" /></p>
            <p slot="tag">
              <Tag
                ><p slot="icon"><Fingerprint class="text-slate-400 w-4" /></p>
                <p slot="text">Public</p></Tag
              >
            </p></Card
          >
        </a>
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
