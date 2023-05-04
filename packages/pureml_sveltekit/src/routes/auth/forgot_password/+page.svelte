<script>
  import Button from "$lib/components/Button.svelte";
  import Input from "$lib/components/Input.svelte";
  import logo from "$lib/logos/PureMLLogoText.svg";
</script>

<div class="flex justify-center">
  <div class="w-fit text-center">
    <div class="flex justify-center items-center pb-16">
      <img src={logo} alt="Logo" class="w-36" />
    </div>
    <form method="post" class="text-slate-600 flex flex-col text-left">
      <div class="flex flex-col gap-y-12">
        <label for="email" class="font-medium">
          <div class="text-slate-600 text-left">Email</div>
          <Input type="email" required intent="primary" fullWidth={false} />
        </label>
        <Button intent="primary">Send Link</Button>
      </div>
      <span class="w-[24rem] text-slate-600 pt-6">
        Reset password link will be sent to your mail ID given above. Click link
        to change or reset password.
      </span>
    </form>
  </div>
</div>
