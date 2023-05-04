<script>
  import Button from "$lib/components/Button.svelte";
  import Input from "$lib/components/Input.svelte";
  import logo from "$lib/logos/PureMLLogoText.svg";
</script>

<svelte:head>
  <title>Reset Password | PureML</title>
  <meta name="description" content="Start your journey with PureML." />
</svelte:head>

<div class="w-screen h-screen flex items-center justify-center">
  <div
    class="md:flex md:flex-col justify-center items-center md:pt-0 text-white"
  >
    <div class="w-fit text-center">
      <div class="flex justify-center items-center pb-16">
        <img src={logo} alt="Logo" class="w-36" />
      </div>
      <form method="post" class="text-slate-600 flex flex-col text-left">
        <div class="flex flex-col gap-y-12">
          <div class="flex flex-col gap-y-6">
            <label for="oldpassword" class="font-medium">
              <div class="text-slate-600 text-left">Enter Old password</div>
              <Input
                type="password"
                required
                intent="primary"
                fullWidth={false}
              />
            </label>
            <label for="newpassword" class="font-medium">
              <div class="text-slate-600 text-left">Enter New Password</div>
              <Input
                type="password"
                required
                intent="primary"
                fullWidth={false}
              />
            </label>
          </div>
          <Button intent="primary">Reset</Button>
        </div>
      </form>
    </div>
  </div>
</div>
