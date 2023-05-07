<script lang="ts">
  import { tv } from "tailwind-variants";

  const dropdownStyles = tv({
    slots: {
      base: "focus:outline-none z-50",
      content:
        "dropdown-content bg-white shadow w-52 text-sm text-slate-600 rounded z-50 mt-2",
    },
    variants: {
      intent: {
        start: "flex justify-between items-center",
        end: "flex justify-between items-center",
        // orgType:
        //   "flex justify-between items-center text-sm text-slate-600 rounded border border-slate-600 h-8 px-2",
      },
      fullWidth: {
        true: "w-full",
        false: "w-28",
      },
    },
    defaultVariants: {
      intent: "start",
      fullWidth: true,
    },
  });

  const { base, content } = dropdownStyles();
</script>

{#if $$restProps.intent === "start"}
  <div class="dropdown dropdown-start">
    <label for="trigger" tabIndex={0} class={base()}>
      <slot name="trigger" />
    </label>
    <ul tabIndex={0} class={content()}>
      <slot name="items" />
    </ul>
  </div>
{:else}
  <div class="dropdown dropdown-end">
    <label for="trigger" tabIndex={0} class={base()}>
      <slot />
    </label>
    <ul tabIndex={0} class={content()}>
      <slot name="items" />
    </ul>
  </div>
{/if}
