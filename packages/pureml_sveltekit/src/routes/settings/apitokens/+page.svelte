<script>
  import Tabbar from "$lib/components/Tabbar.svelte";
</script>

<div class="flex justify-center w-full border-b-2 border-slate-100">
  <div class="w-full 2xl:max-w-screen-2xl">
    <Tabbar tabType="primary" tabContent="primarySettings" tab="apiToken" />
  </div>
</div>
<div class="flex justify-center w-full">
  <div class="bg-slate-50 h-screen overflow-hidden w-full 2xl:max-w-screen-2xl">
    <div
      class="flex flex-col gap-y-10 pt-8 pb-12 px-12 w-full h-4/5 overflow-auto rounded-tl-lg rounded-tr-lg"
    >
      <div class="font-medium">API Token</div>
      <div class="overflow-auto h-[57vh] rounded-tl-lg rounded-tr-lg">
        <table class="table w-3/4 xl:w-2/5">
          <thead>
            <tr class="border-b border-slate-200">
              <th>API Key ID</th>
              <th>Created</th>
              <th>Actions</th>
            </tr>
          </thead>
          <tbody>
            <!-- {data ? (
                  <>
                    {data.map((token: any, index: number) => (
                      <tr
                        class="border border-slate-200 hover:bg-slate-50"
                        key={index}
                      >
                        <td class="text-lg">{token.uuid}</td>
                        <td class="text-lg">
                          {token.created_at
                            .replace("T", " ")
                            .replace("Z", "")
                            .slice(0, 19)}
                        </td>
                        <td class="text-lg">
                          <Form method="post" reloadDocument class="w-full">
                            <input
                              class="hidden"
                              type="text"
                              name="deleteToken"
                              defaultValue={token.uuid}
                            />
                            <button class="text-lg p-2 hover:bg-slate-100 hover:rounded-lg">
                              Delete
                            </button>
                          </Form>
                        </td>
                      </tr>
                    ))}
                  </>
                ) : (
                  <div>No API keys found</div>
                )} -->
          </tbody>
        </table>
      </div>
      <form method="post">
        <input
          class="hidden"
          type="text"
          name="createToken"
          value="createToken"
        />
        <!-- <Modal
                btnName={
                  <Button intent="primary" fullWidth={false}>
                    Create API Key
                  </Button>
                }
                title="Token created"
              >
                <div class="text-slate-600">
                  This is the last time you can see the API Key Secret! Sorry,
                  security...
                </div>
                {actionData && actionData.newToken && (
                  <div class="flex flex-col gap-y-6 pt-8">
                    <label htmlFor="keyid" class="pb-1">
                      API Key ID
                      <div class="input-icons">
                        {console.log("mod=", actionData)}
                        <input
                          id="keyid"
                          type="text"
                          name="keyid"
                          defaultValue={
                            actionData.newToken.data[0].uuid || "API Key ID"
                          }
                          value={actionData.newToken.data[0].uuid}
                          aria-label="keyid"
                          data-testid="keyid"
                          required
                          disabled
                          class="input-field rounded !w-full truncate"
                        />
                        <i>
                          <Copy
                            class="text-slate-400 hover:text-slate-600 w-4 cursor-pointer"
                            onClick={() =>
                              copyOrgId(actionData.newToken.data[0].uuid)
                            }
                          />
                        </i>
                      </div>
                    </label>
                    <label htmlFor="keysecret" class="pb-1">
                      API Key Secret
                      <div class="input-icons">
                        <input
                          id="orgid"
                          type="text"
                          name="orgid"
                          defaultValue={
                            actionData.newToken.data[0].api_token_secret ||
                            "API Key Secret"
                          }
                          value={actionData.newToken.data[0].api_token_secret}
                          aria-label="orgid"
                          data-testid="orgid"
                          required
                          disabled
                          class="input-field rounded !w-full !pr-12 truncate"
                        />
                        <i>
                          <Copy
                            class="text-slate-400 hover:text-slate-600 w-4 cursor-pointer"
                            onClick={() =>
                              copyOrgId(
                                actionData.newToken.data[0].api_token_secret
                              )
                            }
                          />
                        </i>
                      </div>
                    </label>
                  </div>
                )}
              </Modal> -->
      </form>
    </div>
  </div>
</div>
