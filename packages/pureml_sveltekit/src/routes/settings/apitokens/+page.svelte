<script>
  import { enhance } from "$app/forms";
  import Modal from "$lib/components/Modal.svelte";
  import Tabbar from "$lib/components/Tabbar.svelte";
  import { Copy } from "lucide-svelte";
  import toast from "svelte-french-toast";

  function copyOrgId() {
    navigator.clipboard.writeText("text");
    toast.success("Copied to clipboard!");
  }

  export let data;
  /** @type {import('./$types').ActionData} */
  export let form;

  if (form) {
    if (form.removeToken) {
      if (form.removeToken.status === 200 || form.removeToken.status === 202)
        toast.success("API Key ID deleted successfully!");
      else toast.error("Something went wrong!");
    }
  }
</script>

<svelte:head>
  <title>API Tokens | Settings</title>
</svelte:head>

<div class="flex justify-center w-full border-b-2 border-slate-100">
  <div class="w-full 2xl:max-w-screen-2xl px-6">
    <Tabbar tabType="primary" tabContent="primarySettings" tab="apiToken" />
  </div>
</div>
<div class="flex justify-center w-full">
  <div class="bg-slate-50 h-screen overflow-hidden w-full 2xl:max-w-screen-2xl">
    <div
      class="flex flex-col gap-y-10 pt-8 pb-12 px-12 w-full h-4/5 overflow-auto rounded-tl-lg rounded-tr-lg"
    >
      <div class="font-medium">API Token</div>
      <div class="overflow-auto max-h-[57vh] rounded-tl-lg rounded-tr-lg">
        <table class="table w-3/4 xl:w-2/5">
          <thead>
            <tr class="border-b border-slate-200">
              <th>API Key ID</th>
              <th>Created</th>
              <th>Actions</th>
            </tr>
          </thead>
          <tbody>
            {#if data && data.apiTokens}
              {#each data.apiTokens as token}
                <tr class="border border-slate-200 hover:bg-slate-50">
                  <td class="text-lg">{token.uuid}</td>
                  <td class="text-lg">
                    {token.created_at
                      .replace("T", " ")
                      .replace("Z", "")
                      .slice(0, 19)}
                  </td>
                  <td class="text-lg">
                    <form method="POST" class="w-full">
                      <input
                        class="hidden"
                        type="text"
                        name="deleteToken"
                        value={token.uuid}
                      />
                      <button
                        class="text-lg p-2 hover:bg-slate-100 hover:rounded-lg"
                      >
                        Delete
                      </button>
                    </form>
                  </td>
                </tr>
              {/each}
            {:else}
              <tr class="border border-slate-200 hover:bg-slate-50"
                ><div class="p-4">No API keys found</div></tr
              >
            {/if}
          </tbody>
        </table>
      </div>
      <form method="POST" use:enhance>
        <input
          class="hidden"
          type="text"
          name="createToken"
          value="createToken"
        />
        <Modal title="Token Created" open={false}
          ><p slot="btnName">Create API Key</p>
          <div class="text-slate-600">
            This is the last time you can see the API Key Secret! Sorry,
            security :)
          </div>
          {#if form && form.newToken}
            <div class="flex flex-col gap-y-6 pt-8">
              <label for="keyid" class="pb-1">
                API Key ID
                <div class="input-icons">
                  <input
                    id="keyid"
                    type="text"
                    name="keyid"
                    value={form.newToken.data[0].uuid || "API Key ID"}
                    aria-label="keyid"
                    data-testid="keyid"
                    required
                    disabled
                    class="input-field rounded !w-full truncate"
                  />
                  <i>
                    <Copy
                      class="text-slate-400 hover:text-slate-600 w-4 cursor-pointer"
                      on:click={() => copyOrgId()}
                    />
                    copyOrgId(form.newToken.data[0].uuid)
                  </i>
                </div>
              </label>
              <label for="keysecret" class="pb-1">
                API Key Secret
                <div class="input-icons">
                  <input
                    id="orgid"
                    type="text"
                    name="orgid"
                    value={form.newToken.data[0].api_token_secret ||
                      "API Key Secret"}
                    aria-label="orgid"
                    data-testid="orgid"
                    required
                    disabled
                    class="input-field rounded !w-full !pr-12 truncate"
                  />
                  <i>
                    <Copy
                      class="text-slate-400 hover:text-slate-600 w-4 cursor-pointer"
                      on:click={() => copyOrgId()}
                    />
                    copyOrgId(form.newToken.data[0].api_token_secret)}
                  </i>
                </div>
              </label>
            </div>
          {/if}
        </Modal>
      </form>
    </div>
  </div>
</div>
