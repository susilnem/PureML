<script>
  import Button from "$lib/components/Button.svelte";
  import Input from "$lib/components/Input.svelte";
  import Tabbar from "$lib/components/Tabbar.svelte";
</script>

<div class="flex justify-center w-full border-b-2 border-slate-100">
  <div class="w-full 2xl:max-w-screen-2xl">
    <Tabbar tabType="primary" tabContent="primarySettings" tab="profile" />
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
        <label for="name" class="labelInput">
          <div class="text-left">Name</div>
          <Input
            intent="primary"
            type="text"
            name="name"
            fullWidth={false}
            defaultValue="Enter Name"
            required
          />
        </label>
      </div>
      <div class="pb-4">
        <label for="bio" class="labelInput">
          <div class="text-left">Bio</div>
          <Input
            intent="primary"
            type="text"
            name="bio"
            fullWidth={false}
            defaultValue="profileData[0].bio"
            aria-label="bio"
            data-testid="bio"
            required
          />
        </label>
      </div>
      <div class="pb-4">
        <label for="email" class="labelInput">
          <div class="text-left">Email</div>
          <Input
            intent="primary"
            type="text"
            name="email"
            fullWidth={false}
            defaultValue="profileData[0].email"
            aria-label="email"
            data-testid="email"
            required
            disabled
          />
        </label>
      </div>
      <div class="pb-8">
        <label for="username" class="labelInput">
          <div class="text-left">Username</div>
          <Input
            intent="primary"
            type="text"
            name="username"
            fullWidth={false}
            defaultValue="profileData[0].handle"
            aria-label="username"
            data-testid="username"
            required
          />
        </label>
      </div>
      <div class="w-fit">
        <Button fullWidth={false}>Save changes</Button>
      </div>
    </form>
  </div>
</div>
