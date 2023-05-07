<script lang="ts">
  import type { VariantProps } from "tailwind-variants";
  import { tv } from "tailwind-variants";

  const cardStyles = tv({
    base: "text-slate-600 border border-slate-200 justify-center focus:outline-none",
    variants: {
      intent: {
        badged:
          "font-normal rounded-lg cursor-pointer px-6 py-4 hover:bg-slate-100 transform transition duration-300",
        info: "items-center p-8 text-lg font-medium rounded-xl",
      },
      fullWidth: {
        true: "w-full",
      },
    },
    defaultVariants: {
      intent: "badged",
      fullWidth: true,
    },
  });

  export let title: string;
  export let description: string;
  export let count: string;
</script>

{#if $$restProps.intent === "badged"}
  <div class={cardStyles({ ...$$restProps })}>
    <header class="pb-0 text-slate-600">
      <div class="flex items-center">
        <slot name="icon" />
        <!-- <Box class="text-slate-400 w-4" /> -->
        <span class="ml-2 truncate font-medium">{title}</span>
      </div>
    </header>
    <div class="text-sm font-normal truncate pt-2">
      {description}
    </div>
    <div class="flex pt-4">
      <slot name="tag" />
    </div>
  </div>
{:else}
  <div class={cardStyles({ ...$$restProps })}>
    <div class="text-3xl">{count}</div>
    <div class="text-lg">{title}</div>
  </div>
{/if}
