<script>
  import Tabbar from "$lib/components/Tabbar.svelte";
  import Select from "$lib/components/Select.svelte";
  import { ListboxOption } from "@rgossiaux/svelte-headlessui";
  import Avatar from "$lib/components/Avatar.svelte";
  import Modal from "$lib/components/Modal.svelte";
  import Input from "$lib/components/Input.svelte";
  import Button from "$lib/components/Button.svelte";
  import { PlusCircle } from "lucide-svelte";

  export let data;

  const listRoles = [
    { id: 1, name: "Owner" },
    { id: 2, name: "Member" },
  ];
</script>

<div class="flex justify-center w-full border-b-2 border-slate-100">
  <div class="w-full 2xl:max-w-screen-2xl">
    <Tabbar tabType="primary" tabContent="primarySettings" tab="members" />
  </div>
</div>
<div class="flex justify-center w-full">
  <div class="bg-slate-50 flex flex-col h-screen w-full 2xl:max-w-screen-2xl">
    <div class="pt-8 px-12">
      <div class="flex justify-between pb-10">
        <div class="font-medium">Organization Members</div>
        <Modal title="Add Member">
          <p slot="btnName">
            <Button intent="primary" fullWidth={false}>
              <PlusCircle class="text-white w-4" />
              <div class="pl-2">Add member</div>
            </Button>
          </p>
          <form method="post" class="w-full">
            <div>
              <input class="hidden" type="text" name="orgid" />
              <!-- value={data[0].uuid} -->
              <label for="addEmail" class="text-sm pb-1">
                Email
                <Input
                  intent="valuePrimary"
                  type="text"
                  name="addEmail"
                  fullWidth={false}
                  defaultValue=""
                  aria-label="addEmail"
                  data-testid="addEmail"
                  required
                />
              </label>
            </div>
            <div class="pt-12 grid justify-items-end w-full">
              <div class="flex justify-end w-1/2">
                <Button intent="primary" fullWidth={false}>Add</Button>
              </div>
            </div>
          </form>
        </Modal>
      </div>
      <div class="overflow-auto h-[57vh] rounded-tl-lg rounded-tr-lg">
        <table class="table w-full">
          <thead>
            <tr class="border-b border-slate-200">
              <th>Name</th>
              <th>Email</th>
              <th>Role</th>
              <th>{""}</th>
            </tr>
          </thead>
          <tbody>
            {#if data.orgDetails[0].members}
              {#each data.orgDetails[0].members as member}
                <tr class="hover:bg-slate-100">
                  <!-- <th class="w-4">
                        <input type="checkbox" class="checkbox" />
                      </th> -->
                  <td class="lg:w-96 2xl:w-[56rem]">
                    <div class="flex items-center space-x-3">
                      <div class="mask w-12 flex items-center">
                        <Avatar intent="large">
                          {member.name.charAt(0).toUpperCase() || "U"}
                        </Avatar>
                      </div>
                      <div>
                        <div class="text-slate-600">
                          {member.name || "Name"}
                        </div>
                        <div class="text-sm text-slate-400">
                          {member.handle || "username"}
                        </div>
                      </div>
                    </div>
                  </td>
                  <td class="lg:w-96 2xl:w-[56rem]">
                    {member.email || "abc@gmail.com"}
                  </td>
                  <td>
                    <form method="post">
                      <!-- onChange={roleChange} -->
                      <input class="hidden" type="text" name="orgid" />
                      <!-- value={data[0].uuid} -->
                      <input class="hidden" type="text" name="roleEmail" />
                      <!-- value={member.email} -->
                      <!-- <SelectPrimitive.Root name="role">
                                <SelectPrimitive.Trigger
                                  asChild
                                  aria-label="Role"
                                  class="focus:outline-none"
                                >
                                  <button class="flex foxus:outline-none">
                                    <SelectPrimitive.Icon class="capitalize focus:outline-none flex justify-center items-center text-sm gap-x-2 text-slate-600 border border-slate-200 rounded px-2 py-1">
                                      <div class="hidden">{role}</div>
                                      {member.role}{" "}
                                      <ChevronDown class="w-4" />
                                    </SelectPrimitive.Icon>
                                  </button>
                                </SelectPrimitive.Trigger>
                                <SelectPrimitive.Content
                                  position="popper"
                                  sideOffset={7}
                                  align="start"
                                  class="z-50"
                                >
                                  <SelectPrimitive.Viewport class="bg-white rounded shadow-lg">
                                    <SelectPrimitive.Group>
                                      {allRoles.map(
                                        (role: any, index: number) => (
                                          <SelectPrimitive.Item
                                            key={`${role}-${index}`}
                                            value={role.value}
                                            class="flex items-center w-24 px-2 py-2 rounded text-sm text-slate-600 cursor-pointer hover:bg-slate-100 hover:border-none focus:outline-none"
                                          >
                                            <SelectPrimitive.ItemText class="text-sm">
                                              {role.label}
                                            </SelectPrimitive.ItemText>
                                          </SelectPrimitive.Item>
                                        )
                                      )}
                                    </SelectPrimitive.Group>
                                  </SelectPrimitive.Viewport>
                                </SelectPrimitive.Content>
                              </SelectPrimitive.Root> -->
                      <Select intent="primary" name="members" title="Role"
                        >{#each listRoles as role (role.id)}
                          <ListboxOption
                            value={role}
                            class="p-2 border border-slate-200 hover:bg-slate-200"
                          >
                            {role.name}
                          </ListboxOption>
                        {/each}</Select
                      >
                    </form>
                  </td>
                  <th class="lg:w-96">
                    <!-- <Modal
                              btnName={
                                <button class="font-normal">Remove</button>
                              }
                              title="Remove member"
                            >
                              <form
                                method="post"
                                class="w-full"
                              >
                                <input
                                  class="hidden"
                                  type="text"
                                  name="orgid"
                                  defaultValue={data[0].uuid}
                                />
                                <input
                                  class="hidden"
                                  type="text"
                                  name="removeEmail"
                                  defaultValue={member.email}
                                />
                                <div class="text-slate-600">
                                  Are you sure you want to remove this user?
                                </div>
                                <div class="pt-12 grid justify-items-end w-full">
                                  <div class="flex justify-end w-1/2">
                                    <Button
                                      intent="primary"
                                      type="submit"
                                      class="pl-2"
                                      fullWidth={false}
                                    >
                                      Remove
                                    </Button>
                                  </div>
                                </div>
                              </form>
                            </Modal> -->
                    Remove
                  </th>
                </tr>
              {/each}
            {:else}
              <div>No members found</div>
            {/if}
          </tbody>
        </table>
      </div>
    </div>
  </div>
</div>
