<script lang="ts">
  import { X } from "lucide-svelte";
  import { tv } from "tailwind-variants";

  const modalBtnStyles = tv({
    base: "font-medium",
    variants: {
      intent: {
        primary: "btn btn-primary btn-sm text-white",
        secondary:
          "btn btn-secondary btn-sm bg-transparent text-slate-600 border border-slate-400",
      },
    },
    defaultVariants: {
      intent: "primary",
    },
  });

  export let title: string;

  export let open = false;
  console.log(open);
</script>

<!-- The button to open modal -->
<button
  class={modalBtnStyles({ ...$$restProps })}
  on:click={() => (open = !open)}
>
  <label for="modalBtn"><slot name="btnName" /></label>
</button>

<!-- Body -->
<input type="checkbox" id="modalBtn" class="modal-toggle" bind:checked={open} />
<div class="modal">
  <div class="modal-box relative focus:outline-none rounded-lg">
    <label
      for="modalBtn"
      class="rounded-full h-6 w-6 flex items-center justify-center absolute top-4 right-4"
    >
      <X />
    </label>
    <div class="w-full">
      <div class="text-base text-slate-800 font-medium">
        {title}
      </div>
      <div class="pt-6 text-sm text-slate-600">
        <slot />
      </div>
    </div>
  </div>
</div>
