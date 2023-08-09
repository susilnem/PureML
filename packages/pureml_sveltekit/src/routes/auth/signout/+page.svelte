<script>
  import Button from "$lib/components/Button.svelte";
  import { base } from '$app/paths';
</script>

<form method="POST">
  <div class="modal-box focus:outline-none rounded-lg w-96">
    <div class="w-full pb-8">
      <div class="text-base text-slate-800 font-medium">
        Do you want to sign out?
      </div>
      <div class="pt-6 text-sm text-slate-600">
        You can come back whenever you want.
      </div>
    </div>
    <div class="flex justify-end gap-x-6">
      <a href="{base}/models"
        ><Button intent="secondary" fullWidth={false}>Cancel</Button></a
      >
      <Button intent="primary" fullWidth={false}>Logout</Button>
    </div>
  </div>
</form>
