<script>
  import { page } from "$app/stores";
  import Breadcrumbs from "$lib/components/Breadcrumbs.svelte";
  import Tabbar from "$lib/components/Tabbar.svelte";
  import VerticalLayout from "$lib/components/VerticalLayout.svelte";
  import { Box, FileCheck, Copy, Clock, Edit, Globe } from "lucide-svelte";

  export let data;

  function copy() {
    navigator.clipboard.writeText("text");
    // toast.success("Copied to clipboard!");
  }
</script>

<svelte:head>
  <title>{`Model Card | ${$page.params.modelId}`}</title>
</svelte:head>

<VerticalLayout>
  <div slot="tabbar" class="flex justify-between w-full">
    <Breadcrumbs />
    <Tabbar
      tabType="primary"
      tabContent="primaryModel"
      tab="modelCard"
      fullWidth={false}
    />
  </div>
  <div slot="sidebar">
    <div class="text-slate-800 font-medium text-base">Model details</div>
    <div class="pt-2 text-slate-400">
      <div class="flex justify-between items-center py-1">
        <span class="w-1/7 flex items-center">
          <Box class="w-4" />
        </span>
        <span class="w-1/2 pl-2">Model Name</span>
        <span class="w-1/2 pl-2 text-slate-600 font-medium">
          {`${data.modelDetails[0].name}` || "Model Name"}
        </span>
      </div>
      <div class="flex justify-between items-center py-1">
        <span class="w-1/7 flex items-center">
          <FileCheck class="w-4" />
        </span>
        <span class="w-1/2 pl-2">Model ID</span>
        <span class="w-[48%] text-slate-600 font-medium flex justify-between">
          <div class="w-4/5 pl-2 overflow-hidden truncate">
            {`${data.modelDetails[0].uuid}` || "Model UUID"}
          </div>
          <Copy
            class="text-slate-400 hover:text-slate-600 w-4 cursor-pointer"
            on:click={() => copy()}
          />
        </span>
      </div>
      <div class="flex justify-between items-center py-1">
        <span class="w-1/7 flex items-center">
          <Clock class="w-4" />
        </span>
        <span class="w-1/2 pl-2">Created By</span>
        <span class="w-1/2 pl-2 text-slate-600 font-medium">
          {`${data.modelDetails[0].created_by.name}` || "Created By"}
        </span>
      </div>
      <div class="flex justify-between items-center py-1">
        <span class="w-1/7 flex items-center">
          <Edit class="w-4" />
        </span>
        <span class="w-1/2 pl-2">Last updated by</span>
        <span class="w-1/2 pl-2 text-slate-600 font-medium">
          {data.modelDetails[0].updated_by.name || "User X"}
        </span>
      </div>
      <div class="flex justify-between items-center py-1">
        <span class="w-1/7 flex items-center">
          <Globe class="w-4" />
        </span>
        <span class="w-1/2 pl-2">Public</span>
        <span class="w-1/2 pl-2 text-slate-600 font-medium">
          {`${data.modelDetails[0].is_public ? "Yes" : "No"}`}
        </span>
      </div>
    </div>
  </div>
  <div
    slot="content"
    class="pl-12 pr-8 pt-8 space-y-4 h-[50vh] w-full overflow-auto"
  >
    <!-- {edit ? (
                <>
                  <Form method="post" reloadDocument class="h-full">
                    <div class="flex justify-between h-full">
                      <ClientOnly
                        fallback={
                          <div class="w-2/5">Editor Failed to Load!</div>
                        }
                      >
                        {() => (
                          <Quill defaultValue={html} setContent={setContent} />
                        )}
                      </ClientOnly>
                      <Button
                        intent="icon"
                        fullWidth={false}
                        type="submit"
                        onClick={() => {
                          toast.success("Saved your model card!");
                        }}
                      >
                        <Save class="w-4" />{" "}
                        <div class="pl-2">Save</div>
                      </Button>
                    </div>
                    <input name="content" value={content} class="hidden" />
                  </Form>
                </>
              ) : (
                <>
                  {html ? (
                    <div class="flex justify-between">
                      <div
                        dangerouslySetInnerHTML={{ __html: html }}
                        class="pr-6 text-slate-600"
                      />
                      <Button
                        intent="icon"
                        fullWidth={false}
                        onClick={() => {
                          setEdit(true);
                        }}
                      >
                        <Pencil class="w-4" />{" "}
                        <div class="pl-2">Edit</div>
                      </Button>
                    </div>
                  ) : (
                    <div class="flex flex-col justify-center items-center gap-y-1">
                      <div class="text-base font-medium pt-44">
                        Readme not available
                      </div>
                      <div>
                        No details available in model card. Please contact
                        author of the model.
                      </div>
                    </div>
                  )}
                </>
              )} -->
    <div class="prose">
      {data.modelDetails[0].readme.latest_version.content ||
        "No readme added yet."}
    </div>
  </div>
</VerticalLayout>
