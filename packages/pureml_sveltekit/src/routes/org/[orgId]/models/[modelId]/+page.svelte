<script>
  import Breadcrumbs from "$lib/components/Breadcrumbs.svelte";
  import Tabbar from "$lib/components/Tabbar.svelte";
  import { Box, FileCheck, Copy, Clock, Edit, Globe } from "lucide-svelte";

  function copy() {
    navigator.clipboard.writeText("text");
    // toast.success("Copied to clipboard!");
  }
</script>

<div
  class="flex justify-center sticky top-0 bg-slate-50 w-full border-b border-slate-200"
>
  <div class="flex justify-between px-12 2xl:pr-0 w-full max-w-screen-2xl">
    <Breadcrumbs />
    <Tabbar tabType="primary" tabContent="primaryModel" tab="modelCard" fullWidth={false} />
  </div>
</div>
<div class="flex justify-center w-full">
  <div
    class="bg-slate-50 flex flex-col h-screen overflow-hidden w-full 2xl:max-w-screen-2xl"
  >
    <div class="flex justify-between h-full">
      <div class="pl-12 pr-8 pt-8 space-y-4 h-[50vh] w-full overflow-auto">
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
                        class="pr-6 text-slate-600 text-sm"
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
      </div>
      <div
        class="px-6 pt-8 w-1/4 max-w-[400px] bg-slate-50 border-l-2 border-slate-100"
      >
        <div class="text-slate-800 font-medium text-base">Model details</div>
        <div class="pt-2 text-sm text-slate-400">
          <div class="flex justify-between items-center py-1">
            <span class="w-1/7 flex items-center">
              <Box class="w-4" />
            </span>
            <span class="w-1/2 pl-2">Name</span>
            <span class="w-1/2 pl-2 text-slate-600 font-medium">
              <!-- {`${modelDetails[0].name}` || "Model Name"} -->
              Churn
            </span>
          </div>
          <div class="flex justify-between items-center py-1">
            <span class="w-1/7 flex items-center">
              <FileCheck class="w-4" />
            </span>
            <span class="w-1/2 pl-2">Model Id</span>
            <span
              class="w-[48%] text-slate-600 font-medium flex justify-between"
            >
              <div class="w-4/5 pl-2 overflow-hidden truncate">
                <!-- {`${modelDetails[0].uuid}` || "Model Id"} -->
                123456789
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
              <!-- {`${modelDetails[0].created_by.name}` || "Created By"} -->
              Person
            </span>
          </div>
          <div class="flex justify-between items-center py-1">
            <span class="w-1/7 flex items-center">
              <Edit class="w-4" />
            </span>
            <span class="w-1/2 pl-2">Last updated by</span>
            <span class="w-1/2 pl-2 text-slate-600 font-medium">
              <!-- {modelDetails[0].updated_by.name || "User X"} -->
              Person
            </span>
          </div>
          <div class="flex justify-between items-center py-1">
            <span class="w-1/7 flex items-center">
              <Globe class="w-4" />
            </span>
            <span class="w-1/2 pl-2">Public</span>
            <span class="w-1/2 pl-2 text-slate-600 font-medium">
              <!-- {`${modelDetails[0].is_public ? "Yes" : "No"}`} -->
              Yes
            </span>
          </div>
        </div>
      </div>
    </div>
  </div>
</div>
