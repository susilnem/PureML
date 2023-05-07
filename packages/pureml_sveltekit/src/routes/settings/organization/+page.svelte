<script>
  import Button from "$lib/components/Button.svelte";
  import Input from "$lib/components/Input.svelte";
  import Tabbar from "$lib/components/Tabbar.svelte";
  import { Copy } from "lucide-svelte";

  function copyOrgId() {
    navigator.clipboard.writeText("text");
    // toast.success("Copied to clipboard!");
  }
</script>

<div class="flex justify-center w-full border-b-2 border-slate-100">
  <div class="w-full 2xl:max-w-screen-2xl">
    <Tabbar tabType="primary" tabContent="primarySettings" tab="organization" />
  </div>
</div>
<div class="flex justify-center w-full">
  <div
    class="bg-slate-50 flex flex-col h-screen overflow-hidden w-full 2xl:max-w-screen-2xl"
  >
    <form
      method="post"
      class="py-8 px-12 w-full h-[80%] overflow-auto text-slate-600"
    >
      <div class="pb-4">
        <label for="orgname" class="labelInput pb-1">
          <div class="text-left">Organization Name</div>
          <Input
            intent="primary"
            type="text"
            name="orgname"
            fullWidth={false}
            defaultValue="orgData[0].name"
            aria-label="orgname"
            data-testid="orgname"
            required
          />
        </label>
      </div>
      <div class="pb-4">
        <label for="orgdesc" class="labelInput">
          <div class="text-left">Organization Description</div>
          <Input
            intent="primary"
            type="text"
            name="orgdesc"
            fullWidth={false}
            defaultValue="Enter email..."
            aria-label="orgdesc"
            data-testid="orgdesc"
            required
          />
        </label>
      </div>
      <div class="pb-8">
        <label for="orgid" class="labelInput pb-1">
          <div class="text-left">Organization ID</div>
          <div class="input-icons">
            <input class="hidden" name="orgid" value="uuid" />
            <input
              id="orgid"
              type="text"
              name="orgid"
              value="Organization ID"
              aria-label="orgid"
              data-testid="orgid"
              required
              disabled
              class="input-field rounded !w-[24rem]"
            />
            <i>
              <Copy
                class="text-slate-400 hover:text-slate-600 w-4 cursor-pointer"
                on:click={() => copyOrgId()}
              />
            </i>
          </div>
        </label>
      </div>
      <div class="w-fit">
        <Button fullWidth={false}>Save changes</Button>
      </div>
      <!-- <div class="pt-12 font-medium text-slate-800">
              Delete Organization
            </div>
            <div class="pb-8 text-slate-600">
              Delete this organization permanently, this action is irreversible
              All its repositories (models, datasets) will be deleted.
            </div>
            <div class="w-fit">
              <Button intent="danger" fullWidth={false}>
                Delete Organization
              </Button>
            </div> -->
    </form>
  </div>
</div>
