<script lang="ts">
  import type { VariantProps } from "tailwind-variants";
  import { tv } from "tailwind-variants";

  const tagStyles = tv({
    base: "badge badge-secondary badge-outline badge-lg rounded !text-slate-600",
    variants: {
      intent: {
        text: "border-slate-100 pt-2 pb-2",
        icon: "bg-slate-50 border-slate-100",
      },
    },
    defaultVariants: {
      intent: "text",
    },
  });

  interface $$Props extends VariantProps<typeof tagStyles> {}
</script>

<div class="pr-2">
  <div class={tagStyles({ ...$$restProps })}>
    {#if $$restProps.intent === "text"}
      <div><slot /></div>
    {:else}
      <i><slot name="icon" /></i>
      <div class="pl-2 text-xs"><slot name="text" /></div>
    {/if}
  </div>
</div>
